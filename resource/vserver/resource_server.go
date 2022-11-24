package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
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
				ForceNew: true,
			},
			"attach_floating": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encryption_volume": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"expire_password": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"root_disk_encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"root_disk_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"root_disk_type_id": {
				Type:     schema.TypeString,
				Required: true,
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
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_password": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"ssh_key_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_disk_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_data": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"base64_encode": {
				Default:  false,
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
func resourceServerStateRefreshFunc(cli *client.Client, serverID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET1(context.TODO(), projectID, serverID)

		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing : %s", GetResponseBody(httpResponse))
		}

		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		server := resp.Data
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
		EncryptionVolume:       d.Get("encryption_volume").(bool),
		ExpirePassword:         d.Get("expire_password").(bool),
		FlavorId:               d.Get("flavor_id").(string),
		ImageId:                d.Get("image_id").(string),
		Name:                   d.Get("name").(string),
		NetworkId:              d.Get("network_id").(string),
		OsLicence:              d.Get("os_licence").(bool),
		RootDiskEncryptionType: d.Get("root_disk_encryption_type").(string),
		RootDiskSize:           int32(d.Get("root_disk_size").(int)),
		RootDiskTypeId:         d.Get("root_disk_type_id").(string),
		SecurityGroup:          securityGroup,
		SshKeyId:               d.Get("ssh_key").(string),
		SubnetId:               d.Get("subnet_id").(string),
		ServerGroupId:          d.Get("server_group_id").(string),
		UserName:               d.Get("user_name").(string),
		UserPassword:           d.Get("user_password").(string),
		UserData:               d.Get("user_data").(string),
		UserDataBase64Encoded:  d.Get("base64_encode").(bool),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.CreateServerUsingPOST1(context.TODO(), server, projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverCreating,
		Target:     serverCreated,
		Refresh:    resourceServerStateRefreshFunc(cli, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create server (%s) %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	serverID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET1(context.TODO(), projectID, serverID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return err
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	server := resp.Data
	d.Set("name", server.Name)
	d.Set("network_id", server.InternalInterfaces[0].NetworkUuid)
	d.Set("subnet_id", server.InternalInterfaces[0].SubnetUuid)
	d.Set("encryption_volume", server.EncryptionVolume)
	d.Set("flavor_id", server.Flavor.FlavorId)
	d.Set("image_id", server.Image.Id)
	d.Set("os_info", server.Image.ImageVersion)
	d.Set("ssh_key_name", server.SshKeyName)
	d.Set("server_group_id", server.ServerGroupId)
	d.Set("root_disk_id", server.BootVolumeId)
	var internalInterfaces []map[string]string
	for _, internalInterface := range server.InternalInterfaces {
		internalInterfaceMap := make(map[string]string)
		internalInterfaceMap["fixed_ip"] = internalInterface.FixedIp
		internalInterfaceMap["floating_ip"] = internalInterface.FloatingIp
		internalInterfaceMap["interface_type"] = internalInterface.InterfaceType
		internalInterfaceMap["mac"] = internalInterface.Mac
		internalInterfaceMap["network_uuid"] = internalInterface.NetworkUuid
		internalInterfaceMap["port_uuid"] = internalInterface.PortUuid
		internalInterfaceMap["product"] = internalInterface.Product
		internalInterfaceMap["server_uuid"] = internalInterface.ServerUuid
		internalInterfaceMap["status"] = internalInterface.Status
		internalInterfaceMap["subnet_uuid"] = internalInterface.SubnetUuid
		internalInterfaceMap["type"] = internalInterface.Type_
		internalInterfaceMap["uuid"] = internalInterface.Uuid
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
		externalInterfaceMap["product"] = externalInterface.Product
		externalInterfaceMap["server_uuid"] = externalInterface.ServerUuid
		externalInterfaceMap["status"] = externalInterface.Status
		externalInterfaceMap["subnet_uuid"] = externalInterface.SubnetUuid
		externalInterfaceMap["type"] = externalInterface.Status
		externalInterfaceMap["uuid"] = externalInterface.Uuid
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
	if d.HasChange("root_disk_size") || d.HasChange("root_disk_type_id") {
		return resourceResizeBootVolume(d, m)
	}
	return resourceServerRead(d, m)

}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	deleteServer := vserver.DeleteServerRequest{
		DeleteAllVolume: false,
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.DeleteServerUsingDELETE1(context.TODO(), deleteServer, projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverDeleting,
		Target:     serverDeleted,
		Refresh:    resourceServerDeleteStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for delete server (%s) : %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
func resourceServerResize(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	serverResizeRequest := vserver.ResizeServerRequest{
		ServerId: d.Id(),
		FlavorId: d.Get("flavor_id").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.ResizeServerUsingPUT1(context.TODO(), projectID, serverResizeRequest, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldFlavor, _ := d.GetChange("flavor_id")
		d.Set("flavor_id", oldFlavor)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverResizing,
		Target:     serverResize,
		Refresh:    resourceServerStateRefreshFunc(cli, d.Id(), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for resize server (%s) : %s", d.Id(), err)
	}
	return resourceServerRead(d, m)
}
func resourceServerReboot(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.RebootServerUsingPUT1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldAction, _ := d.GetChange("action")
		d.Set("action", oldAction)
		return errorResponse
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

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.StopServerUsingPUT1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		d.Set("action", "start")
		return errorResponse
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

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.StartServerUsingPUT1(context.TODO(), projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		d.Set("action", "stop")
		return errorResponse
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
	resp, httpResponse, err := cli.VserverClient.ServerRestControllerApi.UpdateSecGroupServerUsingPUT1(context.TODO(), serverChangeSecGroup, projectID, d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldSecGroups, _ := d.GetChange("security_group")
		d.Set("security_group", oldSecGroups)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    serverChangingSecGroup,
		Target:     serverChangedSecGroup,
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

func resourceResizeBootVolume(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	resizeVolume := vserver.ResizeVolumeRequest{
		NewSize:         int32(d.Get("root_disk_size").(int)),
		NewVolumeTypeId: d.Get("root_disk_type_id").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VolumeRestControllerApi.ResizeVolumeUsingPUT1(context.TODO(), projectID, resizeVolume, d.Get("root_disk_id").(string))
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldSize, _ := d.GetChange("root_disk_size")
		oldType, _ := d.GetChange("root_disk_type_id")
		d.Set("root_disk_size", oldSize)
		d.Set("root_disk_type_id", oldType)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    volumeResizing,
		Target:     volumeAttached,
		Refresh:    resourceVolumeStateRefreshFunc(cli, d.Get("root_disk_id").(string), projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for volume to be resized: %s", err)
	}
	d.Set("root_disk_size", resizeVolume.NewSize)
	d.Set("root_disk_type_id", resizeVolume.NewVolumeTypeId)
	return nil
}

func resourceServerDeleteStateRefreshFunc(cli *client.Client, serverID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.ServerRestControllerApi.GetServerUsingGET1(context.TODO(), projectID, serverID)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.Server{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		server := resp.Data
		return server, server.Status, nil
	}
}
