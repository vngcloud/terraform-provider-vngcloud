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
	listPackageResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetFlavors1(context.TODO(), type_, version, 0)
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
		}
	}
	if d.Id() == "" {
		return errors.New("no package name '" + d.Get("name").(string) + "' found")
	}
	return nil
}
