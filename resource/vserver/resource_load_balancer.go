package vserver

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/utils"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoadBalancerCreate,
		Read:   resourceLoadBalancerRead,
		// Update: resourceLoadBalancerUpdate,
		Delete: resourceLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:LoadBalancerID", d.Id())
				}
				d.Set("project_id", idParts[0])
				d.SetId(idParts[1])
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			// "algorithm": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "ROUND_ROBIN",
			// },
			// "health_check_method": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
			// "health_check_path": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "health_check_protocol": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
			// "healthy_threshold": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Default:  3,
			// },
			// "interval": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Default:  30,
			// },
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			// "members": {
			// 	Type: schema.TypeList,
			// 	Elem: &schema.Schema{
			// 		Type: vserver.CreateMemberRequest,
			// 	},
			// },
			"package_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			// "pool_name": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
			// "pool_protocol": {
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// },
			// "timeout": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Default:  5,
			// },
			"scheme": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "Internet",
				// ExactlyOneOf: []string{"Internet", "Internal"},
			},
			// "stickiness": {
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Default:  false,
			// },
			"subnet_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			// "success_code": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				// ExactlyOneOf: []string{"Layer 4", "Layer 7"},
			},
			// "unhealthy_threshold": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Default:  3,
			// },
		},
	}
}

func resourceLoadBalancerCreate(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	req := vserver.CreateLoadBalancerRequest{
		// Algorithm:              d.Get("algorithm").(string),
		// CertificateAuthorities: d.Get("certificate_authorities").([]string),
		// HealthCheckMethod:      d.Get("health_check_method").(string),
		// HealthCheckPath:        d.Get("health_check_path").(string),
		// HealthCheckProtocol:    d.Get("health_check_protocol").(string),
		// HealthyThreshold:       d.Get("healthy_threshold").(int64),
		// Interval:               d.Get("interval").(int64),
		// Members: ,
		Name:      d.Get("name").(string),
		PackageId: d.Get("package_id").(string),
		// PoolName:           d.Get("pool_name").(string),
		// PoolProtocol:       d.Get("pool_protocol").(string),
		Scheme: d.Get("scheme").(string),
		// Stickiness:         d.Get("stickiness").(bool),
		SubnetId: d.Get("subnet_id").(string),
		// SuccessCode:        d.Get("success_code").(string),
		// Timeout:            d.Get("timeout").(int64),
		Type_: d.Get("type").(string),
		// UnhealthyThreshold: d.Get("unhealthy_threshold").(int64),
	}
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	log.Printf("Create request: %v", req)
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.CreateLoadBalancerUsingPOST(context.TODO(), req, projectId)
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	log.Printf("Waiting for creation")
	stateConf := &resource.StateChangeConf{
		Pending: loadBalancerCreating,
		Target:  loadBalancerCreated,
		Refresh: func() (interface{}, string, error) {
			lb, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetLoadBalancerUsingGET(context.TODO(), resp.Data.Uuid, projectId)
			if err != nil {
				return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
			}
			return lb, lb.Data.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      30 * time.Second,
		MinTimeout: 20 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}
	log.Printf("Created successfully")
	d.SetId(resp.Data.Uuid)
	return resourceLoadBalancerRead(d, m)
}

func resourceLoadBalancerRead(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	lb, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetLoadBalancerUsingGET(context.TODO(), d.Id(), projectId)

	if err != nil {
		return utils.GetErrorMessage(err)
	}

	d.Set("project_id", lb.Data.ProjectId)
	d.Set("name", lb.Data.Name)
	d.Set("package_id", lb.Data.PackageId)
	d.Set("subnet_id", lb.Data.SubnetId)
	d.Set("type", lb.Data.Type_)
	d.Set("scheme", lb.Data.LoadBalancerSchema)

	return nil
}

// func resourceLoadBalancerUpdate(d *schema.ResourceData, m interface{}) error {
// 	return nil
// }

func resourceLoadBalancerDelete(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	req := vserver.DeleteLoadBalancerRequest{
		LoadBalancerId: d.Id(),
	}
	_, _, err := cli.VserverClient.LoadBalancerRestControllerApi.DeleteLoadBalancerUsingDELETE(context.TODO(), req, projectId)
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	return nil
}
