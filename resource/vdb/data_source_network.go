package vdb

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"log"
	"strconv"
)

func DataSourceNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source network")

	cli := m.(*client.Client)

	listNetworkResp, _, err := cli.VdbClient.VdbNetworkEndPointApi.GetListNetworkUsingGET(context.TODO(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when getting network info: %s", err)
	}
	if !listNetworkResp.Success {
		return fmt.Errorf("Error when getting network info: %s", listNetworkResp.ErrorMsg)
	}

	numNetworks := len(listNetworkResp.ListNetwork)
	log.Println("[DEBUG] [trihm4] Num networks: " + strconv.Itoa(numNetworks))

	for i := 0; i < numNetworks; i++ {
		if listNetworkResp.ListNetwork[i].Name == d.Get("name") {
			d.SetId(listNetworkResp.ListNetwork[i].Id)
		}
	}

	return nil
}
