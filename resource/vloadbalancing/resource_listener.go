package vloadbalancing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vloadbalancing"
)

const (
	errorListenerCreate         = "error creating Listener: %s"
	errorListenerRead           = "error reading Listener with Id(%s): %s"
	errorListenerDelete         = "error deleting Listener (%s): %s"
	errorListenerUpdate         = "error updating Listener (%s): %s"
	errorCheckingListenerStatus = "error checking Listener: %s"
	errorWaitingListenerUpdate  = "error waiting for Listener (%s) to be ready: %s"
	errorWaitingListenerCreate  = "error waiting for Listener (%s) to be ready: %s"
	errorImportListener         = "import format error: to import a Listener use the format {projectId}:{loadBalancerId}:{listenerId}"
	errorListenerNotFound       = "[WARN] Listener (%s) not found, removing from state"
	errorRetryDeleteListener    = "[DEBUG] Retrying Deleting Listener with Id(%s) follow error: %s"
	errorRetryUpdateListener    = "[DEBUG] Retrying Updating Listener with Id(%s) follow error: %s"
	errorRetryCreateListener    = "[DEBUG] Retrying Creating Listener with follow error: %s"
)

func ResourceListener() *schema.Resource {
	return resourceListenerInstance()
}

func resourceListenerInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceListenerCreate,
		ReadContext:   resourceListenerRead,
		UpdateContext: resourceListenerUpdate,
		DeleteContext: resourceListenerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceListenerImportState,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowed_cidrs": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0/0",
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"timeout_client": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"timeout_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"timeout_member": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"default_pool_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_authorities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},
			"default_certificate_authority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_certificate_authentication": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceListenerImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return nil, fmt.Errorf(errorImportListener)
	}

	projectId := parts[0]
	loadBalancerId := parts[1]
	listenerId := parts[2]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	if err := d.Set("load_balancer_id", loadBalancerId); err != nil {
		return nil, fmt.Errorf("error setting `load_balancer_id`: %s", err)
	}

	d.SetId(listenerId)
	return []*schema.ResourceData{d}, nil
}

func resourceListenerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Creating listener")
	var listenerId string
	cli := m.(*client.Client)

	retryErr := resource.RetryContext(ctx, 10*time.Minute, func() *resource.RetryError {
		listener, httpResponse, err := createListenerResource(ctx, d, cli)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryCreateListener, responseError.ErrorMessage())
				// Delay for 60 seconds before continuing
				time.Sleep(30 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorListenerCreate, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}

		value, ok := listener["uuid"]
		if ok {
			listenerId = value
		}

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutCreate)
	stateConf := &resource.StateChangeConf{
		Pending:    listenerCreating,
		Target:     listenerCreated,
		Refresh:    resourceListenerStateRefreshFunc(ctx, d, cli, listenerId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	// Wait, catching any errors
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingListenerCreate, listenerId, err))
	}

	d.SetId(listenerId)

	log.Printf("Created listener successfully")
	return resourceListenerRead(ctx, d, m)
}

func createListenerResource(ctx context.Context, d *schema.ResourceData, cli *client.Client) (map[string]string, *http.Response, error) {
	loadBalancerId := d.Get("load_balancer_id").(string)
	projectId := d.Get("project_id").(string)

	req := vloadbalancing.CreateListenerRequestV2{
		AllowedCidrs:         d.Get("allowed_cidrs").(string),
		ListenerName:         d.Get("name").(string),
		ListenerProtocol:     d.Get("protocol").(string),
		ListenerProtocolPort: int32(d.Get("protocol_port").(int)),
		TimeoutClient:        int32(d.Get("timeout_client").(int)),
		TimeoutConnection:    int32(d.Get("timeout_connection").(int)),
		TimeoutMember:        int32(d.Get("timeout_member").(int)),
		DefaultPoolId:        d.Get("default_pool_id").(string),
	}

	if v, ok := d.GetOk("default_certificate_authority"); ok {
		req.DefaultCertificateAuthority = new(string)
		*req.DefaultCertificateAuthority = v.(string)
	}
	if v, ok := d.GetOk("client_certificate_authentication"); ok {
		req.ClientCertificateAuthentication = new(string)
		*req.ClientCertificateAuthentication = v.(string)
	}

	if v, ok := d.GetOk("certificate_authorities"); ok {
		certificateAuthoritiesInterface := v.([]interface{})
		certificateAuthorities := make([]string, 0)
		for _, cert := range certificateAuthoritiesInterface {
			certificateAuthorities = append(certificateAuthorities, cert.(string))
		}
		req.CertificateAuthorities = &certificateAuthorities
	}

	if v, ok := d.GetOk("headers"); ok {
		headersSet, ok := v.(*schema.Set)
		if ok && headersSet.Len() > 0 {
			headers := make([]string, headersSet.Len())
			for i, header := range headersSet.List() {
				headers[i] = header.(string)
			}
			req.Headers = &headers
		}
	}

	listener, httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.CreateListenerUsingPOST(ctx, req, loadBalancerId, projectId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be created")
	log.Printf("-------------------------------------\n")

	return listener, httpResponse, err
}

func resourceListenerStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, listenerID string) resource.StateRefreshFunc {
	loadBalancerId := d.Get("load_balancer_id").(string)
	projectId := d.Get("project_id").(string)
	return func() (interface{}, string, error) {
		// Check the current state of the resource
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.GetListenersUsingGET(ctx, listenerID, loadBalancerId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			return resp, "", fmt.Errorf(errorCheckingListenerStatus, getResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		listener := resp.Data
		return listener, listener.ProgressStatus, nil
	}
}

func resourceListenerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Reading listener")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.GetListenersUsingGET(ctx, listenerId, loadBalancerId, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorListenerNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf(errorListenerRead, listenerId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", resp.Data.Name)
	d.Set("allowed_cidrs", resp.Data.AllowedCidrs)
	d.Set("protocol", resp.Data.Protocol)
	d.Set("protocol_port", resp.Data.ProtocolPort)
	d.Set("timeout_client", resp.Data.TimeoutClient)
	d.Set("timeout_connection", resp.Data.TimeoutConnection)
	d.Set("timeout_member", resp.Data.TimeoutMember)
	d.Set("default_pool_id", resp.Data.DefaultPoolId)
	d.Set("headers", resp.Data.Headers)
	d.Set("default_certificate_authority", resp.Data.DefaultCertificateAuthority)
	d.Set("certificate_authorities", resp.Data.CertificateAuthorities)
	d.Set("client_certificate_authentication", resp.Data.ClientCertificateAuthentication)
	d.Set("status", resp.Data.Status)

	log.Printf("Read listener successfully")
	return nil
}

func resourceListenerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if shouldSkipUpdate(d) {
		return nil
	}

	log.Printf("Updating listener")
	listenerId := d.Id()
	cli := m.(*client.Client)

	retryErr := resource.RetryContext(ctx, 10*time.Minute, func() *resource.RetryError {
		httpResponse, err := updateListener(ctx, d, cli)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryUpdateListener, listenerId, responseError.ErrorMessage())
				// Delay for 30 seconds before continuing
				time.Sleep(30 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorListenerUpdate, listenerId, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutUpdate)
	stateConf := &resource.StateChangeConf{
		Pending:    listenerUpdating,
		Target:     listenerUpdated,
		Refresh:    resourceListenerStateRefreshFunc(ctx, d, cli, d.Id()),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingListenerUpdate, listenerId, err))
	}

	log.Printf("Updated listener successfully")

	return resourceListenerRead(ctx, d, m)
}

func shouldSkipUpdate(d *schema.ResourceData) bool {
	keysToCheck := []string{"allowed_cidrs", "default_pool_id", "timeout_client",
		"timeout_connection", "timeout_member", "default_certificate_authority", "client_certificate_authentication", "headers"}
	if d.HasChanges(keysToCheck...) {
		return false
	}

	return true
}

func updateListener(ctx context.Context, d *schema.ResourceData, cli *client.Client) (*http.Response, error) {
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Id()

	req := vloadbalancing.UpdateListenerRequestV2{
		AllowedCidrs:      d.Get("allowed_cidrs").(string),
		DefaultPoolId:     d.Get("default_pool_id").(string),
		TimeoutClient:     int32(d.Get("timeout_client").(int)),
		TimeoutConnection: int32(d.Get("timeout_connection").(int)),
		TimeoutMember:     int32(d.Get("timeout_member").(int)),
	}

	if v, ok := d.GetOk("default_certificate_authority"); ok {
		req.DefaultCertificateAuthority = new(string)
		*req.DefaultCertificateAuthority = v.(string)
	}
	if v, ok := d.GetOk("client_certificate_authentication"); ok {
		req.ClientCertificate = new(string)
		*req.ClientCertificate = v.(string)
	}

	if v, ok := d.GetOk("headers"); ok {
		headersSet, ok := v.(*schema.Set)
		if ok && headersSet.Len() > 0 {
			headers := make([]string, headersSet.Len())
			for i, header := range headersSet.List() {
				headers[i] = header.(string)
			}
			req.Headers = &headers
		}
	}

	httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.UpdateListenerUsingPUT(ctx, listenerId, loadBalancerId, projectId, req)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
	log.Printf("-------------------------------------\n")
	return httpResponse, err
}

func resourceListenerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting listener")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Id()

	cli := m.(*client.Client)
	err := resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		httpResponse, _ := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.DeleteListenerUsingDELETE(ctx, loadBalancerId, listenerId, projectId)

		if checkErrorResponse(httpResponse) {
			errorResponse := fmt.Errorf(errorListenerDelete, listenerId, getResponseBody(httpResponse))
			if httpResponse.StatusCode == 500 {
				log.Printf(errorRetryDeleteListener, listenerId, errorResponse)
				return resource.RetryableError(errorResponse)
			} else {
				return resource.NonRetryableError(errorResponse)
			}
		}

		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	log.Printf("Deleted listener successfully")
	return nil
}
