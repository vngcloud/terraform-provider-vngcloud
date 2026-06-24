package vdbv2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func DataSourcePostgreSQLClusterPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePostgrePackageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"backup_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourcePostgrePackageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Data source package read")

	cli := m.(*client.Client)

	var opts *vdbv2.PostgreSQLClusterAPIApiListFlavorPostgreClusterOpts
	if zoneID, ok := d.GetOk("zone_id"); ok {
		opts = &vdbv2.PostgreSQLClusterAPIApiListFlavorPostgreClusterOpts{
			ZoneId: optional.NewString(zoneID.(string)),
		}
	}
	listPackageResp, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.ListFlavorPostgreCluster(context.TODO(), opts)
	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if listPackageResp.Data == nil || len(listPackageResp.Data) == 0 {
		d.SetId("")
		return nil
	}

	numPackages := len(listPackageResp.Data)
	log.Println("[DEBUG] Num packages: " + strconv.Itoa(numPackages))

	for i := 0; i < numPackages; i++ {
		if listPackageResp.Data[i].Name == d.Get("name") {
			d.SetId(listPackageResp.Data[i].Id)
			d.Set("cpu", listPackageResp.Data[i].Vcpus)
			d.Set("ram", listPackageResp.Data[i].Ram)
			d.Set("backup_size", listPackageResp.Data[i].BackupSize)
			d.Set("zone_id", listPackageResp.Data[i].LocateZoneId)
		}
	}
	if d.Id() == "" {
		return errors.New("no package name '" + d.Get("name").(string) + "' found")
	}
	return nil
}
