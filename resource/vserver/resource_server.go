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

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:ServerID", d.Id())
				}
				projectID := idParts[0]
				serverID := idParts[1]
				d.SetId(serverID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attach_floating": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data_disk_encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"data_disk_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_volume_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_volume": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"expire_password": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_licence": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"root_disk_encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"root_disk_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_group": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_interfaces": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			"internal_interfaces": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			"os_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"share": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ssh_key_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceServerStateRefreshFunc(cli *client.Client, serverID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, serverID)
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
		server := resp.Servers[0]
		return server, server.Status, nil
	}
}
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	securityGroupInterface := d.Get("security_group").([]interface{})
	var securityGroup []string
	for _, s := range securityGroupInterface {
		securityGroup = append(securityGroup, s.(string))
	}
	if _, ok := d.GetOk("root_disk_size"); !ok {
		return fmt.Errorf(`The argument "root_disk_size" is required, but no definition was found.`)
	}
	if _, ok := d.GetOk("root_disk_type_id"); !ok {
		return fmt.Errorf(`The argument "root_disk_type_id" is required, but no definition was found.`)
	}
	server := vserver.CreateServerRequest{
		AttachFloating:         d.Get("attach_floating").(bool),
		DataDiskEncryptionType: d.Get("data_disk_encryption_type").(string),
		DataDiskSize:           int32(d.Get("data_disk_size").(int)),
		DataDiskTypeId:         d.Get("data_disk_type_id").(string),
		DataVolumeName:         d.Get("data_volume_name").(string),
		EncryptionVolume:       d.Get("encryption_volume").(bool),
		ExpirePassword:         d.Get("expire_password").(bool),
		FlavorId:               d.Get("flavor_id").(string),
		ImageId:                d.Get("image_id").(string),
		IsPoc:                  d.Get("is_poc").(bool),
		Name:                   d.Get("name").(string),
		NetworkId:              d.Get("network_id").(string),
		OsLicence:              d.Get("os_licence").(bool),
		RootDiskEncryptionType: d.Get("root_disk_encryption_type").(string),
		RootDiskSize:           int32(d.Get("root_disk_size").(int)),
		RootDiskTypeId:         d.Get("root_disk_type_id").(string),
		SecurityGroup:          securityGroup,
		SourceType:             d.Get("source_type").(string),
		SshKeyId:               d.Get("ssh_key").(string),
		SubnetId:               d.Get("subnet_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.CreateServerUsingPOST(context.TODO(), server, projectID)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverCreating,
		Target:     serverCreated,
		Refresh:    resourceServerStateRefreshFunc(cli, resp.Servers[0].Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", resp.Servers[0].Uuid, err)
	}
	d.SetId(resp.Servers[0].Uuid)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	serverID := d.Id()
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, serverID)
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
	if len(resp.Servers) == 0 {
		d.SetId("")
		return nil
	}
	server := resp.Servers[0]
	d.Set("name", server.Name)
	d.Set("network_id", server.InternalInterfaces[0].NetworkUuid)
	d.Set("subnet_id", server.InternalInterfaces[0].SubnetUuid)
	d.Set("encryption_volume", server.EncryptionVolume)
	d.Set("flavor_id", server.FlavorId)
	d.Set("image_id", server.ImageId)
	d.Set("os_info", server.OsInfo)
	d.Set("owner_email", server.OwnerEmail)
	d.Set("share", server.Share)
	d.Set("ssh_key_name", server.SshKeyName)
	var internalInterfaces []map[string]string
	for _, internalInterface := range server.InternalInterfaces {
		internalInterfaceMap := make(map[string]string)
		internalInterfaceMap["fixed_ip"] = internalInterface.FixedIp
		internalInterfaceMap["floating_ip"] = internalInterface.FloatingIp
		internalInterfaceMap["interface_type"] = internalInterface.InterfaceType
		internalInterfaceMap["mac"] = internalInterface.Mac
		internalInterfaceMap["network_uuid"] = internalInterface.NetworkUuid
		internalInterfaceMap["port_uuid"] = internalInterface.PortUuid
		internalInterfaceMap["status"] = internalInterface.Status
		internalInterfaceMap["subnet_uuid"] = internalInterface.SubnetUuid
		internalInterfaceMap["type"] = internalInterface.Status
		internalInterfaces = append(internalInterfaces, internalInterfaceMap)
	}
	d.Set("internal_interfaces", internalInterfaces)
	var externalInterfaces []map[string]string
	for _, externalInterface := range server.ExternalInterfaces {
		externalInterfaceMap := make(map[string]string)
		externalInterfaceMap["fixed_ip"] = externalInterface.FixedIp
		externalInterfaceMap["floating_ip"] = externalInterface.FloatingIp
		externalInterfaceMap["interface_type"] = externalInterface.InterfaceType
		externalInterfaceMap["mac"] = externalInterface.Mac
		externalInterfaceMap["network_uuid"] = externalInterface.NetworkUuid
		externalInterfaceMap["port_uuid"] = externalInterface.PortUuid
		externalInterfaceMap["status"] = externalInterface.Status
		externalInterfaceMap["subnet_uuid"] = externalInterface.SubnetUuid
		externalInterfaceMap["type"] = externalInterface.Status
		externalInterfaces = append(externalInterfaces, externalInterfaceMap)
	}
	d.Set("external_interfaces", externalInterfaces)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("flavor_id") {
		return resourceServerResize(d, m)
	}
	if d.HasChange("action") {
		switch d.Get("action").(string) {
		case "reboot":
			return resourceServerReboot(d, m)
		case "start":
			return resourceServerStart(d, m)
		case "stop":
			return resourceServerStop(d, m)
		}
	}
	if d.HasChange("security_group") {
		return resourceServerUpdateSecgroup(d, m)
	}
	return resourceServerRead(d, m)

}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteServer := vserver.DeleteServerRequest{
		ServerId:    d.Id(),
		ForceDelete: true,
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.DeleteServerInTrashUsingDELETE(context.TODO(), deleteServer, projectID)
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
		resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if len(resp.Servers) == 0 {
			d.SetId("")
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance to be created but was in state %s", resp.Servers[0].Status))
	})
}
func resourceServerResize(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	serverResize := vserver.ResizeServerRequest{
		Poc:      d.Get("is_poc").(bool),
		ServerId: d.Id(),
		FlavorId: d.Get("flavor_id").(string),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.ResizeServerUsingPUT(context.TODO(), projectID, serverResize)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	return resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		resp, _, err := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET(context.TODO(), projectID, d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", err))
		}
		if !resp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error describing instance: %s", resp.ErrorMsg))
		}
		if resp.Servers[0].FlavorId == d.Get("flavor_id").(string) {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("Expected instance FlavorId to be %s but was in state %s", d.Get("flavor_id").(string), resp.Servers[0].FlavorId))
	})
}
func resourceServerReboot(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	serverReboot := vserver.UpdateServerRequest{
		ServerId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.RebootServerUsingPUT(context.TODO(), projectID, serverReboot)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverRebooting,
		Target:     serverRebooted,
		Refresh:    resourceServerStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	return resourceServerRead(d, m)
}
func resourceServerStop(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	serverStop := vserver.UpdateServerRequest{
		ServerId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.StopServerUsingPUT(context.TODO(), projectID, serverStop)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverStopping,
		Target:     serverStopped,
		Refresh:    resourceServerStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	return resourceServerRead(d, m)
}
func resourceServerStart(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	serverStart := vserver.UpdateServerRequest{
		ServerId: d.Id(),
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.StartServerUsingPUT(context.TODO(), projectID, serverStart)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverStarting,
		Target:     serverStarted,
		Refresh:    resourceServerStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	return resourceServerRead(d, m)
}
func resourceServerUpdateSecgroup(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	securityGroupInterface := d.Get("security_group").([]interface{})
	var securityGroup []string
	for _, s := range securityGroupInterface {
		securityGroup = append(securityGroup, s.(string))
	}
	serverChangeSecGroup := vserver.ChangeSecGroupRequest{
		ServerId:      d.Id(),
		SecurityGroup: securityGroup,
	}
	cli := m.(*client.Client)
	resp, _, err := cli.VserverClient.ServerRestControllerApi.UpdateSecGroupServerUsingPUT(context.TODO(), serverChangeSecGroup, projectID)
	if err != nil {
		return err
	}
	if !resp.Success {
		err := fmt.Errorf("request fail with errMsg=%s", resp.ErrorMsg)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	return resourceServerRead(d, m)
}
