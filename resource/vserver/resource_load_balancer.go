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
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"scheme": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "Internet",
			},
			"subnet_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceLoadBalancerCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create load balancer")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	req := vserver.CreateLoadBalancerRequest{
		Name:      d.Get("name").(string),
		PackageId: d.Get("package_id").(string),
		Scheme:    d.Get("scheme").(string),
		SubnetId:  d.Get("subnet_id").(string),
		Type_:     d.Get("type").(string),
	}
	// log.Printf("Create request: %v", req)
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.CreateLoadBalancerUsingPOST(context.TODO(), req, projectId)
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	log.Printf("Waiting for creation")
	stateConf := &resource.StateChangeConf{
		Pending: loadBalancerCreating,
		Target:  loadBalancerCreated,
		Refresh: func() (interface{}, string, error) {
			lb, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetLoadBalancerUsingGET(context.TODO(), resp.Data.Uuid, projectId)
			respJSON, _ := json.Marshal(lb)
			log.Printf("-------------------------------------\n")
			log.Printf("%s\n", string(respJSON))
			log.Printf("-------------------------------------\n")
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
	log.Printf("Read load balancer")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	lb, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetLoadBalancerUsingGET(context.TODO(), d.Id(), projectId)

	respJSON, _ := json.Marshal(lb)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	d.Set("project_id", lb.Data.ProjectId)
	d.Set("name", lb.Data.Name)
	d.Set("package_id", lb.Data.PackageId)
	d.Set("subnet_id", lb.Data.SubnetId)
	d.Set("type", lb.Data.Type_)
	d.Set("scheme", lb.Data.LoadBalancerSchema)

	log.Printf("Read load balancer successfully")

	return nil
}

// func resourceLoadBalancerUpdate(d *schema.ResourceData, m interface{}) error {
// 	return nil
// }

func resourceLoadBalancerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Delete load balancer")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	req := vserver.DeleteLoadBalancerRequest{
		LoadBalancerId: d.Id(),
	}
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.DeleteLoadBalancerUsingDELETE(context.TODO(), req, projectId)
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	log.Printf("Delete load balancer successfully")
	return nil
}
