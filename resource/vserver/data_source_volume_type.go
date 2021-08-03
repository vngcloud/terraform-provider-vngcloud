package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
)

func DataSourceVolumeType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolumeTypeRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_type_zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iops": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"min_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"through_put": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
func dataSourceVolumeTypeRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	zoneID := d.Get("volume_type_zone_id").(string)
	name := d.Get("name").(string)
	cli := m.(*client.VSRClient)
	resp, _, err := cli.VserverClient.VolumeTypeRestControllerApi.ListVolumeTypeUsingGET(context.TODO(), projectID, zoneID)
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
	for _, volumeType := range resp.VolumeTypes {
		if volumeType.Name == name {
			d.Set("iops", volumeType.Iops)
			d.Set("max_size", volumeType.MaxSize)
			d.Set("min_size", volumeType.MinSize)
			d.Set("through_put", volumeType.ThroughPut)
			d.SetId(volumeType.Id)
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", name)
}
