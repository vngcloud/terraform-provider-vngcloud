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
				Description: "id of volume acttach",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "id of server acttach",
			},
		},
	}
}
func resourceVolumeAttach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	attachVolume := vserver.AttachVolumeRequest{
		VolumeId:   d.Get("volume_id").(string),
		InstanceId: d.Get("instance_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.VolumeRestControllerApi.AttachVolumeUsingPUT(context.TODO(), attachVolume, projectID)
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
		Pending:    volumeAttaching,
		Target:     volumeAttached,
		Refresh:    resourceVolumeStateRefreshFunc(cli, attachVolume.VolumeId, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", attachVolume.VolumeId, err)
	}
	d.SetId(attachVolume.VolumeId)
	return nil
}

func resourceVolumeDetach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	detachVolume := vserver.DetachVolumeRequest{
		VolumeId:   d.Get("volume_id").(string),
		InstanceId: d.Get("instance_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.VolumeRestControllerApi.DetachVolumeUsingPUT(context.TODO(), detachVolume, projectID)
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
		Pending:    volumeDetaching,
		Target:     volumeDetached,
		Refresh:    resourceVolumeStateRefreshFunc(cli, detachVolume.VolumeId, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", detachVolume.InstanceId, err)
	}
	d.SetId("")
	return nil
}
