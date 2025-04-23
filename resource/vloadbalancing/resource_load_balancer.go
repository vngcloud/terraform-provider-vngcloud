package vloadbalancing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vloadbalancing"
)

const (
	errorLoadBalancerCreate         = "error creating LoadBalancer: %s"
	errorLoadBalancerRead           = "error reading LoadBalancer with Id(%s): %s"
	errorLoadBalancerDelete         = "error deleting LoadBalancer (%s): %s"
	errorLoadBalancerResize         = "error resize LoadBalancer (%s): %s"
	errorWaitingLoadBalancerCreate  = "error waiting for LoadBalancer (%s) to be ready: %s"
	errorWaitingLoadBalancerUpdate  = "error waiting for LoadBalancer (%s) to be ready: %s"
	errorCheckingLoadBalancerStatus = "error checking Load Balancer: %s"
	errorImportLoadBalancer         = "import format error: to import a Load Balancer use the format {projectId}:{loadBalancerId}"
	errorLoadBalancerNotFound       = "Load Balancer with Id(%s) not found, removing from state"
	errorRetryDeleteLB              = "[DEBUG] Retrying Deleting Load Balancer with Id(%s) follow error: %s"
	errorRetryLoadBalancerRead      = "[DEBUG] Retrying Load Balancer with Id(%s) follow error: %s"
	errorWaitingLoadBalancerDelete  = "error waiting for Load Balancer (%s) to be deleted: %s"
)

func ResourceLoadBalancer() *schema.Resource {
	return resourceLoadBalancerInstance()
}

func resourceLoadBalancerInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLoadBalancerCreate,
		ReadContext:   resourceLoadBalancerRead,
		UpdateContext: resourceLoadBalancerUpdate,
		DeleteContext: resourceLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceLoadBalancerImportState,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheme": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "Internet",
			},
			"subnet_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_subnet_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLoadBalancerImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 2 || parts[0] == "" {
		return nil, fmt.Errorf(errorImportLoadBalancer)
	}

	projectId := parts[0]
	loadBalancerId := parts[1]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	d.SetId(loadBalancerId)
	return []*schema.ResourceData{d}, nil
}

func resourceLoadBalancerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Create load balancer")
	var loadBalancerId string
	cli := m.(*client.Client)

	loadbalancer, httpResponse, err := createLBResource(ctx, d, cli)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorLoadBalancerCreate, responseBody)
		return diag.FromErr(errResponse)
	}

	value, ok := loadbalancer["uuid"]
	if ok {
		loadBalancerId = value
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutCreate)
	stateConf := &resource.StateChangeConf{
		Pending:    loadBalancerCreating,
		Target:     loadBalancerCreated,
		Refresh:    resourceLBStateRefreshFunc(ctx, d, cli, loadBalancerId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	// Wait, catching any errors
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingLoadBalancerCreate, loadBalancerId, err))
	}

	d.SetId(loadBalancerId)

	log.Printf("Created load balancer successfully")
	return resourceLoadBalancerRead(ctx, d, m)
}

func createLBResource(ctx context.Context, d *schema.ResourceData, cli *client.Client) (map[string]string, *http.Response, error) {
	projectId := d.Get("project_id").(string)
	req := vloadbalancing.CreateLoadBalancerRequestV2{
		Name:      d.Get("name").(string),
		PackageId: d.Get("package_id").(string),
		Scheme:    d.Get("scheme").(string),
		SubnetId:  d.Get("subnet_id").(string),
		Type_:     d.Get("type").(string),
		ZoneId:    d.Get("zone_id").(string),
	}

	loadbalancer, httpResponse, err := cli.VlbClient.LoadBalancerRestControllerV2Api.CreateLoadBalancerUsingPOST(ctx, req, projectId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be created")
	log.Printf("-------------------------------------\n")

	return loadbalancer, httpResponse, err
}

func resourceLBStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, loadBalancerID string) resource.StateRefreshFunc {
	projectId := d.Get("project_id").(string)
	return func() (interface{}, string, error) {
		// Check the current state of the resource
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerRestControllerV2Api.GetLoadBalancerUsingGET(ctx, loadBalancerID, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			return resp, "", fmt.Errorf(errorCheckingLoadBalancerStatus, getResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		loadbalancer := resp.Data
		return loadbalancer, loadbalancer.ProgressStatus, nil
	}
}

func resourceLoadBalancerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read load balancer")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VlbClient.LoadBalancerRestControllerV2Api.GetLoadBalancerUsingGET(ctx, loadBalancerId, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorLoadBalancerNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf(errorLoadBalancerRead, loadBalancerId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", resp.Data.Name)
	d.Set("package_id", resp.Data.PackageId)
	d.Set("subnet_id", resp.Data.PrivateSubnetId)
	d.Set("type", resp.Data.Type)
	d.Set("scheme", resp.Data.LoadBalancerSchema)
	d.Set("address", resp.Data.Address)
	d.Set("private_subnet_cidr", resp.Data.PrivateSubnetCidr)
	d.Set("status", resp.Data.Status)
	d.Set("zone_id", resp.Data.Zone.Uuid)
	log.Printf("Read load balancer successfully")
	return nil
}

func resourceLoadBalancerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("package_id") {
		return nil
	}
	log.Printf("Resize load balancer")
	loadBalancerId := d.Id()

	cli := m.(*client.Client)
	retryErr := resource.RetryContext(ctx, 10*time.Minute, func() *resource.RetryError {
		httpResponse, err := resizeLoadBalancer(ctx, d, cli)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryLoadBalancerRead, loadBalancerId, responseError.ErrorMessage())
				// Delay for 30 seconds before continuing
				time.Sleep(30 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorLoadBalancerResize, loadBalancerId, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	stateConf := &resource.StateChangeConf{
		Pending:    loadBalancerUpdating,
		Target:     loadBalancerUpdated,
		Refresh:    resourceLBStateRefreshFunc(ctx, d, cli, loadBalancerId),
		Timeout:    45 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingLoadBalancerUpdate, loadBalancerId, err))
	}

	log.Printf("Updated Load Balancer successfully")
	return resourceLoadBalancerRead(ctx, d, m)
}

func resizeLoadBalancer(ctx context.Context, d *schema.ResourceData, cli *client.Client) (*http.Response, error) {
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Id()

	req := vloadbalancing.ResizeLoadBalancerRequest{
		PackageId: d.Get("package_id").(string),
	}

	httpResponse, err := cli.VlbClient.LoadBalancerRestControllerV2Api.ResizeLoadBalancerUsingPUT(ctx, loadBalancerId, projectId, req)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
	log.Printf("-------------------------------------\n")
	return httpResponse, err
}
func resourceLoadBalancerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting load balancer")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Id()

	cli := m.(*client.Client)
	retryErr := resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		httpResponse, _ := cli.VlbClient.LoadBalancerRestControllerV2Api.DeleteLoadBalancerUsingDELETE(ctx, loadBalancerId, projectId)
		if checkErrorResponse(httpResponse) {
			errorResponse := fmt.Errorf(errorLoadBalancerDelete, loadBalancerId, getResponseBody(httpResponse))
			if httpResponse.StatusCode == http.StatusInternalServerError {
				log.Printf(errorRetryDeleteLB, loadBalancerId, errorResponse)
				return resource.RetryableError(errorResponse)
			} else {
				return resource.NonRetryableError(errorResponse)
			}
		}

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutDelete)
	stateConf := &resource.StateChangeConf{
		Pending:    loadBalancerDeleting,
		Target:     loadBalancerDeleted,
		Refresh:    resourceLoadBalancerDeleteStateRefreshFunc(ctx, d, cli, loadBalancerId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingLoadBalancerDelete, loadBalancerId, err))
	}

	d.SetId("")

	log.Printf("Deleted load balancer successfully")
	return nil
}

func resourceLoadBalancerDeleteStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, loadBalancerId string) resource.StateRefreshFunc {
	projectId := d.Get("project_id").(string)
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerRestControllerV2Api.GetLoadBalancerUsingGET(ctx, loadBalancerId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vloadbalancing.LoadBalancerDto{ProgressStatus: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf(errorLoadBalancerRead, loadBalancerId, getResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		lb := resp.Data
		return lb, "DELETING", nil
	}
}
