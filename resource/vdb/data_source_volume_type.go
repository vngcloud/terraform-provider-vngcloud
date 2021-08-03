package vdb

import (
	"context"
	"fmt"
	"git.vngcloud.tech/vdb/vdb-terraform/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func DataSourceVolumeType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolumeTypeRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_type_zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVolumeTypeRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source volume type")

	cli := m.(*client.Client)

	listVolumeTypeResp, _, err := cli.VdbClient.VdbVolumeTypeEndPointApi.GetListActiveVolumeTypeUsingGET(context.TODO(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when getting volume type info: %s", err)
	}
	if !listVolumeTypeResp.Success {
		return fmt.Errorf("Error when getting volume type info: %s", listVolumeTypeResp.ErrorMsg)
	}

	numVolumeTypes := len(listVolumeTypeResp.Data)
	log.Println("[DEBUG] [trihm4] Num volume types: " + strconv.Itoa(numVolumeTypes))

	for i := 0; i < numVolumeTypes; i++ {
		if listVolumeTypeResp.Data[i].Type_ == d.Get("type") {
			d.Set("volume_type_zone_id", listVolumeTypeResp.Data[i].VolumeTypeZoneId)

			d.SetId(d.Get("type").(string))
		}
	}

	return nil
}
