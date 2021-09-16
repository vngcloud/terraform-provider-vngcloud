package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceImageRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"flavor_zone_id": {
				Required: true,
				Type:     schema.TypeString,
			},
			"licence": {
				Computed: true,
				Type:     schema.TypeBool,
			},
		},
	}
}
func dataSourceImageRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)
	flavorID := d.Get("flavor_zone_id")
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ImageRestControllerApi.ListOSImageUsingGET(context.TODO(), projectID)
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
	for _, image := range resp.Images {
		if image.ImageVersion == name {
			for _, FlavorID := range image.FlavorZoneIds {
				if FlavorID == flavorID {
					d.SetId(image.Id)
					d.Set("licence", image.Licence)
					return nil
				}
			}

		}
	}
	return fmt.Errorf("not found image with config %s ", name)
}
