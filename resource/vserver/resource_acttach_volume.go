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
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceAttachVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeAttach,
		Delete: resourceVolumeDetach,
		Read:   resourceVolumeRead,
		Importer: &schema.ResourceImporter{
			StateContext: resourceAttachVolumeImport,
		},
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

func resourceAttachVolumeImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), ":")
	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:VolumeID:ServerID", d.Id())
	}
	projectID := idParts[0]
	volumeID := idParts[1]
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET2(context.TODO(), projectID, volumeID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("Get Volume fail with errMsg : %s", responseBody)
		return nil, err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	volume := resp.Data
	if volume.Status != "IN-USE" {
		err := fmt.Errorf("Volume status %s unexpected (IN-USE)", volume.Status)
		return nil, err
	}
	serverID := idParts[2]
	_, httpResponse, _ = cli.VserverClient.ServerRestControllerApi.GetServerUsingGET1(context.TODO(), projectID, serverID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("Get Server fail with errMsg : %s", responseBody)
		return nil, err
	}
	if volume.MultiAttach {
		if !contains(volume.ServerIdList, serverID) {
			err := fmt.Errorf("Server is not attached in this volume")
			return nil, err
		}
	} else if volume.ServerId != serverID {
		err := fmt.Errorf("Server is not attached in this volume")
		return nil, err
	}
	d.SetId(volumeID)
	d.Set("project_id", projectID)
	d.Set("server_id", serverID)
	d.Set("volume_id", volumeID)
	return []*schema.ResourceData{d}, nil
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
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
