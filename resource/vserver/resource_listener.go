package vserver

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/utils"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceListener() *schema.Resource {
	return &schema.Resource{
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
			// "allowed_cidrs": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Default:  "0.0.0.0/0",
			// },
			"listener_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listener_protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"listener_protocol_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"timeout_client": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"timeout_connection": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"timeout_member": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			// "certificate_authorities": {
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// },
		},
	}
}

func resourceListenerCreate(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	req := vserver.CreateListenerRequest{
		ListenerName:         d.Get("listener_name").(string),
		ListenerProtocol:     d.Get("listener_protocol").(string),
		ListenerProtocolPort: d.Get("listener_protocol_port").(int32),
		LoadBalancerId:       d.Get("load_balancer_id").(string),
		ProjectId:            d.Get("project_id").(string),
		TimeoutClient:        d.Get("timeout_client").(int32),
		TimeoutConnection:    d.Get("timeout_connection").(int32),
		TimeoutMember:        d.Get("timeout_member").(int32),
	}
	listener, _, err := cli.VserverClient.LoadBalancerRestControllerApi.CreateListenerUsingPOST(context.TODO(), req, projectId)
	if err != nil {
		return fmt.Errorf(utils.GetErrorMessage(err))
	}
	stateConf := &resource.StateChangeConf{
		Pending: listenerCreating,
		Target:  listenerCreated,
		Refresh: func() (interface{}, string, error) {
			resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetListenerUsingGET(context.TODO(), listener.Data.Uuid, projectId)
			if err != nil {
				return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
			}
			return resp, resp.Data.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      0,
		MinTimeout: 10 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	d.SetId(listener.Data.Uuid)
	return nil
}

func resourceListenerRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)

	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetListenerUsingGET(context.TODO(), d.Id(), projectId)
	if err != nil {
		return fmt.Errorf(utils.GetErrorMessage(err))
	}

	d.Set("project_id", resp.Data.ProjectId)
	d.Set("load_balancer_id", resp.Data.LoadBalancerId)
	d.Set("listener_name", resp.Data.Name)
	d.Set("listener_protocol", resp.Data.Protocol)
	d.Set("listener_protocol_port", resp.Data.ProtocolPort)
	d.Set("timeout_client", resp.Data.TimeoutClient)
	d.Set("timeout_connection", resp.Data.TimeoutConnection)
	d.Set("timeout_member", resp.Data.TimeoutMember)

	return nil
}

func resourceListenerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceListenerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
