package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfaceCreate,
		Delete: resourceNetworkInterfaceDelete,
		Read:   resourceNetworkInterfaceRead,
		Update: resourceNetworkInterfaceUpdate,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:networkInterfaceId", d.Id())
				}
				projectID := idParts[0]
				networkInterfaceId := idParts[1]
				d.SetId(networkInterfaceId)
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
			"vpc_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_name": {
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

func resourceNetworkInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	createNetworkInterface := vserver.CreateNetworkInterfaceRequest{
		Name:   d.Get("name").(string),
		ZoneId: d.Get("zone_id").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.NetworkInterfaceRestControllerApi.CreateNetworkInterfaceElasticUsingPOST(context.TODO(), createNetworkInterface, projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceNetworkInterfaceRead(d, m)
}

func resourceNetworkInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)

	httpResponse, err := cli.VserverClient.NetworkInterfaceRestControllerApi.DeleteNetworkInterfaceElasticUsingDELETE(context.TODO(), d.Id(), projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be deleted: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func resourceNetworkInterfaceRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceId := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceRestControllerApi.GetNetworkInterfaceElasticUsingGET(context.TODO(), networkInterfaceId, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	networkInterfaceElastic := resp.Data
	d.Set("name", networkInterfaceElastic.Name)
	d.Set("vpc_name", networkInterfaceElastic.VpcName)
	d.Set("ip", networkInterfaceElastic.Ip)
	d.Set("server_name", networkInterfaceElastic.ServerName)
	d.Set("zone_id", networkInterfaceElastic.Zone.Uuid)
	return nil
}

func resourceNetworkInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("name") {
		projectID := d.Get("project_id").(string)
		renameNetworkInterface := vserver.RenameNetworkInterfaceRequest{
			Name: d.Get("name").(string),
		}

		cli := m.(*client.Client)
		resp, httpResponse, err := cli.VserverClient.NetworkInterfaceRestControllerApi.RenameNetworkInterfaceElasticUsingPUT(context.TODO(), d.Id(), projectID, renameNetworkInterface)

		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			oldName, _ := d.GetChange("name")
			d.Set("name", oldName)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if err != nil {
			return fmt.Errorf("Error waiting for instance (%s) to be updated: %s", resp.Data.Uuid, err)
		}
		return nil
	}
	return resourceNetworkInterfaceRead(d, m)
}
