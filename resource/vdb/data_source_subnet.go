package vdb

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"log"
	"strconv"
)

func DataSourceSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubnetRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Data source subnet")

	cli := m.(*client.Client)

	listSubnetResp, _, err := cli.VdbClient.VdbNetworkEndPointApi.GetListSubnetUsingGET(
		context.TODO(), d.Get("network_id").(string), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when getting subnet info: %s", err)
	}
	if !listSubnetResp.Success {
		return fmt.Errorf("Error when getting subnet info: %s", listSubnetResp.ErrorMsg)
	}

	numSubnets := len(listSubnetResp.InterfaceSubnetList)
	log.Println("[DEBUG] [trihm4] Num subnets: " + strconv.Itoa(numSubnets))

	for i := 0; i < numSubnets; i++ {
		if listSubnetResp.InterfaceSubnetList[i].Name == d.Get("name") {
			d.SetId(listSubnetResp.InterfaceSubnetList[i].Uuid)
		}
	}

	return nil
}
