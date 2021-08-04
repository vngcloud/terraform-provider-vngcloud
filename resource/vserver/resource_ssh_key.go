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
	resp, _, err := cli.VserverClient.SshKeyRestControllerApi.ImportSSHKeyUsingPOST(context.TODO(), sshKey, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	d.SetId(resp.SshKeys[0].Id)
	return resourceSSHKeyRead(d, m)
}

func resourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	sshKeyID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SshKeyRestControllerApi.GetSSHKeyUsingGET(context.TODO(), projectID, sshKeyID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	if len(resp.SshKeys) == 0 {
		d.SetId("")
		return nil
	}
	sshKey := resp.SshKeys[0]
	d.Set("name", sshKey.Name)
	d.Set("public_key", sshKey.PubKey)
	return nil
}

func resourceSSHKeyDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteSSHKey := vserver.SdnSshKeyDeleteRequest{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SshKeyRestControllerApi.DeleteSSHKeyUsingDELETE(context.TODO(), deleteSSHKey, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	return resourceSSHKeyRead(d, m)
}
