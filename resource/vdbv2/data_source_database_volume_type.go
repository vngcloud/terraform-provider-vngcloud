package vdbv2

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceDatabaseVolumeType() *schema.Resource {
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
			"max_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceVolumeTypeRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source volume type")

	cli := m.(*client.Client)

	listVolumeTypeResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetVolumeTypes2(context.TODO(), 0)
	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if listVolumeTypeResp.Data == nil || len(listVolumeTypeResp.Data.Data) == 0 {
		d.SetId("")
		return nil
	}

	numVolumeTypes := len(listVolumeTypeResp.Data.Data)
	log.Println("[DEBUG] Num volume types: " + strconv.Itoa(numVolumeTypes))

	for i := 0; i < numVolumeTypes; i++ {
		if listVolumeTypeResp.Data.Data[i].Type_ == d.Get("type") {
			d.Set("volume_type_zone_id", listVolumeTypeResp.Data.Data[i].VolumeTypeZoneId)
			d.Set("max_size", listVolumeTypeResp.Data.Data[i].MaxVolumeSize)
			d.Set("min_size", listVolumeTypeResp.Data.Data[i].MinVolumeSize)
			d.SetId(d.Get("type").(string))
		}
	}

	return nil
}
