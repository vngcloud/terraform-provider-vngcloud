package vdb

import (
	"context"
	"git.vngcloud.tech/vdb/vdb-terraform/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func DataSourceDb() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDbRead,
		Schema: map[string]*schema.Schema{
			"dbs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDbRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	listDbResp, _, _ := cli.VdbClient.VdbInstanceEndPointApi.
		GetListDbInstanceByEngineGroupUsingGET(context.TODO(), 1, cli.ProjectId)

	numDbs := len(listDbResp.Data)

	dbs := make([]interface{}, numDbs, numDbs)

	for i := 0; i < numDbs; i++ {
		db := make(map[string]interface{})

		db["id"] = listDbResp.Data[i].Id
		db["name"] = listDbResp.Data[i].Name
		db["status"] = listDbResp.Data[i].Status

		dbs[i] = db
	}

	d.Set("dbs", dbs)

	curTime := time.Now().Unix()
	timeStr := strconv.Itoa(int(curTime))

	d.SetId(timeStr)

	return nil
}
