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

func ResourceServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerGroupCreate,
		Read:   resourceServerGroupRead,
		Update: resourceServerGroupUpdate,
		Delete: resourceServerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:ServerGroupId", d.Id())
				}
				projectID := idParts[0]
				serverGroupId := idParts[1]
				d.SetId(serverGroupId)
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
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServerGroupCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	createServerGroupRequest := vserver.CreateServerGroupRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		PolicyId:    d.Get("policy_id").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.ServerGroupRestControllerApi.CreateServerGroupUsingPOST1(context.TODO(), createServerGroupRequest, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	d.SetId(resp.Data.Uuid)
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("description") {
		projectID := d.Get("project_id").(string)
		updateServerGroupRequest := vserver.UpdateServerGroupRequestV2{
			Description: d.Get("description").(string),
		}
		cli := m.(*client.Client)
		resp, httpResponse, _ := cli.VserverClient.ServerGroupRestControllerApi.UpdateServerGroupUsingPUT1(context.TODO(), projectID, d.Id(), updateServerGroupRequest)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		d.SetId(resp.Data.Uuid)
	}
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	serverGroupId := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.ServerGroupRestControllerApi.GetServerGroupUsingGET1(context.TODO(), projectID, serverGroupId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	serverGroup := resp.Data
	d.Set("name", serverGroup.Name)
	d.Set("description", serverGroup.Description)
	d.Set("policy_id", serverGroup.PolicyId)
	return nil
}

func resourceServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.ServerGroupRestControllerApi.DeleteServerGroupUsingDELETE1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	d.SetId("")
	return nil
}
