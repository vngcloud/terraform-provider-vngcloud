package vserver

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"golang.org/x/net/context"
	"log"
)

func DataSourceK8sVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceK8sVersionRead,
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

func dataSourceK8sVersionRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	k8sVersionName := d.Get("name").(string)

	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.ListK8sVersionUsingGET(context.TODO(), projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	for _, k8sVersion := range resp {
		if k8sVersion.Name == k8sVersionName {
			d.SetId(k8sVersion.Uuid)
			d.Set("id", string(k8sVersion.Uuid))
			return nil
		}
	}
	return fmt.Errorf("not found resource with name %s ", k8sVersionName)
}
