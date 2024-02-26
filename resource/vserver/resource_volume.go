package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeCreate,
		Read:   resourceVolumeRead,
		Update: resourceVolumeUpdate,
		Delete: resourceVolumeDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:VolumeID", d.Id())
				}
				projectID := idParts[0]
				volumeID := idParts[1]
				d.SetId(volumeID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bootable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multi_attach": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceVolumeStateRefreshFunc(cli *client.Client, volumeID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET2(context.TODO(), projectID, volumeID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		volume := resp.Data
		return volume, volume.Status, nil
	}
}
func resourceVolumeCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	a := vserver.CreateVolumeRequest{
		EncryptionType: d.Get("encryption_type").(string),
		Name:           d.Get("name").(string),
		Size:           int32(d.Get("size").(int)),
		VolumeTypeId:   d.Get("volume_type_id").(string),
		MultiAttach:    d.Get("multi_attach").(bool),
		IsPoc:          d.Get("is_poc").(bool),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.CreateVolumeUsingPOST1(context.TODO(), a, projectID)
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
		Pending:    volumeCreating,
		Target:     volumeCreated,
		Refresh:    resourceVolumeStateRefreshFunc(cli, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceVolumeRead(d, m)
}

func resourceVolumeRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	volumeID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET2(context.TODO(), projectID, volumeID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	volume := resp.Data
	d.Set("encryption_type", volume.EncryptionType)
	d.Set("name", volume.Name)
	d.Set("size", int(volume.Size))
	d.Set("volume_type_id", volume.VolumeTypeId)
	d.Set("bootable", volume.Bootable)
	return nil
}

func resourceVolumeUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("size") || d.HasChange("volume_type_id") {
		projectID := d.Get("project_id").(string)
		resizeVolume := vserver.ResizeVolumeRequest{
			NewSize:         int32(d.Get("size").(int)),
			NewVolumeTypeId: d.Get("volume_type_id").(string),
		}
		cli := m.(*client.Client)
		resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.ResizeVolumeUsingPUT1(context.TODO(), projectID, resizeVolume, d.Id())
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			oldSize, _ := d.GetChange("size")
			oldType, _ := d.GetChange("volume_type_id")
			d.Set("size", oldSize)
			d.Set("volume_type_id", oldType)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		stateConf := &resource.StateChangeConf{
			Pending:    volumeResizing,
			Target:     volumeResized,
			Refresh:    resourceVolumeStateRefreshFunc(cli, resp.Data.Uuid, projectID),
			Timeout:    d.Timeout(schema.TimeoutCreate),
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		_, err = stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Data.Uuid, err)
		}
		return nil
	}
	return resourceVolumeRead(d, m)

}

func resourceVolumeDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.DeleteVolumeUsingDELETE1(context.TODO(), projectID, d.Id())
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
		Pending:    volumeDeleting,
		Target:     volumeDeleted,
		Refresh:    resourceVolumeDeleteStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func resourceVolumeDeleteStateRefreshFunc(cli *client.Client, volumeId string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET2(context.TODO(), projectID, volumeId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.Volume{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		volume := resp.Data
		return volume, volume.Status, nil
	}
}
