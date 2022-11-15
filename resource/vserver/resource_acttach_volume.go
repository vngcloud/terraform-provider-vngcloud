package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceAttachVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeAttach,
		Delete: resourceVolumeDetach,
		Read:   resourceVolumeRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "id of volume",
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
func resourceVolumeAttach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	volumeID := d.Get("volume_id").(string)
	serverID := d.Get("server_id").(string)
	attachVolume := vserver.AttachVolumeRequest{}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.AttachVolumeUsingPUT1(context.TODO(), attachVolume, projectID, serverID, volumeID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    volumeAttaching,
		Target:     volumeAttached,
		Refresh:    resourceVolumeStateRefreshFunc(cli, volumeID, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for attach volume (%s) %s", volumeID, err)
	}
	d.SetId(volumeID)
	return nil
}

func resourceVolumeDetach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	volumeID := d.Get("volume_id").(string)
	serverID := d.Get("server_id").(string)
	detachVolume := vserver.DetachVolumeRequest{}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.DetachVolumeUsingPUT1(context.TODO(), detachVolume, projectID, serverID, volumeID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    volumeDetaching,
		Target:     volumeDetached,
		Refresh:    resourceVolumeStateRefreshFunc(cli, volumeID, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for detach volume (%s) %s", volumeID, err)
	}
	d.SetId("")
	return nil
}
