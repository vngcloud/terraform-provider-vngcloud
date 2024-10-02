package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
)

func ResourceRelationalConfigurationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationGroupCreate,
		Read:   resourceConfigurationGroupRead,
		Delete: resourceConfigurationGroupDelete,

		Schema: map[string]*schema.Schema{
			"datastore_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"datastore_version": {
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"values": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `List of parameters to be applied to the configuration group.`,
			},
		},
	}
}

func resourceConfigurationGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group create")

	cli := m.(*client.Client)
	createRequest := vdb.ConfigurationRequest{
		DatastoreType:    d.Get("datastore_type").(string),
		DatastoreVersion: d.Get("datastore_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		EngineGroup:      int32(1),
	}

	body, _ := json.Marshal(createRequest)

	resp, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.CreateConfig1(context.TODO(), string(body))

	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	log.Println("[DEBUG]  Created " + resp.Data.Id)

	d.SetId(resp.Data.Id)

	time.Sleep(1 * time.Second)
	_, ok := d.Get("values").([]interface{})
	if ok {
		err := resourceConfigurationGroupUpdate(d, m)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return resourceConfigurationGroupRead(d, m)
}

func resourceConfigurationGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group read")

	cli := m.(*client.Client)
	configID := d.Id()

	resp, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.GetConfigsById1(context.TODO(), configID)
	if err != nil {
		return err
	}

	if resp.Data == nil {
		d.SetId("")
		return nil
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	d.Set("datastore_type", resp.Data.DatastoreName)
	d.Set("datastore_version", resp.Data.DatastoreVersionName)
	d.Set("description", resp.Data.Description)
	d.Set("name", resp.Data.Name)
	d.Set("created", resp.Data.Created)
	d.Set("updated", resp.Data.Updated)
	d.Set("instance_count", resp.Data.InstanceCount)
	instances := parseInstances(&resp.Data.Instances)
	d.Set("instances", instances)
	d.Set("values", resp.Data.Values)

	return nil
}

func parseInstances(instances *[]vdbv2.TinyDbInstanceInfo) interface{} {
	res := make([]interface{}, len(*instances))
	for i, instance := range *instances {
		cur := make(map[string]interface{})

		cur["id"] = instance.Id
		cur["name"] = instance.Name
		res[i] = cur
	}
	return res
}

func resourceConfigurationGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group delete")

	cli := m.(*client.Client)
	configID := d.Id()
	request := make([]vdb.ConfigurationRequest, 1)

	deleteReq := vdb.ConfigurationRequest{
		Id: configID,
	}
	request[0] = deleteReq
	body, _ := json.Marshal(request)

	_, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.DeleteConfigs1(context.TODO(), string(body))
	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	log.Println("[DEBUG]  Deleted " + d.Id())
	d.SetId("")

	return nil
}

func resourceConfigurationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	//valuesInput, ok := d.Get("values").([]interface{})
	//var values map[string]interface{}
	//if ok {
	//	values = getValues(valuesInput)
	//} else {
	//	values = nil
	//}

	createRequest := vdb.ConfigurationRequest{
		DatastoreType:    d.Get("datastore_type").(string),
		DatastoreVersion: d.Get("datastore_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		Values:           getValues(d.Get("values").(map[string]interface{})),
		EngineGroup:      int32(1),
	}

	body, _ := json.Marshal(createRequest)

	_, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.UpdateConfig1(context.TODO(), string(body))

	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	return nil
}

//func getValues(input []interface{}) map[string]interface{} {
//	values := make(map[string]interface{}, len(input))
//
//	for _, rawValue := range input {
//		valueDto := rawValue.(map[string]interface{})
//		switch valueDto["data_type"].(string) {
//		case "integer":
//			values[valueDto["key"].(string)] = valueDto["value"].(int)
//		case "string":
//			values[valueDto["key"].(string)] = valueDto["value"].(string)
//		case "boolean":
//			values[valueDto["key"].(string)] = valueDto["value"].(bool)
//		}
//	}
//
//	return values
//}

func getValues(input map[string]interface{}) map[string]interface{} {
	values := make(map[string]interface{}, len(input))

	for k, v := range input {
		if intValue, err := strconv.Atoi(v.(string)); err == nil {
			values[k] = intValue
		} else if boolValue, err := strconv.ParseBool(v.(string)); err == nil {
			values[k] = boolValue
		} else {
			values[k] = v.(string)
		}
	}

	return values
}
