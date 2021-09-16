package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceServerGroupPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerGroupPolicyRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
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
func dataSourceServerGroupPolicyRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("project_id").(string)
	name := d.Get("name").(string)

	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerGroupRestControllerApi.ListServerGroupPolicyUsingGET(context.TODO(), id)
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
	for _, policy := range resp.ServerGroupPolicies {
		if policy.Name == name {
			d.SetId(policy.Uuid)
			d.Set("status", policy.Status)
			d.Set("description", policy.Description)
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", name)
}
