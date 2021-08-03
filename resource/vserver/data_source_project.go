package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
)

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceProjectRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ProjectRestControllerApi.ListProjectUsingGET(context.TODO())
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
	for _, project := range resp.Projects {
		if project.Name == name {
			d.SetId(project.Id)
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", name)
}
