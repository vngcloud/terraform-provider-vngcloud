package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/utils"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourcePool() *schema.Resource {
	return &schema.Resource{
		Create: resourcePoolCreate,
		Read:   resourcePoolRead,
		Update: resourcePoolUpdate,
		Delete: resourcePoolDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:PoolID", d.Id())
				}
				d.Set("project_id", idParts[0])
				d.SetId(idParts[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ROUND_ROBIN",
			},
			"health_check_method": {
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},
			"health_check_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check_protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"healthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"monitor_port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"stickiness": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"success_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioning_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func buildCreateMemberRequests(d *schema.ResourceData) []vserver.CreateMemberRequest {
	members := d.Get("members").([]interface{})
	createMemberRequest := make([]vserver.CreateMemberRequest, len(members))

	for i, m := range members {
		member := m.(map[string]interface{})
		var request vserver.CreateMemberRequest
		request.Backup = member["backup"].(bool)
		request.IpAddress = member["ip_address"].(string)
		request.MonitorPort = int32(member["monitor_port"].(int))
		request.Name = member["name"].(string)
		request.Port = int32(member["port"].(int))
		request.SubnetId = member["subnet_id"].(string)
		request.Weight = int32(member["weight"].(int))

		createMemberRequest[i] = request
	}
	return createMemberRequest
}

func resourcePoolCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create pool")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	req := vserver.CreatePoolRequest{
		LoadBalancerId: d.Get("load_balancer_id").(string),
		Algorithm:      d.Get("algorithm").(string),
		PoolName:       d.Get("name").(string),
		PoolProtocol:   d.Get("protocol").(string),
		Stickiness:     d.Get("stickiness").(bool),
		HealthMonitor: &vserver.CreateHealthMonitorRequest{
			HealthCheckMethod:   d.Get("health_check_method").(string),
			HealthCheckPath:     d.Get("health_check_path").(string),
			HealthCheckProtocol: d.Get("health_check_protocol").(string),
			HealthyThreshold:    int64(d.Get("healthy_threshold").(int)),
			UnhealthyThreshold:  int64(d.Get("unhealthy_threshold").(int)),
			Interval:            int64(d.Get("interval").(int)),
			SuccessCode:         d.Get("success_code").(string),
			Timeout:             int64(d.Get("timeout").(int)),
		},
		Members: buildCreateMemberRequests(d),
	}
	pool, _, err := cli.VserverClient.LoadBalancerRestControllerApi.CreatePoolUsingPOST(context.TODO(), req, projectId)

	respJSON, _ := json.Marshal(pool)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		return utils.GetErrorMessage(err)
	}

	stateConf := &resource.StateChangeConf{
		Pending: poolCreating,
		Target:  poolCreated,
		Refresh: func() (interface{}, string, error) {
			resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetPoolUsingGET(context.TODO(), pool.Data.Uuid, projectId)
			respJSON, _ := json.Marshal(resp)
			log.Printf("-------------------------------------\n")
			log.Printf("%s\n", string(respJSON))
			log.Printf("-------------------------------------\n")
			if err != nil {
				return nil, "", utils.GetErrorMessage(err)
			}
			return resp, resp.Data.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      30 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	d.SetId(pool.Data.Uuid)
	log.Printf("Create pool successfully")

	return nil
}

func resourcePoolRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("Read pool")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	pool, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetPoolUsingGET(context.TODO(), d.Id(), projectId)
	respJSON, _ := json.Marshal(pool)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		return utils.GetErrorMessage(err)
	}

	d.Set("name", pool.Data.Name)
	d.Set("load_balancer_id", pool.Data.LoadBalancerId)
	d.Set("algorithm", pool.Data.LoadBalanceMethod)
	d.Set("protocol", pool.Data.Protocol)
	var stickiness bool
	if pool.Data.SessionPersistence != 1 {
		stickiness = false
	} else {
		stickiness = true
	}
	d.Set("stickiness", stickiness)

	d.Set("health_check_method", pool.Data.HealthMonitor.HealthCheckMethod)
	d.Set("health_check_protocol", pool.Data.HealthMonitor.HealthCheckProtocol)
	d.Set("health_check_path", pool.Data.HealthMonitor.HealthCheckPath)
	d.Set("healthy_threshold", pool.Data.HealthMonitor.HealthyThreshold)
	d.Set("unhealthy_threshold", pool.Data.HealthMonitor.UnhealthyThreshold)
	d.Set("interval", pool.Data.HealthMonitor.Interval)
	d.Set("timeout", pool.Data.HealthMonitor.Timeout)
	d.Set("success_code", pool.Data.HealthMonitor.SuccessCode)
	d.Set("status", pool.Data.Status)
	d.Set("provisioning_status", pool.Data.ProvisioningStatus)
	d.Set("operating_status", pool.Data.OperatingStatus)

	err = memberRead(d, m)
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	log.Printf("Read pool successfully")
	return nil
}

func memberRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("Read member")

	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	pagingMember, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetMemberFromPoolUsingGET(context.TODO(), d.Id(), projectId, nil)
	respJSON, _ := json.Marshal(pagingMember)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		return err
	}

	log.Printf("Parse member")
	if pagingMember.ListData != nil {
		memberList := make([]interface{}, len(pagingMember.ListData))
		for i, member := range pagingMember.ListData {
			memberTemp := make(map[string]interface{})
			if member.Backup == 0 {
				memberTemp["backup"] = false
			} else {
				memberTemp["backup"] = true
			}
			memberTemp["ip_address"] = member.Address
			memberTemp["monitor_port"] = member.MonitorPort
			memberTemp["port"] = member.ProtocolPort
			memberTemp["subnet_id"] = member.SubnetId
			memberTemp["weight"] = member.Weight
			memberTemp["name"] = member.Name

			memberList[i] = memberTemp
		}
		d.Set("members", memberList)
	}

	log.Printf("Read member successfully")
	return nil
}

func resourcePoolUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Update pool")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	if d.HasChangeExcept("members") {
		req := vserver.UpdatePoolRequest{
			PoolId:     d.Id(),
			Algorithm:  d.Get("algorithm").(string),
			Stickiness: d.Get("stickiness").(bool),
			HealthMonitor: &vserver.UpdateHealthMonitorRequest{
				HealthCheckMethod:  d.Get("health_check_method").(string),
				HealthCheckPath:    d.Get("health_check_path").(string),
				HealthyThreshold:   int64(d.Get("healthy_threshold").(int)),
				UnhealthyThreshold: int64(d.Get("unhealthy_threshold").(int)),
				Interval:           int64(d.Get("interval").(int)),
				Timeout:            int64(d.Get("timeout").(int)),
				SuccessCode:        d.Get("success_code").(string),
			},
		}
		resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.UpdatePoolUsingPUT(context.TODO(), projectId, req)
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		if err != nil {
			return utils.GetErrorMessage(err)
		}

		stateConf := &resource.StateChangeConf{
			Pending: poolUpdating,
			Target:  poolCreated,
			Refresh: func() (interface{}, string, error) {
				resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetPoolUsingGET(context.TODO(), d.Id(), projectId)
				respJSON, _ := json.Marshal(resp)
				log.Printf("-------------------------------------\n")
				log.Printf("%s\n", string(respJSON))
				log.Printf("-------------------------------------\n")

				if err != nil {
					return nil, "", utils.GetErrorMessage(err)
				}

				return resp, resp.Data.Status, err
			},
			Timeout:    d.Timeout(schema.TimeoutUpdate),
			Delay:      10 * time.Second,
			MinTimeout: 10 * time.Second,
		}

		_, err = stateConf.WaitForState()
		if err != nil {
			return utils.GetErrorMessage(err)
		}
	}
	if d.HasChange("members") {
		memberUpdate(d, m)
	}

	log.Printf("Update pool successfully")
	return resourcePoolRead(d, m)
}

func memberUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Update member")

	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	createMemberRequest := buildCreateMemberRequests(d)
	updateMembersRequest := vserver.UpdateMembersRequest{
		Members: createMemberRequest,
	}

	_, err := cli.VserverClient.LoadBalancerRestControllerApi.UpdateMembersUsingPUT(context.TODO(), d.Id(), projectId, updateMembersRequest)

	if err != nil {
		log.Printf("%v\n", err)
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending: poolUpdating,
		Target:  poolCreated,
		Refresh: func() (interface{}, string, error) {
			resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetPoolUsingGET(context.TODO(), d.Id(), projectId)
			respJSON, _ := json.Marshal(resp)
			log.Printf("-------------------------------------\n")
			log.Printf("%s\n", string(respJSON))
			log.Printf("-------------------------------------\n")

			if err != nil {
				return nil, "", utils.GetErrorMessage(err)
			}

			return resp, resp.Data.Status, err
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	log.Printf("Update member successfully")
	return nil
}

func resourcePoolDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Delete pool")
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)

	req := vserver.DeletePoolRequest{
		PoolId: d.Id(),
	}
	resp, httpResp, err := cli.VserverClient.LoadBalancerRestControllerApi.DeletePoolUsingDELETE(context.TODO(), req, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if httpResp.StatusCode != 200 {
		log.Printf("status code %v\n", httpResp.StatusCode)
		log.Printf("status %v\n", httpResp.Status)
		if err != nil {
			return utils.GetErrorMessage(err)
		}
	}

	log.Printf("Delete pool successfully")

	return nil
}
