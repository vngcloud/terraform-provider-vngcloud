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

func ResourceAttachExternalInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalInterfaceAttach,
		Delete: resourceExternalInterfaceDetach,
		Read:   resourceExternalInterfaceRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_interface_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "id of network interface",
			},
			"server_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "id of server",
			},
		},
	}
}
func resourceExternalInterfaceAttach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceId := d.Get("network_interface_id").(string)
	serverID := d.Get("server_id").(string)
	attachExternalNetworkInterface := vserver.AttachExternalNetworkInterfaceRequest{}
	attachExternalNetworkInterface.ExternalNetworkInterfaceId = networkInterfaceId
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.AttachExternalNetworkInterfaceUsingPOST(context.TODO(), attachExternalNetworkInterface, projectID, serverID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if err != nil {
		return fmt.Errorf("error attach External Network Interface (%s) %s", networkInterfaceId, err)
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	d.SetId(resp.Data.Uuid)
	return nil
}

func resourceExternalInterfaceRead(d *schema.ResourceData, m interface{}) error {

}

func resourceExternalInterfaceDetach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceId := d.Get("network_interface_id").(string)
	serverID := d.Get("server_id").(string)
	detachExternalNetworkInterface := vserver.DetachExternalNetworkInterfaceRequest{}
	detachExternalNetworkInterface.NetworkInterfaceId = networkInterfaceId
	cli := m.(*client.Client)

	httpResponse, err := cli.VserverClient.ServerRestControllerApi.DetachExternalNetworkInterfaceUsingDELETE(context.TODO(), detachExternalNetworkInterface, projectID, serverID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if err != nil {
		return fmt.Errorf("error for detach External Network Interface (%s) %s", networkInterfaceId, err)
	}
	d.SetId("")
	return nil
}
