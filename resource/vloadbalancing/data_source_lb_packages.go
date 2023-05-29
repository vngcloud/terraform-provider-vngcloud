package vloadbalancing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
)

func DataSourceLBPackages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLBPackagesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"package_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data_transfer": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lb_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceLBPackagesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Reading lb packages")
	projectId := d.Get("project_id").(string)

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VlbClient.LoadBalancerRestControllerV2Api.GetPackagesUsingGET(ctx, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return diag.FromErr(errorResponse)
	}
	respJSON, _ = json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return diag.FromErr(err)
	}

	// create flatten packages
	var packages []interface{}
	for _, v := range resp.ListData {
		packageMap := map[string]interface{}{
			"uuid":              v.Uuid,
			"name":              v.Name,
			"package_type":      v.Type_,
			"connection_number": v.ConnectionNumber,
			"data_transfer":     v.DataTransfer,
			"mode":              v.Mode,
			"lb_type":           v.LbType,
		}
		packages = append(packages, packageMap)
	}

	if err := d.Set("packages", packages); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	log.Printf("Read lb packages successfully")
	return nil
}
