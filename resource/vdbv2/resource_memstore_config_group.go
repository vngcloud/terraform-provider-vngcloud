package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
)

func ResourceMemStoreConfigurationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemStoreConfigurationGroupCreate,
		Read:   resourceMemStoreConfigurationGroupRead,
		Delete: resourceMemStoreConfigurationGroupDelete,
		Update: resourceMemStoreConfigurationGroupUpdate,

		Schema: map[string]*schema.Schema{
			"engine_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_version": {
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

func resourceMemStoreConfigurationGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group create")

	cli := m.(*client.Client)
	createRequest := vdb.ConfigurationRequest{
		DatastoreType:    d.Get("engine_type").(string),
		DatastoreVersion: d.Get("engine_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		EngineGroup:      2,
	}

	body, _ := json.Marshal(createRequest)

	resp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreConfigurationGroupAPIApi.CreateConfig(context.TODO(), string(body))

	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	log.Println("[DEBUG]  Created " + resp.Data.Id)
	time.Sleep(10 * time.Second)
	d.SetId(resp.Data.Id)

	err := resourceMemStoreConfigurationGroupUpdate(d, m)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Second)

	return resourceMemStoreConfigurationGroupRead(d, m)
}

func resourceMemStoreConfigurationGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group read")

	cli := m.(*client.Client)
	configID := d.Id()

	resp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreConfigurationGroupAPIApi.GetConfigsById(context.TODO(), configID)
	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if resp.Data == nil {
		d.SetId("")
		return nil
	}

	d.Set("engine_type", resp.Data.DatastoreName)
	d.Set("engine_version", resp.Data.DatastoreVersionName)
	d.Set("description", resp.Data.Description)
	d.Set("name", resp.Data.Name)
	d.Set("created", resp.Data.Created)
	d.Set("updated", resp.Data.Updated)
	d.Set("instance_count", resp.Data.InstanceCount)
	instances := parseInstances(&resp.Data.Instances)
	d.Set("instances", instances)

	return nil
}

func resourceMemStoreConfigurationGroupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Configuration group delete")

	cli := m.(*client.Client)
	configID := d.Id()
	request := make([]vdb.ConfigurationRequest, 1)

	deleteReq := vdb.ConfigurationRequest{
		Id: configID,
	}
	request[0] = deleteReq
	body, _ := json.Marshal(request)

	resp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreConfigurationGroupAPIApi.DeleteConfigs(context.TODO(), string(body))
	//if err != nil {
	//	return err
	//}

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
		Refresh:    resourceMemStoreConfigGroupDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for delete config group (%s) : %s", d.Id(), err)
	}
	d.SetId("")

	return nil
}

func resourceMemStoreConfigurationGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	//valuesInput, ok := d.Get("values").([]interface{})
	//var values map[string]interface{}
	//if ok {
	//	values = getValues(valuesInput)
	//} else {
	//	values = nil
	//}

	values := getValues(d.Get("values").(map[string]interface{}))

	if len(values) == 0 {
		log.Println("Values empty to be updated")
		return nil
	}

	updateRequest := vdb.ConfigurationRequest{
		Id:               d.Id(),
		DatastoreType:    d.Get("engine_type").(string),
		DatastoreVersion: d.Get("engine_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
		Values:           values,
		EngineGroup:      1,
	}

	body, _ := json.Marshal(updateRequest)

	_, httpResponse, _ := cli.Vdbv2Client.MemoryStoreConfigurationGroupAPIApi.UpdateConfig(context.TODO(), string(body))

	//if err != nil {
	//	return err
	//}

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

func resourceMemStoreConfigGroupDeleteStateRefreshFunc(cli *client.Client, configId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreConfigurationGroupAPIApi.GetConfigsById(context.TODO(), configId)
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
