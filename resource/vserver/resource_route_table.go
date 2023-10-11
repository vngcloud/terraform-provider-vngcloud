package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	errorRouteTableCreate         = "error creating Route Table: %s"
	errorRouteTableRead           = "error reading Route Table with Id(%s): %s"
	errorRouteTableDelete         = "error deleting Route Table (%s): %s"
	errorRouteTableUpdate         = "error update Route Table (%s): %s"
	errorWaitingRouteTableCreate  = "error waiting for Route Table (%s) to be ready: %s"
	errorWaitingRouteTableUpdate  = "error waiting for Route Table (%s) to be ready: %s"
	errorCheckingRouteTableStatus = "error checking Route Table: %s"
	errorImportRouteTable         = "import format error: to import a Route Table use the format {projectId}:{routeTableId}"
	errorRouteTableNotFound       = "Route Table with Id(%s) not found, removing from state"
	errorWaitingRouteTableDelete  = "error waiting for Route Table (%s) to be deleted: %s"
)

func ResourceRouteTable() *schema.Resource {
	return resourceRouteTableInstance()
}

func resourceRouteTableInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouteTableCreate,
		ReadContext:   resourceRouteTableRead,
		UpdateContext: resourceRouteTableUpdate,
		DeleteContext: resourceRouteTableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceRouteTableImportState,
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
			"network_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"route": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_cidr_block": {
							Type:     schema.TypeString,
							Required: true,
						},
						"target": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceRouteTableImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 2 || parts[0] == "" {
		return nil, fmt.Errorf(errorImportRouteTable)
	}

	projectId := parts[0]
	routeTableId := parts[1]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	d.SetId(routeTableId)
	return []*schema.ResourceData{d}, nil
}

func resourceRouteTableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Create route table")
	var routeTableId string
	cli := m.(*client.Client)

	routeTable, httpResponse, err := createRoutable(ctx, d, cli)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorRouteTableCreate, responseBody)
		return diag.FromErr(errResponse)
	}

	value, ok := routeTable["uuid"]
	if ok {
		routeTableId = value
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutCreate)
	stateConf := &resource.StateChangeConf{
		Pending:    routeTableCreating,
		Target:     routeTableCreated,
		Refresh:    resourceRouteTableStateRefreshFunc(ctx, d, cli, routeTableId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	// Wait, catching any errors
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingRouteTableCreate, routeTableId, err))
	}

	d.SetId(routeTableId)

	log.Printf("Created route table successfully")
	return resourceRouteTableRead(ctx, d, m)
}

func createRoutable(ctx context.Context, d *schema.ResourceData, cli *client.Client) (map[string]string, *http.Response, error) {
	projectId := d.Get("project_id").(string)
	req := vserver.CreateRouteTableRequest{
		Name:      d.Get("name").(string),
		NetworkId: d.Get("network_id").(string),
		Routes:    expandRoutes(d.Get("route").(*schema.Set).List()),
	}

	routeTable, httpResponse, err := cli.VserverClient.RouteTableControllerApi.CreateRouteTableUsingPOST1(ctx, req, projectId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be created")
	log.Printf("-------------------------------------\n")

	return routeTable, httpResponse, err
}

func expandRoutes(tfRoutes []interface{}) []vserver.RouteRequest {
	var expandedRoutes []vserver.RouteRequest

	for _, r := range tfRoutes {
		routeMap := r.(map[string]interface{})
		expandedRoute := vserver.RouteRequest{
			DestinationCidrBlock: routeMap["destination_cidr_block"].(string),
			Target:               routeMap["target"].(string),
		}

		expandedRoutes = append(expandedRoutes, expandedRoute)
	}

	return expandedRoutes
}

func resourceRouteTableStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, routeTableId string) resource.StateRefreshFunc {
	projectId := d.Get("project_id").(string)
	return func() (interface{}, string, error) {
		// Check the current state of the resource
		resp, httpResponse, _ := cli.VserverClient.RouteTableControllerApi.GetRouteTableUsingGET1(ctx, projectId, routeTableId)
		if httpResponse.StatusCode != http.StatusOK {
			return resp, "", fmt.Errorf(errorCheckingRouteTableStatus, GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		routeTable := resp.Data
		return routeTable, routeTable.Status, nil
	}
}

func resourceRouteTableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read route table")

	projectId := d.Get("project_id").(string)
	routeTableId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.RouteTableControllerApi.GetRouteTableUsingGET1(ctx, projectId, routeTableId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorRouteTableNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf(errorRouteTableRead, routeTableId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", resp.Data.Name)
	d.Set("network_id", resp.Data.NetworkId)
	d.Set("route", flattenRoutes(resp.Data.Routes))
	d.Set("status", resp.Data.Status)
	log.Printf("Read route table successfully")
	return nil
}

func flattenRoutes(routes []vserver.RoutesDto) []interface{} {
	var flattenedRoutes []interface{}

	for _, r := range routes {
		flattenedRoute := map[string]interface{}{
			"destination_cidr_block": r.DestinationCidrBlock,
			"target":                 r.Target,
		}

		flattenedRoutes = append(flattenedRoutes, flattenedRoute)
	}

	return flattenedRoutes
}

func resourceRouteTableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !d.HasChange("route") {
		return nil
	}
	log.Printf("Updating route table")
	routeTableId := d.Id()

	cli := m.(*client.Client)

	httpResponse, err := updateRouteTable(ctx, d, cli)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorRouteTableUpdate, routeTableId, responseBody)
		return diag.FromErr(errResponse)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutUpdate)
	stateConf := &resource.StateChangeConf{
		Pending:    routeTableUpdating,
		Target:     routeTableUpdated,
		Refresh:    resourceRouteTableStateRefreshFunc(ctx, d, cli, routeTableId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingRouteTableUpdate, routeTableId, err))
	}

	log.Printf("Updated Route Table successfully")
	return resourceRouteTableRead(ctx, d, m)
}

func updateRouteTable(ctx context.Context, d *schema.ResourceData, cli *client.Client) (*http.Response, error) {
	projectId := d.Get("project_id").(string)
	routeTableId := d.Id()

	req := vserver.ChangeRoutesRequest{
		Routes: expandRoutes(d.Get("route").(*schema.Set).List()),
	}

	httpResponse, err := cli.VserverClient.RouteTableControllerApi.UpdateRouteTableDetailUsingPUT(ctx, req, projectId, routeTableId)

	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", "Request Accepted !!!. Waiting for the resource to be updated")
	log.Printf("-------------------------------------\n")
	return httpResponse, err
}

func resourceRouteTableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting route table")

	projectId := d.Get("project_id").(string)
	routeTableId := d.Id()

	cli := m.(*client.Client)

	httpResponse, _ := cli.VserverClient.RouteTableControllerApi.DeleteRouteTableUsingDELETE1(ctx, projectId, routeTableId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf(errorRouteTableDelete, routeTableId, responseBody)
		return diag.FromErr(err)
	}

	// Wait for the resource to reach the desired state
	timeout := d.Timeout(schema.TimeoutDelete)
	stateConf := &resource.StateChangeConf{
		Pending:    routeTableDeleting,
		Target:     routeTableDeleted,
		Refresh:    resourceRouteTableDeleteStateRefreshFunc(ctx, d, cli, routeTableId),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf(errorWaitingRouteTableDelete, routeTableId, err))
	}

	d.SetId("")

	log.Printf("Deleted route table successfully")
	return nil
}

func resourceRouteTableDeleteStateRefreshFunc(ctx context.Context, d *schema.ResourceData, cli *client.Client, routeTableId string) resource.StateRefreshFunc {
	projectId := d.Get("project_id").(string)
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.RouteTableControllerApi.GetRouteTableUsingGET1(ctx, projectId, routeTableId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.RouteTableDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf(errorRouteTableRead, routeTableId, GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		routeTable := resp.Data
		return routeTable, "DELETING", nil
	}
}
