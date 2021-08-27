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
	resp, _, err := cli.VserverClient.ServerGroupRestControllerApi.CreateServerGroupUsingPOST(context.TODO(), createServerGroupRequest, projectID)
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
	d.SetId(resp.ServerGroups[0].Uuid)
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("description") {
		projectID := d.Get("project_id").(string)
		updateServerGroupRequest := vserver.UpdateServerGroupRequest{
			Description:   d.Get("description").(string),
			ServerGroupId: d.Id(),
		}
		cli := m.(*client.Client)
		resp, _, err := cli.VserverClient.ServerGroupRestControllerApi.UpdateServerGroupUsingPUT(context.TODO(), projectID, updateServerGroupRequest)
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
		d.SetId(resp.ServerGroups[0].Uuid)
	}
	return resourceServerGroupRead(d, m)
}

func resourceServerGroupRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	serverGroupId := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerGroupRestControllerApi.GetServerGroupUsingGET(context.TODO(), projectID, serverGroupId)
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
	if len(resp.ServerGroups) == 0 {
		d.SetId("")
		return nil
	}
	serverGroup := resp.ServerGroups[0]
	d.Set("name", serverGroup.Name)
	d.Set("description", serverGroup.Description)
	d.Set("policy_id", serverGroup.PolicyId)
	return nil
}

func resourceServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteServerGroup := vserver.DeleteServerGroupRequest{
		ServerGroupId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerGroupRestControllerApi.DeleteServerGroupUsingDELETE(context.TODO(), deleteServerGroup, projectID)
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
	return resourceServerGroupRead(d, m)
}
