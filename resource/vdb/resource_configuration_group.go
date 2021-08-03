package vdb

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vdb"
)

func ResourceConfigurationGroup() *schema.Resource {
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
		},
	}
}

func resourceConfigurationGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [nhannt10] Configuration group create")

	cli := m.(*client.Client)
	projectID := cli.ProjectId
	createRequest := vdb.ConfigurationRequest{
		DatastoreType:    d.Get("datastore_type").(string),
		DatastoreVersion: d.Get("datastore_version").(string),
		Description:      d.Get("description").(string),
		Name:             d.Get("name").(string),
	}

	resp, _, err := cli.VdbClient.VdbConfigurationGroupEndPointApi.CreateConfigUsingPOST(context.TODO(), projectID, createRequest)

	if err != nil {
		return err
	}

	if !resp.Success {
		err = errors.New(resp.ErrorMsg)
		return err
	}
	log.Println("[DEBUG] [nhannt10] Created " + resp.Data.Id)

	d.SetId(resp.Data.Id)
	resourceConfigurationGroupRead(d, m)

	return nil
}

func resourceConfigurationGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [nhannt10] Configuration group read")

	cli := m.(*client.Client)
	configID := d.Id()
	projectID := cli.ProjectId

	resp, _, err := cli.VdbClient.VdbConfigurationGroupEndPointApi.GetConfigByIdUsingGET(context.TODO(), configID, projectID)
	if err != nil {
		return err
	}

	if resp.Data == nil {
		d.SetId("")
		return nil
	}

	if !resp.Success {
		return errors.New(resp.ErrorMsg)
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

func parseInstances(instances *[]vdb.DbInstancePartialInfo) interface{} {
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
	log.Println("[DEBUG] [nhannt10] Configuration group delete")

	cli := m.(*client.Client)
	configID := d.Id()
	projectID := cli.ProjectId
	request := vdb.ConfigurationRequest{}

	resp, _, err := cli.VdbClient.VdbConfigurationGroupEndPointApi.DeleteConfigUsingDELETE(context.TODO(), configID, projectID, request)
	if err != nil {
		return err
	}

	if !resp.Success {
		err = errors.New(resp.ErrorMsg)
		return err
	}

	log.Println("[DEBUG] [nhannt10] Deleted " + d.Id())
	d.SetId("")

	return nil
}
