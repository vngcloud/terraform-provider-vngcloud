package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
)

func ResourceRelationalConfigurationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRelationalConfigurationGroupCreate,
		Read:   resourceRelationalConfigurationGroupRead,
		Delete: resourceRelationalConfigurationGroupDelete,
		Update: resourceRelationalConfigurationGroupUpdate,

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

func resourceRelationalConfigurationGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group create")

	cli := m.(*client.Client)
	createRequest := vdb.ConfigurationRequest{
		DatastoreType:    d.Get("datastore_type").(string),
		DatastoreVersion: d.Get("datastore_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		EngineGroup:      1,
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
	time.Sleep(10 * time.Second)
	d.SetId(resp.Data.Id)

	_, ok := d.Get("values").([]interface{})
	if ok {
		err := resourceRelationalConfigurationGroupUpdate(d, m)
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Second)
	}
	return resourceRelationalConfigurationGroupRead(d, m)
}

func resourceRelationalConfigurationGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group read")

	cli := m.(*client.Client)
	configID := d.Id()

	resp, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.GetConfigsById1(context.TODO(), configID)
	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if resp.Data == nil {
		d.SetId("")
		return nil
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

func resourceRelationalConfigurationGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group delete")

	cli := m.(*client.Client)
	configID := d.Id()
	request := make([]vdb.ConfigurationRequest, 1)

	deleteReq := vdb.ConfigurationRequest{
		Id: configID,
	}
	request[0] = deleteReq
	body, _ := json.Marshal(request)

	resp, httpResponse, err := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.DeleteConfigs1(context.TODO(), string(body))
	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if resp.Data[0].Success == false {
		errorResponse := fmt.Errorf("request fail with errMsg : %s", resp.Data[0].ErrorMsg)
		return errorResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    resourceRelationalConfigGroupDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for delete config group (%s) : %s", d.Id(), err)
	}
	d.SetId("")

	return nil
}

func resourceRelationalConfigurationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	//valuesInput, ok := d.Get("values").([]interface{})
	//var values map[string]interface{}
	//if ok {
	//	values = getValues(valuesInput)
	//} else {
	//	values = nil
	//}

	updateRequest := vdb.ConfigurationRequest{
		Id:               d.Id(),
		DatastoreType:    d.Get("datastore_type").(string),
		DatastoreVersion: d.Get("datastore_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		Values:           getValues(d.Get("values").(map[string]interface{})),
		EngineGroup:      1,
	}

	body, _ := json.Marshal(updateRequest)

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

func resourceRelationalConfigGroupDeleteStateRefreshFunc(cli *client.Client, configId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalConfigurationGroupAPIApi.GetConfigsById1(context.TODO(), configId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vdbv2.ItemConfigInfo{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing config group: %s", GetResponseBody(httpResponse))
			}
		} else {
			if dbResp.Data == nil {
				return vdbv2.ItemConfigInfo{Status: "DELETED"}, "DELETED", nil
			}
		}
		respJSON, _ := json.Marshal(dbResp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		return vdbv2.ItemConfigInfo{Status: "ACTIVE"}, "ACTIVE", nil
	}
}
