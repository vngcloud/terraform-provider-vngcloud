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

func ResourceSecgroupRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecgroupRuleCreate,
		Read:   resourceSecgroupRuleRead,
		//Update: resourceSecgroupRuleUpdate,
		Delete: resourceSecgroupRuleDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"direction": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ethertype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_range_max": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"port_range_min": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"remote_ip_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSecgroupRuleStateRefreshFunc(cli *client.Client, secgroupRuleID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET(context.TODO(), projectID, secgroupRuleID)
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
		secGroupRule := resp.SecgroupRules[0]
		return secGroupRule, secGroupRule.Status, nil
	}
}
func resourceSecgroupRuleCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secGroupRule := vserver.CreateSecurityGroupRuleRequest{
		Direction:       d.Get("direction").(string),
		Ethertype:       d.Get("ethertype").(string),
		PortRangeMax:    int32(d.Get("port_range_max").(int)),
		PortRangeMin:    int32(d.Get("port_range_min").(int)),
		Protocol:        d.Get("protocol").(string),
		RemoteIpPrefix:  d.Get("remote_ip_prefix").(string),
		SecurityGroupId: d.Get("security_group_id").(string),
		Description:     d.Get("description").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SecgroupRuleRestControllerApi.CreateSecgroupRuleUsingPOST(context.TODO(), secGroupRule, projectID)
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
		Pending:    secgroupRuleCreating,
		Target:     secgroupRuleCreated,
		Refresh:    resourceSecgroupRuleStateRefreshFunc(cli, resp.SecgroupRules[0].Id, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.SecgroupRules[0].Id, err)
	}
	d.SetId(resp.SecgroupRules[0].Id)
	return resourceSecgroupRuleRead(d, m)
}

func resourceSecgroupRuleRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	SecgroupRuleID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET(context.TODO(), projectID, SecgroupRuleID)
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
	if len(resp.SecgroupRules) == 0 {
		d.SetId("")
	}
	return nil
}

func resourceSecgroupRuleUpdate(d *schema.ResourceData, m interface{}) error {
	// if d.HasChange("size") || d.HasChange("volume_type_id") {
	// 	projectID := d.Get("project_id").(string)
	// 	resizeVolume := vSecgroupRule.ResizeVolumeRequest{
	// 		NewSize:         int32(d.Get("size").(int)),
	// 		VolumeId:        d.Id(),
	// 		NewVolumeTypeId: d.Get("volume_type_id").(string),
	// 	}
	// 	cli := m.(*client.Client)
	// 	resp, _, err := cli.VSecgroupRuleClient.VolumeRestControllerApi.ResizeVolumeUsingPUT(context.TODO(), projectID, resizeVolume)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if !resp.Success {
	// 		err := fmt.Errorf(resp.ErrorMsg)
	// 		return err
	// 	}
	// 	respJSON, _ := json.Marshal(resp)
	// 	log.Printf("-------------------------------------\n")
	// 	log.Printf("%s\n", string(respJSON))
	// 	log.Printf("-------------------------------------\n")
	// 	return nil
	// }
	return resourceSecgroupRuleRead(d, m)

}

func resourceSecgroupRuleDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteSecgroupRule := vserver.DeleteSecurityGroupRuleRequest{
		SecgroupRuleId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.SecgroupRuleRestControllerApi.DeleteSecgroupRuleUsingDELETE(context.TODO(), deleteSecgroupRule, projectID)
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
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.SecgroupRules) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.SecgroupRules[0].Status))
	})
}
