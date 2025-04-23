package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetCreate,
		Read:   resourceSubnetRead,
		Delete: resourceSubnetDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:NetworkID:SubnetID", d.Id())
				}
				projectID := idParts[0]
				networkID := idParts[1]
				subnetID := idParts[2]
				d.SetId(subnetID)
				d.Set("project_id", projectID)
				d.Set("network_id", networkID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interface_acl_policy_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_table_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}
func resourceSubnetStateRefreshFunc(cli *client.Client, networkID string, subnetID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), networkID, projectID, subnetID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		subnet := resp
		return subnet, subnet.Status, nil
	}
}
func resourceSubnetCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkID := d.Get("network_id").(string)
	subnet := vserver.CreateSubnetRequest{
		Name:   d.Get("name").(string),
		Cidr:   d.Get("cidr").(string),
		ZoneId: d.Get("zone_id").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.SubnetRestControllerApi.CreateSubnetUsingPOST1(context.TODO(), subnet, networkID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    subnetCreating,
		Target:     subnetCreated,
		Refresh:    resourceSubnetStateRefreshFunc(cli, networkID, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create subnet (%s) %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceSubnetRead(d, m)
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	subnetID := d.Id()
	networkID := d.Get("network_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), networkID, projectID, subnetID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	subnet := resp
	d.Set("name", subnet.Name)
	d.Set("cidr", subnet.Cidr)
	d.Set("zone_id", subnet.Zone.Uuid)
	return nil
}

func resourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	SubnetId := d.Id()
	NetworkId := d.Get("network_id").(string)
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.DeleteNetworkUsingDELETE2(context.TODO(), projectID, SubnetId, NetworkId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	stateConf := &resource.StateChangeConf{
		Pending:    subnetDeleting,
		Target:     subnetDeleted,
		Refresh:    resourceSubnetDeleteStateRefreshFunc(cli, NetworkId, SubnetId, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for delete subnet (%s) %s", SubnetId, err)
	}
	d.SetId("")
	return nil
}

func resourceSubnetDeleteStateRefreshFunc(cli *client.Client, networkID string, subnetID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), networkID, projectID, subnetID)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.SubnetDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		subnet := resp
		return subnet, subnet.Status, nil
	}
}
