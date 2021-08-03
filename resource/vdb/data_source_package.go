package vdb

import (
	"context"
	"encoding/json"
	"fmt"
	"git.vngcloud.tech/vdb/vdb-terraform/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func DataSourcePackage() *schema.Resource {
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
			"package_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"price_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourcePackageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source package")

	cli := m.(*client.Client)

	listPackageResp, _, err := cli.VdbClient.VdbPackageEndpointApi.GetListPackageByEngineTypeAndVersionUsingGET(
		context.TODO(), d.Get("engine_type").(string), d.Get("engine_version").(string), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when getting package info: %s", err)
	}
	if !listPackageResp.Success {
		return fmt.Errorf("Error when getting package info: %s", listPackageResp.ErrorMsg)
	}

	numPackages := len(listPackageResp.Data)
	log.Println("[DEBUG] [trihm4] Num packages: " + strconv.Itoa(numPackages))

	for i := 0; i < numPackages; i++ {
		if listPackageResp.Data[i].Name == d.Get("name") {
			d.Set("package_id", strconv.Itoa(int(listPackageResp.Data[i].Id)))
			d.Set("price_key", listPackageResp.Data[i].PriceKey)
			d.Set("zone_id", listPackageResp.Data[i].ZoneId)

			var pkgConfig packageConfig

			log.Println("[DEBUG] [trihm4] Config: " + listPackageResp.Data[i].Config)
			err = json.Unmarshal([]byte(listPackageResp.Data[i].Config), &pkgConfig)
			if err != nil {
				return fmt.Errorf("Error when reading config of package listPackageResp.Data[i].Id: %s", err)
			}

			d.Set("flavor_id", pkgConfig.FlavorId)
			d.Set("vcpus", pkgConfig.Vcpus)
			d.Set("ram", pkgConfig.Ram)

			d.SetId(d.Get("name").(string))
		}
	}

	return nil
}

type packageConfig struct {
	FlavorId string
	Vcpus    int
	Ram      int
}
