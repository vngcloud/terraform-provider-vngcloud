package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vserver"
)

func ResourceSecgroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecgroupCreate,
		Read:   resourceSecgroupRead,
		Update: resourceSecgroupUpdate,
		Delete: resourceSecgroupDelete,

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
		resp, _, err := cli.VserverClient.SecgroupRestControllerApi.GetSecgroupUsingGET(context.TODO(), projectID, secgroupID)
		if err != nil {
			return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			return nil, "", fmt.Errorf("Error describing instance: %s", resp.ErrorMsg)
		}
		secGroup := resp.Secgroups[0]
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
	resp, _, err := cli.VserverClient.SecgroupRestControllerApi.CreateSecgroupUsingPOST(context.TODO(), secgroup, projectID)
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
	stateConf := &resource.StateChangeConf{
		Pending:    secgroupCreating,
		Target:     secgroupCreated,
		Refresh:    resourceSecgroupStateRefreshFunc(cli, resp.Secgroups[0].Id, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Secgroups[0].Id, err)
	}
	d.SetId(resp.Secgroups[0].Id)
	return resourceSecgroupRead(d, m)
}

func resourceSecgroupRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secgroupID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SecgroupRestControllerApi.GetSecgroupUsingGET(context.TODO(), projectID, secgroupID)
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
	if len(resp.Secgroups) == 0 {
		d.SetId("")
		return nil
	}
	secgroup := resp.Secgroups[0]
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
			SecgroupId:  secgroupID,
		}
		cli := m.(*client.Client)
		resp, _, err := cli.VserverClient.SecgroupRestControllerApi.EditSecgroupUsingPUT(context.TODO(), secgroupUpdate, projectID)
		if err != nil {
			return err
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			err := fmt.Errorf(resp.ErrorMsg)
			return err
		}
	}
	return resourceSecgroupRead(d, m)

}

func resourceSecgroupDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteSecgroup := vserver.DeleteSecurityGroupRequest{
		SecgroupId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SecgroupRestControllerApi.DeleteSecgroupUsingDELETE(context.TODO(), deleteSecgroup, projectID)
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
	return resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.SecgroupRestControllerApi.GetSecgroupUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Secgroups) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Secgroups[0].Status))
	})
}
