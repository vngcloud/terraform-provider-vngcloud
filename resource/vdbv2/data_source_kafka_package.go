package vdbv2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceKafkaPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKafkaPackageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceKafkaPackageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Data source kafka package read")

	cli := m.(*client.Client)

	listFlavorResp, httpResponse, _ := cli.Vdbv2Client.KafkaMiscAPIApi.ListFlavor(context.TODO(), nil)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if len(listFlavorResp.Data) == 0 {
		d.SetId("")
		return nil
	}

	numPackages := len(listFlavorResp.Data)
	log.Println("[DEBUG] Num kafka packages: " + strconv.Itoa(numPackages))

	for i := 0; i < numPackages; i++ {
		if listFlavorResp.Data[i].Name == d.Get("name") {
			d.SetId(listFlavorResp.Data[i].FlavorId)
			d.Set("cpu", listFlavorResp.Data[i].Vcpus)
			d.Set("ram", listFlavorResp.Data[i].Ram)
		}
	}
	if d.Id() == "" {
		return errors.New("no kafka package name '" + d.Get("name").(string) + "' found")
	}
	return nil
}
