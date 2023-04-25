package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
)

func ResourceAttachInternalInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceInternalInterfaceAttach,
		Delete: resourceInternalInterfaceDetach,
		Read:   resourceInternalInterfaceRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"server_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of server",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of Subnet",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Fixed Ip",
			},
			"floating_ip_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of Floating IP",
			},
			"display_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP display",
			},
		},
	}
}

func resourceInternalInterfaceAttach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	subnetID := d.Get("subnet_id").(string)
	serverID := d.Get("server_id").(string)
	floatingIpID := d.Get("floating_ip_id").(string)
	attachInternalNetworkInterface := vserver.AttachNetworkInterfaceWithWanIpRequest{}
	attachInternalNetworkInterface.Ip = d.Get("ip").(string)
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.AttachNetworkInterfaceWithWanIpUsingPOST(context.TODO(), attachInternalNetworkInterface, floatingIpID, projectID, serverID, subnetID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if err != nil {
		return fmt.Errorf("error attach Internal Network Interface %s", err)
	}

	respJSON, _ := json.Marshal(resp.Data)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	d.SetId(resp.Data.Uuid)
	return nil
}
func resourceInternalInterfaceDetach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	interfaceNetworkInterfaceId := d.Id()
	serverID := d.Get("server_id").(string)
	detachInternalNetworkInterface := vserver.DetachNetworkInterfaceWithWanIpRequest{}
	detachInternalNetworkInterface.NetworkInterfaceId = interfaceNetworkInterfaceId
	cli := m.(*client.Client)

	httpResponse, err := cli.VserverClient.ServerRestControllerApi.DetachNetworkInterfaceWithWanIpUsingDELETE(context.TODO(), detachInternalNetworkInterface, projectID, serverID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if err != nil {
		return fmt.Errorf("error for detach External Network Interface (%s) %s", interfaceNetworkInterfaceId, err)
	}
	d.SetId("")
	return nil
}
func resourceInternalInterfaceRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	interfaceNetworkInterfaceId := d.Id()
	cli := m.(*client.Client)

	resp, httpResponse, _ := cli.VserverClient.ServerRestControllerApi.GetExternalNetworkInterfaceUsingGET(context.TODO(), interfaceNetworkInterfaceId, projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	networkInterface := resp.Data
	d.Set("display_ip", networkInterface.FixedIp)
	return nil
}
