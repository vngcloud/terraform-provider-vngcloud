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

func ResourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSHKeyCreate,
		Read:   resourceSSHKeyRead,
		//Update: resourceSSHKeyUpdate,
		Delete: resourceSSHKeyDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:SshKeyID", d.Id())
				}
				projectID := idParts[0]
				sshKeyID := idParts[1]
				d.SetId(sshKeyID)
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
			"public_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSHKeyCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	sshKey := vserver.ImportSshKeyRequest{
		Name:   d.Get("name").(string),
		PubKey: d.Get("public_key").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SshKeyRestControllerApi.ImportSSHKeyUsingPOST1(context.TODO(), sshKey, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	d.SetId(resp.Data.Id)
	return resourceSSHKeyRead(d, m)
}

func resourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	sshKeyID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SshKeyRestControllerApi.GetSSHKeyUsingGET1(context.TODO(), projectID, sshKeyID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	sshKey := resp.Data
	d.Set("name", sshKey.Name)
	d.Set("public_key", sshKey.PubKey)
	return nil
}

func resourceSSHKeyDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.SshKeyRestControllerApi.DeleteSSHKeyUsingDELETE1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	d.SetId("")
	return nil
}
