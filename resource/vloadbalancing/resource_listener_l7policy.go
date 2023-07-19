package vloadbalancing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vloadbalancing"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

const (
	errorL7PolicyCreate         = "error creating L7Policy: %s"
	errorL7PolicyRead           = "error reading L7Policy with Id(%s): %s"
	errorL7PolicyDelete         = "error deleting L7Policy (%s): %s"
	errorL7PolicyUpdate         = "error updating L7Policy (%s): %s"
	errorCheckingL7PolicyStatus = "error checking L7Policy: %s"
	errorWaitingL7PolicyUpdate  = "error waiting for L7Policy (%s) to be ready: %s"
	errorWaitingL7PolicyCreate  = "error waiting for L7Policy (%s) to be ready: %s"
	errorImportL7Policy         = "import format error: to import a L7Policy use the format {projectId}:{loadBalancerId}:{listenerId}:{l7PolicyId}"
	errorPolicyNotFound         = "L7Policy with Id(%s) not found, removing from state"
	errorRetryUpdateL7Policy    = "[DEBUG] Retrying Updating L7Policy with Id(%s) follow error: %s"
	errorRetryCreateL7Policy    = "[DEBUG] Retrying Creating L7Policy follow error: %s"
	errorRetryDeleteL7Policy    = "[DEBUG] Retrying Deleting Policy with Id(%s) follow error: %s"
	errorWaitingL7PolicyDelete  = "error waiting for L7Policy (%s) to be deleted: %s"
)

func ResourceListenerL7Policy() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceListenerL7PolicyRead,
		CreateContext: resourceListenerL7PolicyCreate,
		UpdateContext: resourceListenerL7PolicyUpdate,
		DeleteContext: resourceListenerL7PolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceListenerL7PolicyImportState,
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
			"listener_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compare_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keep_query_string": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"redirect_http_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redirect_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"position": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceListenerL7PolicyImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 4 || parts[0] == "" || parts[1] == "" || parts[2] == "" || parts[3] == "" {
		return nil, fmt.Errorf(errorImportL7Policy)
	}

	projectId := parts[0]
	loadBalancerId := parts[1]
	listenerId := parts[2]
	l7PolicyId := parts[3]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	if err := d.Set("load_balancer_id", loadBalancerId); err != nil {
		return nil, fmt.Errorf("error setting `load_balancer_id`: %s", err)
	}

	if err := d.Set("listener_id", listenerId); err != nil {
		return nil, fmt.Errorf("error setting `listener_id`: %s", err)
	}

	d.SetId(l7PolicyId)
	return []*schema.ResourceData{d}, nil
}

func resourceListenerL7PolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read l7 policy load balancer")

	projectId := d.Get("project_id").(string)
	l7PolicyId := d.Id()
	listenerId := d.Get("listener_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.GetL7PolicyUsingGET(ctx, l7PolicyId, listenerId, loadBalancerId, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorPolicyNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf(errorL7PolicyRead, l7PolicyId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", resp.Data.Name)
	d.Set("rule_type", resp.Data.L7Rule.RuleType)
	d.Set("action", resp.Data.Action)
	d.Set("compare_type", resp.Data.L7Rule.CompareType)
	d.Set("rule_value", resp.Data.L7Rule.RuleValue)
	d.Set("position", resp.Data.Position)
	d.Set("status", resp.Data.DisplayStatus)

	if resp.Data.RedirectPoolId == "" {
		d.Set("keep_query_string", resp.Data.KeepQueryString)
		d.Set("redirect_http_code", resp.Data.RedirectHttpCode)
		d.Set("redirect_url", resp.Data.RedirectUrl)
		d.Set("redirect_pool_id", nil)
	} else {
		d.Set("redirect_pool_id", resp.Data.RedirectPoolId)
		d.Set("keep_query_string", nil)
		d.Set("redirect_http_code", nil)
		d.Set("redirect_url", nil)
	}

	log.Printf("Read l7 policy load balancer successfully")
	return nil
}
func resourceListenerL7PolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Creating l7 policy load balancer")
	var policyId string
	cli := m.(*client.Client)

	retryErr := resource.RetryContext(ctx, 20*time.Minute, func() *resource.RetryError {
		policy, httpResponse, err := createL7Policy(ctx, d, cli)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryCreateL7Policy, responseError.ErrorMessage())
				// Delay for 90 seconds before continuing
				time.Sleep(90 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorL7PolicyCreate, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}

		value, ok := policy["uuid"]
		if ok {
			policyId = value
		}
		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutCreate)
	stateConf := &resource.StateChangeConf{
		Pending:    l7PolicyCreating,
		Target:     l7PolicyCreated,
		Refresh:    resourceL7PolicyStateRefreshFunc(ctx, d, cli, policyId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	// Wait, catching any errors
	_, err := stateConf.WaitForStateContext(context.Background())
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingL7PolicyCreate, policyId, err))
	}

	d.SetId(policyId)

	log.Printf("Created l7 policy load balancer successfully")
	return resourceListenerL7PolicyRead(ctx, d, m)
}

func createL7Policy(ctx context.Context, d *schema.ResourceData, cli *client.Client) (map[string]string, *http.Response, error) {
	listenerId := d.Get("listener_id").(string)
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	req := vloadbalancing.CreateL7PolicyRequestV2{
		Action:      d.Get("action").(string),
		CompareType: d.Get("compare_type").(string),
		Name:        d.Get("name").(string),
		Type_:       d.Get("rule_type").(string),
		Value:       d.Get("rule_value").(string),
	}

	if v, ok := d.GetOk("keep_query_string"); ok {
		req.KeepQueryString = new(bool)
		*req.KeepQueryString = v.(bool)
	}

	if v, ok := d.GetOk("redirect_pool_id"); ok {
		req.RedirectPoolId = new(string)
		*req.RedirectPoolId = v.(string)
	}

	if v, ok := d.GetOk("redirect_url"); ok {
		req.RedirectUrl = new(string)
		*req.RedirectUrl = v.(string)
	}

	if v, ok := d.GetOk("redirect_http_code"); ok {
		req.RedirectHttpCode = new(int32)
		*req.RedirectHttpCode = int32(v.(int))
	}

	policy, httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.CreatePolicyUsingPOST(ctx, req, loadBalancerId, listenerId, projectId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be created")
	log.Printf("-------------------------------------\n")

	return policy, httpResponse, err
}

func resourceL7PolicyStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, l7PolicyId string) resource.StateRefreshFunc {
	listenerId := d.Get("listener_id").(string)
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	return func() (interface{}, string, error) {
		// Check the current state of the resource
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.GetL7PolicyUsingGET(ctx, l7PolicyId, listenerId, loadBalancerId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			return resp, "", fmt.Errorf(errorCheckingL7PolicyStatus, getResponseBody(httpResponse))
		}

		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		policy := resp.Data
		return policy, policy.ProgressStatus, nil

	}
}

func resourceListenerL7PolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if shouldSkipUpdatePolicy(d) {
		return nil
	}

	log.Printf("Updating l7 policy load balancer")
	policyId := d.Id()
	cli := m.(*client.Client)

	retryErr := resource.RetryContext(ctx, 20*time.Minute, func() *resource.RetryError {
		httpResponse, err := updateL7Policy(ctx, d, cli, policyId)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryUpdateL7Policy, policyId, responseError.ErrorMessage())
				// Delay for 90 seconds before continuing
				time.Sleep(90 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorL7PolicyUpdate, policyId, responseError.ErrorMessage())
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
		Pending:    l7PolicyUpdating,
		Target:     l7PolicyUpdated,
		Refresh:    resourceL7PolicyStateRefreshFunc(ctx, d, cli, policyId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingL7PolicyUpdate, policyId, err))
	}

	log.Printf("Updated l7 policy load balancer successfully")
	return resourceListenerL7PolicyRead(ctx, d, m)
}

func shouldSkipUpdatePolicy(d *schema.ResourceData) bool {
	keysToCheck := []string{"action", "compare_type", "keep_query_string",
		"redirect_http_code", "redirect_pool_id", "redirect_url", "rule_type", "rule_value"}
	if d.HasChanges(keysToCheck...) {
		return false
	}
	return true
}

func updateL7Policy(ctx context.Context, d *schema.ResourceData, cli *client.Client, policyId string) (*http.Response, error) {
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Get("listener_id").(string)

	req := vloadbalancing.UpdateL7PolicyRequestV2{
		Action:      d.Get("action").(string),
		CompareType: d.Get("compare_type").(string),
		Type_:       d.Get("rule_type").(string),
		Value:       d.Get("rule_value").(string),
	}

	if v, ok := d.GetOk("keep_query_string"); ok {
		req.KeepQueryString = new(bool)
		*req.KeepQueryString = v.(bool)
	}

	if v, ok := d.GetOk("redirect_pool_id"); ok {
		req.RedirectPoolId = new(string)
		*req.RedirectPoolId = v.(string)
	}

	if v, ok := d.GetOk("redirect_url"); ok {
		req.RedirectUrl = new(string)
		*req.RedirectUrl = v.(string)
	}

	if v, ok := d.GetOk("redirect_http_code"); ok {
		req.RedirectHttpCode = new(int32)
		*req.RedirectHttpCode = int32(v.(int))
	}

	httpResponse, err := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.UpdateL7PolicyUsingPUT(ctx, loadBalancerId, listenerId, policyId, projectId, req)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
	log.Printf("-------------------------------------\n")
	return httpResponse, err
}

func resourceListenerL7PolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting l7 policy load balancer")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Get("listener_id").(string)
	policyId := d.Id()

	cli := m.(*client.Client)
	retryErr := resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		httpResponse, _ := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.DeletePolicyUsingDELETE(ctx, loadBalancerId, listenerId, policyId, projectId)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryDeleteL7Policy, policyId, responseError.ErrorMessage())
				// Delay for 30 seconds before continuing
				time.Sleep(30 * time.Second)
				return resource.RetryableError(responseError)
			}
			err := fmt.Errorf(errorL7PolicyDelete, policyId, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutDelete)
	stateConf := &resource.StateChangeConf{
		Pending:    l7PolicyDeleting,
		Target:     l7PolicyDeleted,
		Refresh:    resourceL7PolicyDeleteStateRefreshFunc(ctx, d, cli, policyId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingL7PolicyDelete, policyId, err))
	}
	d.SetId("")

	log.Printf("Deleted l7 policy load balancer successfully")
	return nil
}

func resourceL7PolicyDeleteStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, policyId string) resource.StateRefreshFunc {
	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	listenerId := d.Get("listener_id").(string)
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerListenerRestControllerV2Api.GetL7PolicyUsingGET(ctx, policyId, listenerId, loadBalancerId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vloadbalancing.L7Policy{ProgressStatus: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf(errorL7PolicyRead, policyId, getResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		policyData := resp.Data
		return policyData, "DELETING", nil
	}
}
