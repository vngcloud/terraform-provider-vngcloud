package vserver

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"golang.org/x/net/context"
	"log"
)

func DataSourceK8sNetworkType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceK8sNetworkTypeRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceK8sNetworkTypeRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	k8sNetworkTypeName := d.Get("name").(string)
	cli := m.(*client.Client)

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.ListK8sNetworkTypeUsingGET(context.TODO(), projectID)

	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	for _, k8sNetworkType := range resp {
		if k8sNetworkType.Name == k8sNetworkTypeName {
			d.SetId(k8sNetworkType.Uuid)
			d.Set("id", string(k8sNetworkType.Uuid))
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", k8sNetworkTypeName)
}
