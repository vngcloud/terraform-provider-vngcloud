package vserver

import (
	"context"
	"encoding/json"
	"log"
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
		// Importer: &schema.ResourceImporter{
		// 	State: ,
		// },
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
			// "members": {
			// 	Type: schema.TypeList,
			// 	Elem: &schema.Schema{
			// 		Type: vserver.CreateMemberRequest,
			// 	},
			// },
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
		},
	}
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
		// Members: ,
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
		stickiness = true
	} else {
		stickiness = false
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

	log.Printf("Read pool successfully")
	return nil
}

func resourcePoolUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Update pool")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

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

	log.Printf("Update pool successfully")
	return nil
}

func resourcePoolDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Delete pool")
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)

	req := vserver.DeletePoolRequest{
		PoolId: d.Id(),
	}
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.DeletePoolUsingDELETE(context.TODO(), req, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	log.Printf("Delete pool successfully")

	return nil
}
