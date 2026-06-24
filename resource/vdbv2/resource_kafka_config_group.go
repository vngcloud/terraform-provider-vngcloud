package vdbv2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func ResourceKafkaConfigurationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaConfigurationGroupCreate,
		Read:   resourceKafkaConfigurationGroupRead,
		Update: resourceKafkaConfigurationGroupUpdate,
		Delete: resourceKafkaConfigurationGroupDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"properties": {
				Type:        schema.TypeMap,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of parameters to be applied to the configuration group.",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceKafkaConfigurationGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Kafka configuration group create")

	cli := m.(*client.Client)

	createRequest := vdbv2.ConfigGroupCreateRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Properties:  parseKafkaConfigGroupProperties(d.Get("properties").(map[string]interface{})),
	}

	resp, httpResponse, _ := cli.Vdbv2Client.KafkaConfigurationGroupAPIApi.CreateConfigGroup(context.TODO(), createRequest)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		return fmt.Errorf("request fail with errMsg : %s", responseBody)
	}

	if resp.Id == "" {
		return fmt.Errorf("Create config group failed. Id is empty")
	}

	log.Println("[DEBUG]  Created kafka config group " + resp.Id)
	d.SetId(resp.Id)
	time.Sleep(10 * time.Second)

	return resourceKafkaConfigurationGroupRead(d, m)
}

func resourceKafkaConfigurationGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Kafka configuration group read")

	cli := m.(*client.Client)
	configID := d.Id()

	resp, httpResponse, _ := cli.Vdbv2Client.KafkaConfigurationGroupAPIApi.GetConfigGroupById(context.TODO(), configID)

	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return fmt.Errorf("config group not found: %s", configID)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		return fmt.Errorf("request fail with errMsg : %s", responseBody)
	}

	d.Set("name", resp.Name)
	d.Set("description", resp.Description)
	d.Set("status", resp.Status)
	d.Set("portal_user_id", resp.PortalUserId)
	d.Set("created_at", resp.CreatedAt.String())
	d.Set("versions", flattenKafkaConfigGroupVersions(resp.Versions))

	if latest, ok := latestKafkaConfigGroupVersion(resp.Versions); ok {
		d.Set("current_version_id", latest.Id)
		d.Set("current_version", latest.Version)
		// d.Set("properties", flattenKafkaConfigGroupProperties(latest.Properties))
	}

	return nil
}

func resourceKafkaConfigurationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Kafka configuration group update")

	if !d.HasChange("properties") {
		return resourceKafkaConfigurationGroupRead(d, m)
	}

	cli := m.(*client.Client)
	configID := d.Id()

	updateRequest := vdbv2.ConfigGroupVersionCreateRequest{
		Properties: parseKafkaConfigGroupProperties(d.Get("properties").(map[string]interface{})),
	}

	_, httpResponse, _ := cli.Vdbv2Client.KafkaConfigurationGroupAPIApi.CreateConfigGroupVersion(context.TODO(), updateRequest, configID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		return fmt.Errorf("request fail with errMsg : %s", responseBody)
	}
	time.Sleep(10 * time.Second)

	return resourceKafkaConfigurationGroupRead(d, m)
}

func resourceKafkaConfigurationGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Kafka configuration group delete")

	cli := m.(*client.Client)
	configID := d.Id()

	_, httpResponse, _ := cli.Vdbv2Client.KafkaConfigurationGroupAPIApi.DeleteConfigGroup(context.TODO(), configID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		return fmt.Errorf("request fail with errMsg : %s", responseBody)
	}
	time.Sleep(10 * time.Second)

	d.SetId("")
	return nil
}

func parseKafkaConfigGroupProperties(input map[string]interface{}) []vdbv2.ConfigGroupPropertyDto {
	properties := make([]vdbv2.ConfigGroupPropertyDto, 0, len(input))
	for k, v := range input {
		properties = append(properties, vdbv2.ConfigGroupPropertyDto{
			Key:   k,
			Value: v.(string),
		})
	}
	return properties
}

func flattenKafkaConfigGroupProperties(properties []vdbv2.ConfigGroupPropertyDto) map[string]interface{} {
	res := make(map[string]interface{}, len(properties))
	for _, p := range properties {
		res[p.Key] = p.Value
	}
	return res
}

func flattenKafkaConfigGroupVersions(versions []vdbv2.ConfigGroupVersionDto) []interface{} {
	res := make([]interface{}, len(versions))
	for i, v := range versions {
		res[i] = map[string]interface{}{
			"id":         v.Id,
			"version":    int(v.Version),
			"created_at": v.CreatedAt.String(),
		}
	}
	return res
}

func latestKafkaConfigGroupVersion(versions []vdbv2.ConfigGroupVersionDto) (vdbv2.ConfigGroupVersionDto, bool) {
	if len(versions) == 0 {
		return vdbv2.ConfigGroupVersionDto{}, false
	}
	latest := versions[0]
	for _, v := range versions[1:] {
		if v.Version > latest.Version {
			latest = v
		}
	}
	return latest, true
}
