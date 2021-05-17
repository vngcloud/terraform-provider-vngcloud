package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vserver"
)

func ResourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: resourceNetworkDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:NetworkID", d.Id())
				}
				projectID := idParts[0]
				networkID := idParts[1]
				d.SetId(networkID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table_default_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceNetworkStateRefreshFunc(cli *client.Client, networkID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET(context.TODO(), networkID, projectID)
		if err != nil {
			return nil, "", fmt.Errorf("Error on network State Refresh: %s", err)
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			return nil, "", fmt.Errorf("Error describing instance: %s", resp.ErrorMsg)
		}
		network := resp.Networks[0]
		return network, network.Status, nil
	}
}
func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	network := vserver.CreateNetworkRequest{
		Name:                d.Get("name").(string),
		Cidr:                d.Get("cidr").(string),
		RouteTableDefaultId: d.Get("route_table_default_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.NetworkRestControllerApi.CreateNetworkUsingPOST(context.TODO(), network, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:    networkCreating,
		Target:     networkCreated,
		Refresh:    resourceNetworkStateRefreshFunc(cli, resp.Networks[0].Id, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Networks[0].Id, err)
	}
	d.SetId(resp.Networks[0].Id)
	return resourceNetworkRead(d, m)
}

func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET(context.TODO(), networkID, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	if len(resp.Networks) == 0 {
		d.SetId("")
		return nil
	}
	network := resp.Networks[0]
	d.Set("name", network.Name)
	d.Set("cidr", network.Cidr)
	return nil
}

func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkID := d.Id()
	cli := m.(*client.Client)
	if d.HasChange("name") {
		networkUpdate := vserver.UpdateNetworkRequest{
			Name:      d.Get("name").(string),
			NetworkId: networkID,
		}
		resp, _, err := cli.VserverClient.NetworkRestControllerApi.EditNetworkUsingPUT(context.TODO(), projectID, networkUpdate)
		if err != nil {
			return err
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		if !resp.Success {
			err := fmt.Errorf(resp.ErrorMsg)
			return err
		}
		return resourceNetworkRead(d, m)
	}
	return nil
}

func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteNetwork := vserver.DeleteNetworkRequest{
		NetworkId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.NetworkRestControllerApi.DeleteNetworkUsingDELETE(context.TODO(), deleteNetwork, projectID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	return resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.NetworkRestControllerApi.GetNetworkUsingGET(context.TODO(), d.Id(), projectID)

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Networks) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Networks[0].Status))
	})
}
