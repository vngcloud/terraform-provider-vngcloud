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

func ResourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: resourceNetworkDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:NetworkID", d.Id())
				}
				projectID := idParts[0]
				networkID := idParts[1]
				d.SetId(networkID)
				d.Set("project_id", projectID)
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
			"route_table_default_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceNetworkStateRefreshFunc(cli *client.Client, networkID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET1(context.TODO(), networkID, projectID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		network := resp
		return network, network.Status, nil
	}
}
func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	network := vserver.CreateNetworkRequest{
		Name: d.Get("name").(string),
		Cidr: d.Get("cidr").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.NetworkRestControllerApi.CreateNetworkUsingPOST1(context.TODO(), network, projectID)
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
		Pending:    networkCreating,
		Target:     networkCreated,
		Refresh:    resourceNetworkStateRefreshFunc(cli, resp.Data.Id, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create network (%s) : %s", resp.Data.Id, err)
	}
	d.SetId(resp.Data.Id)
	return resourceNetworkRead(d, m)
}

func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET1(context.TODO(), networkID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	network := resp
	d.Set("name", network.DisplayName)
	d.Set("cidr", network.Cidr)
	return nil
}

func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkID := d.Id()
	cli := m.(*client.Client)
	if d.HasChange("name") {
		networkUpdate := vserver.UpdateNetworkRequest{
			Name: d.Get("name").(string),
		}
		resp, httpResponse, _ := cli.VserverClient.NetworkRestControllerApi.EditNetworkUsingPATCH(context.TODO(), networkID, projectID, networkUpdate)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		return resourceNetworkRead(d, m)
	}
	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.NetworkRestControllerApi.DeleteNetworkUsingDELETE1(context.TODO(), d.Id(), projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	stateConf := &resource.StateChangeConf{
		Pending:    networkDeleting,
		Target:     networkDeleted,
		Refresh:    resourceNetworkDeleteStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for delete network (%s) : %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func resourceNetworkDeleteStateRefreshFunc(cli *client.Client, networkID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET1(context.TODO(), networkID, projectID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		network := resp
		return network, network.Status, nil
	}
}
