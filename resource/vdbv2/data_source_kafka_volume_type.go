package vdbv2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceKafkaVolumeType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKafkaVolumeTypeRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceKafkaVolumeTypeRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Data source kafka volume type read")

	cli := m.(*client.Client)

	listVolumeTypeResp, httpResponse, _ := cli.Vdbv2Client.KafkaMiscAPIApi.GetVolumeTypes2(context.TODO())

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
	log.Println("[DEBUG] Num kafka volume types: " + strconv.Itoa(numVolumeTypes))

	for i := 0; i < numVolumeTypes; i++ {
		if strings.TrimPrefix(listVolumeTypeResp.Data.Data[i].Type_, "kafka.") == d.Get("type") {
			d.SetId(listVolumeTypeResp.Data.Data[i].KafkaUuid)
			d.Set("max_size", listVolumeTypeResp.Data.Data[i].MaxVolumeSize)
			d.Set("min_size", listVolumeTypeResp.Data.Data[i].MinVolumeSize)
		}
	}
	if d.Id() == "" {
		return errors.New("no kafka volume type '" + d.Get("type").(string) + "' found")
	}
	return nil
}
