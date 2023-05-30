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
	errorPoolCreate         = "error creating Pool: %s"
	errorPoolRead           = "error reading Pool with Id(%s): %s"
	errorPoolMemberRead     = "error reading Pool Member with Pool Id(%s): %s"
	errorHealthMonitorRead  = "error reading Health Monitor with Pool Id(%s): %s"
	errorPoolDelete         = "error deleting Pool with Id(%s): %s"
	errorPoolUpdate         = "error updating Pool (%s): %s"
	errorMemberUpdate       = "error updating Pool Member with Pool Id(%s): %s"
	errorWaitingPoolUpdate  = "error waiting for Pool (%s) to be ready: %s"
	errorWaitingPoolCreate  = "error waiting for Pool (%s) to be ready: %s"
	errorCheckingPoolStatus = "error checking Pool: %s"
	errorImportPool         = "import format error: to import a Pool use the format {projectId}:{loadBalancerId}:{poolId}"
	errorPoolNotFound       = "Pool with Id(%s) not found, removing from state"
	errorRetryDeletePool    = "[DEBUG] Retrying Deleting Pool with Id(%s) follow error: %s"
	errorRetryCreatePool    = "[DEBUG] Retrying Creating Pool follow error: %s"
	errorRetryUpdatePool    = "[DEBUG] Retrying Updating Pool with Id(%s) follow error: %s"
)

func ResourcePool() *schema.Resource {
	return resourcePoolInstance()
}

func resourcePoolInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePoolCreate,
		ReadContext:   resourcePoolRead,
		UpdateContext: resourcePoolUpdate,
		DeleteContext: resourcePoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePoolImport,
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
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stickiness": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tls_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ROUND_ROBIN",
			},
			"health_monitor": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check_method": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"health_check_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"health_check_protocol": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"healthy_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3,
						},
						"unhealthy_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3,
						},
						"interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  30,
						},
						"timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
						"success_code": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"monitor_port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourcePoolImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return nil, fmt.Errorf(errorImportPool)
	}

	projectId := parts[0]
	loadBalancerId := parts[1]
	poolId := parts[2]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	if err := d.Set("load_balancer_id", loadBalancerId); err != nil {
		return nil, fmt.Errorf("error setting `load_balancer_id`: %s", err)
	}

	d.SetId(poolId)
	return []*schema.ResourceData{d}, nil
}

func resourcePoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Creating pool")
	var poolId string

	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	loadBalancerId := d.Get("load_balancer_id").(string)

	retryErr := resource.RetryContext(ctx, 10*time.Minute, func() *resource.RetryError {
		pool, httpResponse, err := createPoolResource(ctx, d, cli, loadBalancerId, projectId)

		if checkErrorResponse(httpResponse) {
			responseError := parseErrorResponse(httpResponse).(*ResponseError)
			if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
				log.Printf(errorRetryCreatePool, responseError.ErrorMessage())
				// Delay for 30 seconds before continuing
				time.Sleep(30 * time.Second)
				return resource.RetryableError(responseError)
			}
			err = fmt.Errorf(errorPoolCreate, responseError.ErrorMessage())
			return resource.NonRetryableError(err)
		}

		value, ok := pool["uuid"]
		if ok {
			poolId = value
		}

		return nil
	})

	if retryErr != nil {
		return diag.FromErr(retryErr)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutCreate)
	stateConf := &resource.StateChangeConf{
		Pending:    poolCreating,
		Target:     poolCreated,
		Refresh:    resourcePoolStateRefreshFunc(ctx, cli, poolId, loadBalancerId, projectId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	// Wait, catching any errors
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingPoolCreate, poolId, err))
	}

	d.SetId(poolId)

	log.Printf("Created pool successfully")
	return resourcePoolRead(ctx, d, m)
}

func createPoolResource(ctx context.Context, d *schema.ResourceData, cli *client.Client, loadBalancerId string, projectId string) (map[string]string, *http.Response, error) {
	req := vloadbalancing.CreatePoolRequestV2{
		Algorithm:     d.Get("algorithm").(string),
		PoolName:      d.Get("name").(string),
		PoolProtocol:  d.Get("protocol").(string),
		HealthMonitor: expandHealthMonitorForCreating(d.Get("health_monitor").([]interface{})),
		Members:       buildCreateMemberRequests(d),
	}

	if v, ok := d.GetOkExists("stickiness"); ok {
		req.Stickiness = new(bool)
		*req.Stickiness = v.(bool)
	}
	if v, ok := d.GetOkExists("tls_encryption"); ok {
		req.TlsEncryption = new(bool)
		*req.TlsEncryption = v.(bool)
	}

	pool, httpResponse, err := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.CreatePoolUsingPOST(ctx, req, loadBalancerId, projectId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be created")
	log.Printf("-------------------------------------\n")

	return pool, httpResponse, err
}

func buildCreateMemberRequests(d *schema.ResourceData) []vloadbalancing.CreateMemberRequest {
	members := d.Get("members").([]interface{})
	createMemberRequest := make([]vloadbalancing.CreateMemberRequest, len(members))

	for i, m := range members {
		member := m.(map[string]interface{})
		var request vloadbalancing.CreateMemberRequest
		request.Backup = member["backup"].(bool)
		request.IpAddress = member["ip_address"].(string)
		request.MonitorPort = int32(member["monitor_port"].(int))
		request.Name = member["name"].(string)
		request.Port = int32(member["port"].(int))
		request.Weight = int32(member["weight"].(int))

		createMemberRequest[i] = request
	}
	return createMemberRequest
}

func expandHealthMonitorForCreating(vHealthMonitor []interface{}) *vloadbalancing.CreateHealthMonitorRequest {
	if len(vHealthMonitor) == 0 || vHealthMonitor[0] == nil {
		return nil
	}
	tfMap, ok := vHealthMonitor[0].(map[string]interface{})
	if !ok {
		return nil
	}

	createHealthMonitor := &vloadbalancing.CreateHealthMonitorRequest{}
	if v, ok := tfMap["health_check_method"]; ok && v != "" {
		createHealthMonitor.HealthCheckMethod = new(string)
		*createHealthMonitor.HealthCheckMethod = v.(string)
	}
	if v, ok := tfMap["health_check_path"]; ok && v != "" {
		createHealthMonitor.HealthCheckPath = new(string)
		*createHealthMonitor.HealthCheckPath = v.(string)
	}
	if vHealthCheckProtocol, ok := tfMap["health_check_protocol"]; ok {
		createHealthMonitor.HealthCheckProtocol = vHealthCheckProtocol.(string)
	}
	if v, ok := tfMap["healthy_threshold"]; ok {
		createHealthMonitor.HealthyThreshold = int64(v.(int))
	}
	if vUnhealthyThreshold, ok := tfMap["unhealthy_threshold"]; ok {
		createHealthMonitor.UnhealthyThreshold = int64(vUnhealthyThreshold.(int))
	}
	if v, ok := tfMap["interval"]; ok {
		createHealthMonitor.Interval = int64(v.(int))
	}
	if v, ok := tfMap["timeout"]; ok {
		createHealthMonitor.Timeout = int64(v.(int))
	}
	if v, ok := tfMap["success_code"]; ok && v != "" {
		createHealthMonitor.SuccessCode = new(string)
		*createHealthMonitor.SuccessCode = v.(string)
	}

	if v, ok := tfMap["http_version"]; ok && v != "" {
		createHealthMonitor.HttpVersion = new(string)
		*createHealthMonitor.HttpVersion = v.(string)
	}

	if v, ok := tfMap["domain_name"]; ok && v != "" {
		createHealthMonitor.DomainName = new(string)
		*createHealthMonitor.DomainName = v.(string)
	}

	return createHealthMonitor
}

func resourcePoolStateRefreshFunc(ctx context.Context, cli *client.Client, poolId string, loadBalancerID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		// Check the current state of the resource
		resp, httpResponse, _ := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.GetPoolUsingGET(ctx, loadBalancerID, poolId, projectID)

		if httpResponse.StatusCode != http.StatusOK {
			return resp, "", fmt.Errorf(errorCheckingPoolStatus, getResponseBody(httpResponse))
		}

		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		pool := resp.Data
		return pool, pool.ProgressStatus, nil
	}
}

func resourcePoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Reading pool")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	poolId := d.Id()

	cli := m.(*client.Client)

	pool, httpResponse, err := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.GetPoolUsingGET(ctx, loadBalancerId, poolId, projectId)

	respPoolJSON, _ := json.Marshal(pool)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respPoolJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] Geting Pool Information failed:  %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorPoolNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf(errorPoolRead, poolId, responseBody)
		return diag.FromErr(err)
	}

	healthCheck, getHealthCheckHttpResponse, err1 := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.GetHealthMonitorFromPoolUsingGET(ctx, loadBalancerId, poolId, projectId)

	respHealthCheckJSON, _ := json.Marshal(healthCheck)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respHealthCheckJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] Geting Pool Information failed:  %s\n", err1)
	}

	if checkErrorResponse(getHealthCheckHttpResponse) {
		responseBody := getResponseBody(getHealthCheckHttpResponse)
		err := fmt.Errorf(errorHealthMonitorRead, poolId, responseBody)
		return diag.FromErr(err)
	}

	members, getMemberHttpResponse, err2 := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.GetMembersFromPoolUsingGET(ctx, poolId, projectId)

	respMemberJSON, _ := json.Marshal(members)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respMemberJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] Geting Pool Information failed:  %s\n", err2)
	}

	if checkErrorResponse(getMemberHttpResponse) {
		responseBody := getResponseBody(getMemberHttpResponse)
		err := fmt.Errorf(errorPoolMemberRead, poolId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", pool.Data.Name)
	d.Set("protocol", pool.Data.Protocol)
	d.Set("stickiness", pool.Data.Stickiness)
	d.Set("tls_encryption", pool.Data.TlsEncryption)
	d.Set("algorithm", pool.Data.LoadBalanceMethod)
	d.Set("status", pool.Data.Status)
	d.Set("health_monitor", flattenHealthMonitor(healthCheck.Data))

	if members.Data != nil {
		memberList := make([]interface{}, len(members.Data))
		for i, member := range members.Data {
			memberList[i] = createMemberMap(member)
		}
		d.Set("members", memberList)
	}

	log.Printf("Read pool successfully")
	return nil
}

func flattenHealthMonitor(healthMonitor *vloadbalancing.HealthMonitor) []interface{} {
	if healthMonitor == nil {
		return []interface{}{}
	}

	return []interface{}{map[string]interface{}{
		"health_check_method":   healthMonitor.HealthCheckMethod,
		"health_check_protocol": healthMonitor.HealthCheckProtocol,
		"health_check_path":     healthMonitor.HealthCheckPath,
		"healthy_threshold":     healthMonitor.HealthyThreshold,
		"unhealthy_threshold":   healthMonitor.UnhealthyThreshold,
		"interval":              healthMonitor.Interval,
		"timeout":               healthMonitor.Timeout,
		"success_code":          healthMonitor.SuccessCode,
		"domain_name":           healthMonitor.DomainName,
		"http_version":          healthMonitor.HttpVersion,
	}}
}

func createMemberMap(member vloadbalancing.Member) map[string]interface{} {
	memberTemp := make(map[string]interface{})
	memberTemp["backup"] = member.Backup
	memberTemp["ip_address"] = member.Address
	memberTemp["monitor_port"] = member.MonitorPort
	memberTemp["port"] = member.ProtocolPort
	memberTemp["weight"] = member.Weight
	memberTemp["name"] = member.Name
	memberTemp["status"] = member.Status

	return memberTemp
}

func resourcePoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Updating pool")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	poolId := d.Id()

	cli := m.(*client.Client)

	if d.HasChangeExcept("members") {
		req := vloadbalancing.UpdatePoolRequestV2{
			Algorithm:     d.Get("algorithm").(string),
			HealthMonitor: expandHealthMonitorForUpdating(d.Get("health_monitor").([]interface{})),
		}

		if v, ok := d.GetOkExists("stickiness"); ok {
			req.Stickiness = new(bool)
			*req.Stickiness = v.(bool)
		}
		if v, ok := d.GetOkExists("tls_encryption"); ok {
			req.TlsEncryption = new(bool)
			*req.TlsEncryption = v.(bool)
		}

		retryError := resource.RetryContext(ctx, 10*time.Minute, func() *resource.RetryError {
			httpResponse, err := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.UpdatePoolUsingPUT(ctx, loadBalancerId, poolId, projectId, req)

			if checkErrorResponse(httpResponse) {
				responseError := parseErrorResponse(httpResponse).(*ResponseError)
				if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
					log.Printf(errorRetryUpdatePool, poolId, responseError.ErrorMessage())
					// Delay for 30 seconds before continuing
					time.Sleep(30 * time.Second)
					return resource.RetryableError(responseError)
				}
				err = fmt.Errorf(errorPoolUpdate, poolId, responseError.ErrorMessage())
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if retryError != nil {
			return diag.FromErr(retryError)
		}

		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
		log.Printf("-------------------------------------\n")

		diagnostics, isError := waitingUpdatePool(ctx, d, cli, loadBalancerId, projectId, poolId)
		if isError {
			return diagnostics
		}
	}

	if d.HasChange("members") {
		log.Printf("Updating member of pool")

		createMemberRequest := buildCreateMemberRequests(d)
		updateMembersRequest := vloadbalancing.UpdateMembersRequestV2{
			Members: createMemberRequest,
		}

		retryErr := resource.RetryContext(ctx, 20*time.Minute, func() *resource.RetryError {
			httpResponse, err := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.UpdateMembersUsingPUT(ctx, loadBalancerId, poolId, projectId, updateMembersRequest)

			if err != nil {
				log.Printf("[ERROR] %s\n", err)
			}

			if checkErrorResponse(httpResponse) {
				responseError := parseErrorResponse(httpResponse).(*ResponseError)
				if errorCodeEquals(responseError, ErrorCodeLoadBalancerNotReady) {
					// Delay for 60 seconds before continuing
					time.Sleep(60 * time.Second)
					log.Printf(errorRetryUpdatePool, poolId, responseError.ErrorMessage())
					return resource.RetryableError(responseError)
				}
				err = fmt.Errorf(errorMemberUpdate, poolId, responseError.ErrorMessage())
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if retryErr != nil {
			return diag.FromErr(retryErr)
		}

		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
		log.Printf("-------------------------------------\n")

		diagnostics, isError := waitingUpdatePool(ctx, d, cli, loadBalancerId, projectId, poolId)
		if isError {
			return diagnostics
		}
		log.Printf("Updated member of pool successfully")
	}

	log.Printf("Updated pool successfully")
	return resourcePoolRead(ctx, d, m)
}

func waitingUpdatePool(ctx context.Context, d *schema.ResourceData, cli *client.Client, loadBalancerId string, projectId string, poolId string) (diag.Diagnostics, bool) {
	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutUpdate)
	stateConf := &resource.StateChangeConf{
		Pending:    poolUpdating,
		Target:     poolUpdated,
		Refresh:    resourcePoolStateRefreshFunc(ctx, cli, d.Id(), loadBalancerId, projectId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingPoolUpdate, poolId, err)), true
	}
	return nil, false
}

func expandHealthMonitorForUpdating(vHealthMonitor []interface{}) *vloadbalancing.UpdateHealthMonitorRequest {

	if len(vHealthMonitor) == 0 || vHealthMonitor[0] == nil {
		return nil
	}
	tfMap, ok := vHealthMonitor[0].(map[string]interface{})
	if !ok {
		return nil
	}

	updateHealthMonitor := &vloadbalancing.UpdateHealthMonitorRequest{}
	if v, ok := tfMap["health_check_method"]; ok && v != "" {
		updateHealthMonitor.HealthCheckMethod = new(string)
		*updateHealthMonitor.HealthCheckMethod = v.(string)
	}
	if v, ok := tfMap["health_check_path"]; ok && v != "" {
		updateHealthMonitor.HealthCheckPath = new(string)
		*updateHealthMonitor.HealthCheckPath = v.(string)
	}
	if v, ok := tfMap["healthy_threshold"]; ok {
		updateHealthMonitor.HealthyThreshold = int64(v.(int))
	}
	if vUnhealthyThreshold, ok := tfMap["unhealthy_threshold"]; ok {
		updateHealthMonitor.UnhealthyThreshold = int64(vUnhealthyThreshold.(int))
	}
	if v, ok := tfMap["interval"]; ok {
		updateHealthMonitor.Interval = int64(v.(int))
	}
	if v, ok := tfMap["timeout"]; ok {
		updateHealthMonitor.Timeout = int64(v.(int))
	}
	if v, ok := tfMap["success_code"]; ok && v != "" {
		updateHealthMonitor.SuccessCode = new(string)
		*updateHealthMonitor.SuccessCode = v.(string)
	}

	if v, ok := tfMap["http_version"]; ok && v != "" {
		updateHealthMonitor.HttpVersion = new(string)
		*updateHealthMonitor.HttpVersion = v.(string)
	}

	if v, ok := tfMap["domain_name"]; ok && v != "" {
		updateHealthMonitor.DomainName = new(string)
		*updateHealthMonitor.DomainName = v.(string)
	}
	return updateHealthMonitor
}

func resourcePoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting pool")

	projectId := d.Get("project_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)
	poolId := d.Id()

	cli := m.(*client.Client)
	err := resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		httpResponse, _ := cli.VlbClient.LoadBalancerPoolRestControllerV2Api.DeletePoolUsingDELETE(
			ctx,
			loadBalancerId,
			poolId,
			projectId,
		)

		if checkErrorResponse(httpResponse) {
			errorResponse := fmt.Errorf(errorPoolDelete, poolId, getResponseBody(httpResponse))
			if httpResponse.StatusCode == 500 {
				log.Printf(errorRetryDeletePool, poolId, errorResponse)
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

	log.Printf("Deleted pool successfully")
	return nil
}
