package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
)

func ResourceChangeClusterSecGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAddSecGroup,
		Delete: resourceRemoveSecGroup,
		Read:   resourceReadSecGroup,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"master": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"sec_group_id_list": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sec_group_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeString,
						},
						"name": {
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func resourceAddSecGroup(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterId := d.Get("cluster_id").(string)
	var typeOfNode string
	if d.Get("master").(bool) == true {
		typeOfNode = "master"
	} else {
		typeOfNode = "minion"
	}
	cli := m.(*client.Client)
	var secGroupIdList []string
	for _, secGroup := range d.Get("sec_group_id_list").([]string) {
		secGroupIdList = append(secGroupIdList, secGroup)
	}

	request := vserver.UpdateClusterSecGroupRequest{
		ClusterId:   clusterId,
		Master:      d.Get("master").(bool),
		SecGroupIds: secGroupIdList,
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.UpdateSecGroupUsingPUT(context.TODO(), projectID, request)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return fmt.Errorf("error attaching secgroup for %s cluster (%s) %s", typeOfNode, clusterId, err)
	}

	d.SetId(clusterId)
	return nil
}

func resourceRemoveSecGroup(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterId := d.Get("cluster_id").(string)
	var typeOfNode string
	if d.Get("master").(bool) == true {
		typeOfNode = "master"
	} else {
		typeOfNode = "minion"
	}
	cli := m.(*client.Client)
	var secGroupIdList []string

	request := vserver.UpdateClusterSecGroupRequest{
		ClusterId:   clusterId,
		Master:      d.Get("master").(bool),
		SecGroupIds: secGroupIdList,
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.UpdateSecGroupUsingPUT(context.TODO(), projectID, request)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return fmt.Errorf("error attaching secgroup for %s cluster (%s) %s", typeOfNode, clusterId, err)
	}

	d.SetId(clusterId)
	return nil
}

func resourceReadSecGroup(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterId := d.Get("cluster_id").(string)
	isMaster := d.Get("master").(bool)
	var typeOfNode string
	if isMaster == true {
		typeOfNode = "master"
	} else {
		typeOfNode = "minion"
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.ListSecGroupUsingGET(context.TODO(), clusterId, projectID, isMaster)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return fmt.Errorf("error get list secgroup for %s cluster (%s) %s", typeOfNode, clusterId, err)
	}

	var secGroupList []map[string]string
	var secGroupIdList []string
	for _, s := range resp {
		m := make(map[string]string)
		m["id"] = s.SecGroupId
		m["name"] = s.SecGroupName
		secGroupIdList = append(secGroupIdList, s.SecGroupId)
		secGroupList = append(secGroupList, m)
	}

	d.Set("cluster_id", clusterId)
	d.Set("master", resp[0].Master)
	d.Set("sec_group_id_list", secGroupIdList)
	d.Set("sec_group_list", secGroupList)

	d.SetId(clusterId)
	return nil
}
