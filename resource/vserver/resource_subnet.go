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
		Update: resourceSubnetUpdate,
		Delete: resourceSubnetDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				// Read needs network_id to build the GET path, so import requires it too.
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
			"secondary_subnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr": {
							Type:     schema.TypeString,
							Required: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					// Hash by cidr + name: cidr is the backend's identity key; name is
					// included so Terraform detects renames (the backend updates the name
					// when the cidr is unchanged). uuid is excluded because it is Computed.
					return schema.HashString(m["cidr"].(string) + "|" + m["name"].(string))
				},
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
	if reqs := buildSecondarySubnetRequests(d); len(reqs) > 0 {
		if err := syncSecondarySubnets(cli, projectID, networkID, resp.Data.Uuid, d.Get("name").(string), reqs); err != nil {
			return err
		}
	}
	return resourceSubnetRead(d, m)
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	subnetID := d.Id()
	networkID := d.Get("network_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), networkID, projectID, subnetID)
	if CheckErrorResponse(httpResponse) {
		if httpResponse.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	subnet := resp
	// The backend soft-deletes subnets: GET keeps returning 200 with status DELETED,
	// so treat that the same as 404 and drop the resource from state.
	if subnet.Status == subnetDeleted[0] {
		d.SetId("")
		return nil
	}
	d.Set("name", subnet.Name)
	d.Set("cidr", subnet.Cidr)
	d.Set("zone_id", subnet.Zone.Uuid)
	secondaries := make([]map[string]interface{}, 0, len(subnet.SecondarySubnets))
	for _, s := range subnet.SecondarySubnets {
		secondaries = append(secondaries, map[string]interface{}{
			"name": s.Name,
			"cidr": s.Cidr,
			"uuid": s.Uuid,
		})
	}
	d.Set("secondary_subnet", secondaries)
	return nil
}

func resourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	SubnetId := d.Id()
	NetworkId := d.Get("network_id").(string)
	cli := m.(*client.Client)
	// The backend refuses to delete a subnet while it still has secondary subnets,
	// so remove them first by syncing an empty list before deleting the subnet.
	if set, ok := d.Get("secondary_subnet").(*schema.Set); ok && set.Len() > 0 {
		if err := syncSecondarySubnets(cli, projectID, NetworkId, SubnetId, d.Get("name").(string), []vserver.CreateSecondarySubnetRequest{}); err != nil {
			return err
		}
	}
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

// buildSecondarySubnetRequests reads the secondary_subnet blocks from state into a request list.
func buildSecondarySubnetRequests(d *schema.ResourceData) []vserver.CreateSecondarySubnetRequest {
	set := d.Get("secondary_subnet").(*schema.Set)
	reqs := make([]vserver.CreateSecondarySubnetRequest, 0, set.Len())
	for _, raw := range set.List() {
		m := raw.(map[string]interface{})
		reqs = append(reqs, vserver.CreateSecondarySubnetRequest{
			Name: m["name"].(string),
			Cidr: m["cidr"].(string),
		})
	}
	return reqs
}

// syncSecondarySubnets sends the FULL desired secondary list via PATCH; the backend diffs it by cidr.
func syncSecondarySubnets(cli *client.Client, projectID, networkID, subnetID, name string,
	reqs []vserver.CreateSecondarySubnetRequest) error {
	updateReq := vserver.UpdateSubnetRequest{
		Name:                    name,
		SecondarySubnetRequests: reqs,
	}
	_, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.EditSubnetUsingPATCH(
		context.TODO(), projectID, subnetID, updateReq, networkID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		return fmt.Errorf("request fail with errMsg : %s", responseBody)
	}
	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	projectID := d.Get("project_id").(string)
	networkID := d.Get("network_id").(string)
	subnetID := d.Id()
	if d.HasChange("secondary_subnet") || d.HasChange("name") {
		reqs := buildSecondarySubnetRequests(d)
		if err := syncSecondarySubnets(cli, projectID, networkID, subnetID, d.Get("name").(string), reqs); err != nil {
			// Keep the prior state: the PATCH was rejected, so nothing changed on the backend.
			d.Partial(true)
			return err
		}
	}
	return resourceSubnetRead(d, m)
}
