package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
)

func DataSourceVolumeTypeZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolumeTypeZoneRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceVolumeTypeZoneRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)
	cli := m.(*client.VSRClient)
	resp, _, err := cli.VserverClient.VolumeTypeZoneRestControllerApi.ListVolumeTypeZoneUsingGET(context.TODO(), projectID)
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
	for _, volumeTypeZone := range resp.VolumeTypeZones {
		if volumeTypeZone.Name == name {
			d.SetId(volumeTypeZone.Id)
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", name)
}
