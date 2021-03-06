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
		},
	}
}
func resourceVolumeStateRefreshFunc(cli *client.Client, volumeID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET(context.TODO(), projectID, volumeID)
		if err != nil {
			return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			return nil, "", fmt.Errorf("Error describing instance: %s", resp.ErrorMsg)
		}
		volume := resp.Volumes[0]
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
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.VolumeRestControllerApi.CreateVolumeUsingPOST(context.TODO(), a, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:    volumeCreating,
		Target:     volumeCreated,
		Refresh:    resourceVolumeStateRefreshFunc(cli, resp.Volumes[0].Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Volumes[0].Uuid, err)
	}
	d.SetId(resp.Volumes[0].Uuid)
	return resourceVolumeRead(d, m)
}

func resourceVolumeRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	volumeID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET(context.TODO(), projectID, volumeID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if len(resp.Volumes) == 0 {
		d.SetId("")
		return nil
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	volume := resp.Volumes[0]
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
			VolumeId:        d.Id(),
			NewVolumeTypeId: d.Get("volume_type_id").(string),
		}
		cli := m.(*client.Client)
		resp, _, err := cli.VserverClient.VolumeRestControllerApi.ResizeVolumeUsingPUT(context.TODO(), projectID, resizeVolume)
		if err != nil {
			return err
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			err := fmt.Errorf(resp.ErrorMsg)
			return err
		}
		return nil
	}
	return resourceVolumeRead(d, m)

}

func resourceVolumeDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteVolume := vserver.DeleteVolumeRequest{
		VolumeId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.VolumeRestControllerApi.DeleteVolumeUsingDELETE(context.TODO(), deleteVolume, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	return resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.VolumeRestControllerApi.GetVolumeUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Volumes) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Volumes[0].Status))
	})
}
