package vdb

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"log"
	"strconv"
)

func DataSourceBackupStoragePackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBackupStoragePackageRead,
		Schema: map[string]*schema.Schema{
			"engine_group": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quota": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeString,
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

	engineGroup := d.Get("engine_group").(int)
	listPackageResp, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.GetListPackageByEngineGroupUsingGET(context.TODO(), int32(engineGroup), cli.ProjectId)
	if listPackageResp.Success == false {
		return errors.New(listPackageResp.ErrorMsg)
	}
	numPackages := len(listPackageResp.Data)
	log.Println("[DEBUG] Num backup storage packages: " + strconv.Itoa(numPackages))

	for i := 0; i < numPackages; i++ {
		if listPackageResp.Data[i].Name == d.Get("name") {
			idStr := strconv.Itoa(int(listPackageResp.Data[i].Id))
			d.SetId(idStr)
			d.Set("quota", listPackageResp.Data[i].Quota)
		}
	}
	if d.Get("quota") == 0 {
		return errors.New("no package name '" + d.Get("name").(string) + "' found")
	}
	return nil
}
