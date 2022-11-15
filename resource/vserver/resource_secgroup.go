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

func ResourceSecgroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecgroupCreate,
		Read:   resourceSecgroupRead,
		Update: resourceSecgroupUpdate,
		Delete: resourceSecgroupDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:SecgroupID", d.Id())
				}
				projectID := idParts[0]
				secgroupID := idParts[1]
				d.SetId(secgroupID)
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSecgroupStateRefreshFunc(cli *client.Client, secgroupID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.SecgroupRestControllerApi.GetSecgroupUsingGET1(context.TODO(), projectID, secgroupID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		secGroup := resp.Data
		return secGroup, secGroup.Status, nil
	}
}
func resourceSecgroupCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secgroup := vserver.CreateSecurityGroupRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.SecgroupRestControllerApi.CreateSecgroupUsingPOST1(context.TODO(), secgroup, projectID)
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
		Pending:    secgroupCreating,
		Target:     secgroupCreated,
		Refresh:    resourceSecgroupStateRefreshFunc(cli, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create security group (%s) %s", resp.Data.Id, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceSecgroupRead(d, m)
}

func resourceSecgroupRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secgroupID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SecgroupRestControllerApi.GetSecgroupUsingGET1(context.TODO(), projectID, secgroupID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	secgroup := resp.Data
	d.Set("name", secgroup.Name)
	d.Set("description", secgroup.Description)
	return nil
}

func resourceSecgroupUpdate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secgroupID := d.Id()
	if d.HasChange("name") || d.HasChange("description") {
		secgroupUpdate := vserver.EditSecurityGroupRequest{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		}
		cli := m.(*client.Client)
		resp, httpResponse, _ := cli.VserverClient.SecgroupRestControllerApi.UpdateSecgroupUsingPUT1(context.TODO(), secgroupUpdate, projectID, secgroupID)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
	}
	return resourceSecgroupRead(d, m)

}

func resourceSecgroupDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.SecgroupRestControllerApi.DeleteSecgroupUsingDELETE1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	d.SetId("")
	return nil
}
