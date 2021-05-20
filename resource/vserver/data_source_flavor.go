package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceFlavor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFlavorRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"flavor_zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bandwidth_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gpu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
func dataSourceFlavorRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)
	flavorZoneId := d.Get("flavor_zone_id").(string)
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.FlavorRestControllerApi.ListFlavorUsingGET(context.TODO(), flavorZoneId, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	for _, flavor := range resp.Flavors {
		if flavor.Name == name {
			d.SetId(flavor.FlavorId)
			d.Set("bandwidth", int(flavor.Bandwidth))
			d.Set("bandwidth_unit", flavor.BandwidthUnit)
			d.Set("cpu", int(flavor.Cpu))
			d.Set("gpu", int(flavor.Gpu))
			d.Set("memory", int(flavor.Memory))
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", name)
}
