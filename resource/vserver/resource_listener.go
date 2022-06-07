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

func ResourceListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceListenerCreate,
		Read:   resourceListenerRead,
		Update: resourceListenerUpdate,
		Delete: resourceListenerDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:ListenerID", d.Id())
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
			"allowed_cidrs": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0/0",
			},
			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
				ForceNew: true,
			},
			"protocol_port": {
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
			"default_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_authorities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},
			"default_certificate_authority": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceListenerCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create listener")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	certificateAuthoritiesInterface := d.Get("certificate_authorities").([]interface{})
	certificateAuthorities := make([]string, 0)
	for _, cert := range certificateAuthoritiesInterface {
		certificateAuthorities = append(certificateAuthorities, cert.(string))
	}
	req := vserver.CreateListenerRequest{
		AllowedCidrs:                d.Get("allowed_cidrs").(string),
		ListenerName:                d.Get("name").(string),
		ListenerProtocol:            d.Get("protocol").(string),
		ListenerProtocolPort:        int32(d.Get("protocol_port").(int)),
		LoadBalancerId:              d.Get("load_balancer_id").(string),
		ProjectId:                   d.Get("project_id").(string),
		TimeoutClient:               int32(d.Get("timeout_client").(int)),
		TimeoutConnection:           int32(d.Get("timeout_connection").(int)),
		TimeoutMember:               int32(d.Get("timeout_member").(int)),
		DefaultPoolId:               d.Get("default_pool_id").(string),
		CertificateAuthorities:      certificateAuthorities,
		DefaultCertificateAuthority: d.Get("default_certificate_authority").(string),
	}
	listener, _, err := cli.VserverClient.LoadBalancerRestControllerApi.CreateListenerUsingPOST(context.TODO(), req, projectId)

	respJSON, _ := json.Marshal(listener)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	stateConf := &resource.StateChangeConf{
		Pending: listenerCreating,
		Target:  listenerCreated,
		Refresh: func() (interface{}, string, error) {
			resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetListenerUsingGET(context.TODO(), listener.Data.Uuid, projectId)
			respJSON, _ := json.Marshal(resp)
			log.Printf("-------------------------------------\n")
			log.Printf("%s\n", string(respJSON))
			log.Printf("-------------------------------------\n")
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

	log.Printf("Create listener successfully")

	return nil
}

func resourceListenerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("Read listener")
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)

	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetListenerUsingGET(context.TODO(), d.Id(), projectId)
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	d.Set("allowed_cidrs", resp.Data.AllowedCidrs)
	d.Set("project_id", resp.Data.ProjectId)
	d.Set("load_balancer_id", resp.Data.LoadBalancerId)
	d.Set("name", resp.Data.Name)
	d.Set("protocol", resp.Data.Protocol)
	d.Set("protocol_port", resp.Data.ProtocolPort)
	d.Set("timeout_client", resp.Data.TimeoutClient)
	d.Set("timeout_connection", resp.Data.TimeoutConnection)
	d.Set("timeout_member", resp.Data.TimeoutMember)
	d.Set("headers", resp.Data.Headers)
	d.Set("default_pool_id", resp.Data.DefaultPoolId)
	// d.Set("certificate_authorities", resp.)

	log.Printf("Read listener successfully")

	return nil
}

func resourceListenerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Update listener")
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)
	req := vserver.UpdateListenerRequest{
		AllowedCidrs:                d.Get("allowed_cidrs").(string),
		DefaultPoolId:               d.Get("default_pool_id").(string),
		ListenerId:                  d.Id(),
		TimeoutClient:               int32(d.Get("timeout_client").(int)),
		TimeoutConnection:           int32(d.Get("timeout_connection").(int)),
		TimeoutMember:               int32(d.Get("timeout_member").(int)),
		DefaultCertificateAuthority: d.Get("default_certificate_authority").(string),
	}
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.UpdateListenerUsingPUT(context.TODO(), projectId, req)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}
	stateConf := &resource.StateChangeConf{
		Pending: listenerUpdating,
		Target:  listenerCreated,
		Refresh: func() (interface{}, string, error) {
			resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.GetListenerUsingGET(context.TODO(), d.Id(), projectId)

			respJSON, _ := json.Marshal(resp)
			log.Printf("-------------------------------------\n")
			log.Printf("%s\n", string(respJSON))
			log.Printf("-------------------------------------\n")
			if err != nil {
				return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
			}
			return resp, resp.Data.Status, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      0,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	log.Printf("Update listener successfully")

	return resourceListenerRead(d, m)
}

func resourceListenerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("Delete listener")
	cli := m.(*client.Client)
	projectId := d.Get("project_id").(string)

	req := vserver.DeleteListenerRequest{
		ListenerId: d.Id(),
	}
	resp, _, err := cli.VserverClient.LoadBalancerRestControllerApi.DeleteListenerUsingDELETE(context.TODO(), req, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return utils.GetErrorMessage(err)
	}

	log.Printf("Delete listener successfully")

	return nil
}
