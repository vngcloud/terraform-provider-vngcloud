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

func ResourceSecgroupRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecgroupRuleCreate,
		Read:   resourceSecgroupRuleRead,
		Delete: resourceSecgroupRuleDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:SecgroupRuleID", d.Id())
				}
				projectID := idParts[0]
				secgroupRuleID := idParts[1]
				d.SetId(secgroupRuleID)
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

func resourceSecgroupRuleStateRefreshFunc(cli *client.Client, secgroupID string, secgroupRuleID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET1(context.TODO(), projectID, secgroupRuleID, secgroupID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		secGroupRule := resp.Data[0]
		return secGroupRule, secGroupRule.Status, nil
	}
}
func resourceSecgroupRuleCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	secGroupRule := vserver.CreateSecurityGroupRuleRequest{
		Direction:       d.Get("direction").(string),
		EtherType:       d.Get("ethertype").(string),
		PortRangeMax:    int32(d.Get("port_range_max").(int)),
		PortRangeMin:    int32(d.Get("port_range_min").(int)),
		Protocol:        d.Get("protocol").(string),
		RemoteIpPrefix:  d.Get("remote_ip_prefix").(string),
		SecurityGroupId: d.Get("security_group_id").(string),
		Description:     d.Get("description").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SecgroupRuleRestControllerApi.CreateSecgroupRuleUsingPOST1(context.TODO(), secGroupRule, secGroupRule.SecurityGroupId, projectID)
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
	return resourceSecgroupRuleRead(d, m)
}

func resourceSecgroupRuleRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	SecgroupRuleID := d.Id()
	SecurityGroupId := d.Get("security_group_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET1(context.TODO(), projectID, SecgroupRuleID, SecurityGroupId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	secGroupRule := resp.Data[0]
	d.Set("direction", secGroupRule.Direction)
	d.Set("ethertype", secGroupRule.EtherType)
	d.Set("port_range_max", int(secGroupRule.PortRangeMax))
	d.Set("port_range_min", int(secGroupRule.PortRangeMin))
	d.Set("protocol", secGroupRule.Protocol)
	d.Set("remote_ip_prefix", secGroupRule.RemoteIpPrefix)
	d.Set("description", secGroupRule.Description)
	return nil
}

func resourceSecgroupRuleDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	SecgroupRuleId := d.Id()
	SecurityGroupId := d.Get("security_group_id").(string)
	cli := m.(*client.Client)
	cli.VserverClient.SecgroupRuleRestControllerApi.DeleteSecgroupRuleUsingDELETE1(context.TODO(), projectID, SecgroupRuleId, SecurityGroupId)
	return resource.Retry(3*time.Minute, func() *resource.RetryError {
		_, httpResponse, _ := cli.VserverClient.SecgroupRuleRestControllerApi.GetSecgroupRuleUsingGET1(context.TODO(), projectID, SecgroupRuleId, SecurityGroupId)
		if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		httpResponse, _ = cli.VserverClient.SecgroupRuleRestControllerApi.DeleteSecgroupRuleUsingDELETE1(context.TODO(), projectID, SecgroupRuleId, SecurityGroupId)
		if CheckErrorResponse(httpResponse) {
			return resource.RetryableError(fmt.Errorf("Error when refreshing security group rule state: %s", GetResponseBody(httpResponse)))
		} else {
			d.SetId("")
			return nil
		}
	})
}
