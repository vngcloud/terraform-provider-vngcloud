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

func DataSourceBackupStoragePackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBackupStoragePackageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quota": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBackupStoragePackageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Data source backup storage package read")

	cli := m.(*client.Client)

	listPackageResp, httpResponse, err := cli.Vdbv2Client.RelationalBackupStorageAPIApi.GetListQuotaPackage1(context.TODO())
	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if listPackageResp.Data == nil || len(listPackageResp.Data) == 0 {
		d.SetId("")
		return nil
	}

	numPackages := len(listPackageResp.Data[0].Packages)
	log.Println("[DEBUG] Num backup storage packages: " + strconv.Itoa(numPackages))

	for i := 0; i < numPackages; i++ {
		if listPackageResp.Data[0].Packages[i].PackageName == d.Get("name") {
			idStr := strconv.Itoa(int(listPackageResp.Data[0].Packages[i].PackageId))
			d.SetId(idStr)
			quota := 0
			if quota, err = strconv.Atoi(listPackageResp.Data[0].Packages[i].PackageQuota); err != nil {
				log.Println("[DEBUG] Invalid quota input of " + idStr)
				return err
			}
			d.Set("quota", quota)
			d.Set("description", listPackageResp.Data[0].Packages[i].Description)
		}
	}
	if d.Get("quota") == 0 {
		return errors.New("no package name '" + d.Get("name").(string) + "' found")
	}
	return nil
}
