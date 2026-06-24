package vdbv2

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func DataSourcePostgreSQLClusterVolumeType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePostgreVolumeTypeRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func dataSourcePostgreVolumeTypeRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source volume type")

	cli := m.(*client.Client)

	var opts *vdbv2.PostgreSQLClusterAPIApiGetVolumeTypesPostgreClusterOpts
	if zoneID, ok := d.GetOk("zone_id"); ok {
		opts = &vdbv2.PostgreSQLClusterAPIApiGetVolumeTypesPostgreClusterOpts{
			ZoneId: optional.NewString(zoneID.(string)),
		}
	}
	listVolumeTypeResp, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.GetVolumeTypesPostgreCluster(context.TODO(), opts)
	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if listVolumeTypeResp.Data == nil || len(listVolumeTypeResp.Data) == 0 {
		d.SetId("")
		return nil
	}

	numVolumeTypes := len(listVolumeTypeResp.Data)
	log.Println("[DEBUG] Num volume types: " + strconv.Itoa(numVolumeTypes))

	for i := 0; i < numVolumeTypes; i++ {
		if listVolumeTypeResp.Data[i].Type_ == d.Get("type") {
			d.Set("volume_type_zone_id", listVolumeTypeResp.Data[i].VolumeTypeZoneId)
			d.Set("max_size", listVolumeTypeResp.Data[i].MaxVolumeSize)
			d.Set("min_size", listVolumeTypeResp.Data[i].MinVolumeSize)
			d.SetId(listVolumeTypeResp.Data[i].Id)
			d.Set("zone_id", listVolumeTypeResp.Data[i].ZoneId)
		}
	}

	return nil
}
