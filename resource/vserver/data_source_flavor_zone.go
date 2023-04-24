package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceFlavorZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFlavorZoneRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of project",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of flavor zone.",
			},
		},
	}
}
func dataSourceFlavorZoneRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	name := d.Get("name").(string)
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.FlavorZoneRestControllerApi.ListFlavorZoneUsingGET(context.TODO(), projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	for _, zone := range resp.FlavorZones {
		if zone.Name == name {
			d.SetId(zone.Id)
			return nil
		}
	}
	return fmt.Errorf("not found flavor zone with name %s ", name)
}
