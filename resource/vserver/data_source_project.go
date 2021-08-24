package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceProjectRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("project_id").(string)
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ProjectRestControllerApi.GetProjectUsingGET(context.TODO(), id)
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
	d.SetId(id)
	return nil
}
