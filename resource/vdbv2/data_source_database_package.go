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

func DataSourceDatabasePackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePackageRead,
		Schema: map[string]*schema.Schema{
			"engine_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"engine_version": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func dataSourcePackageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Data source package read")

	cli := m.(*client.Client)

	type_ := d.Get("engine_type").(string)
	version := d.Get("engine_version").(string)

	var opts *vdbv2.RelationalDatabaseAPIApiGetFlavorsOpts
	if zoneID, ok := d.GetOk("zone_id"); ok {
		opts = &vdbv2.RelationalDatabaseAPIApiGetFlavorsOpts{
			ZoneId: optional.NewString(zoneID.(string)),
		}
	}
	listPackageResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetFlavors(context.TODO(), type_, version, opts)
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
			idStr := strconv.Itoa(int(listPackageResp.Data[i].Id))
			d.SetId(idStr)
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
