package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
	"log"
	"net/http"
)

func ResourceMemStoreDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemStoreDatabaseCreate,
		Read:   resourceMemStoreDatabaseRead,
		Update: resourceMemStoreDatabaseUpdate,
		Delete: resourceMemStoreDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_auto": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"backup_duration": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"backup_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"free_backup_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"public_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"redis_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redis_password_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"replica_source_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_type_zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_used": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"zone_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed_ip_prefix": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMemStoreDatabaseStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG]  State refresh")

		dbResp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), id)
		if err != nil {
			return nil, "", fmt.Errorf("error when refreshing database state: %s", err)
		}
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return nil, "", fmt.Errorf("error when refreshing database state: %s", responseBody)
		}
		if dbResp.Data == nil {
			return nil, "", fmt.Errorf("error when refreshing database state: data is empty")
		}
		log.Println("[DEBUG]  Database status: " + dbResp.Data.Status)

		return dbResp.Data.Id, dbResp.Data.Status, nil
	}
}

func resourceMemStoreDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database create")

	cli := m.(*client.Client)

	var instanceId string
	if d.Get("backup_id").(string) == "" {
		createRequest := generateMemStoreCreateDatabaseRequest(d)
		reqBody, _ := json.Marshal(createRequest)
		log.Println("[DEBUG]  Body: " + string(reqBody))
		if d.Get("replica_source_id").(string) == "" {
			createDbResult, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.CreateMemoryStoreDatabaseInstance(context.TODO(), string(reqBody), nil)
			if err != nil {
				return fmt.Errorf("error when creating database: %s", err)
			}
			if CheckErrorResponse(httpResponse) {
				responseBody := GetResponseBody(httpResponse)
				errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
				return errorResponse
			}
			log.Println("[DEBUG]  Create db result, id: " + createDbResult.Data[0].ResourceId)
			instanceId = createDbResult.Data[0].ResourceId
		} else {
			createDbResult, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.CreateDatabaseInstanceReplica(context.TODO(), string(reqBody), d.Get("replica_source_id").(string), nil)
			if err != nil {
				return fmt.Errorf("error when creating database replica: %s", err)
			}
			if CheckErrorResponse(httpResponse) {
				responseBody := GetResponseBody(httpResponse)
				errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
				return errorResponse
			}
			log.Println("[DEBUG]  Create db replica result, id: " + createDbResult.Data[0].ResourceId)
			instanceId = createDbResult.Data[0].ResourceId
		}
	} else {
		restoreRequest := generateMemStoreRestoreBackupRequest(d)
		reqBody, _ := json.Marshal(restoreRequest)
		log.Println("[DEBUG]  Body: " + string(reqBody))

		createDbResult, httpResponse, err := cli.Vdbv2Client.MemoryStoreBackupAPIApi.RestoreBackup(context.TODO(), string(reqBody), d.Get("backup_id").(string), nil)
		if err != nil {
			return fmt.Errorf("error when creating database: %s", err)
		}
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		log.Println("[DEBUG]  Restore bk result, id: " + createDbResult.Data[0].ResourceId)
		instanceId = createDbResult.Data[0].ResourceId
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseCreatePending,
		Target:     databaseCreateTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, instanceId),
		Timeout:    databaseCreateTimeout,
		Delay:      databaseCreateDelay,
		MinTimeout: databaseCreateMinTimeout,
	}

	id, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be created: %s", err)
	}
	idStr := id.(string)

	log.Println("[DEBUG]  Wait for state done, id: " + idStr)

	d.SetId(idStr)

	allowedIP := d.Get("allowed_ip_prefix").([]interface{})
	log.Printf("[DEBUG] allowed ip prefix: %v\n", allowedIP)

	err = resourceMemStoreDatabaseRead(d, m)
	if err != nil {
		return fmt.Errorf("error when getting database info after create: %s", err)
	}

	if allowedIP != nil && len(allowedIP) > 0 {
		err := secgroupRulesMemStoreUpdate(d, m, allowedIP)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateMemStoreCreateDatabaseRequest(d *schema.ResourceData) vdb.CreateDbInstanceRequest {
	createRequest := vdb.CreateDbInstanceRequest{
		BackupAuto:           d.Get("backup_auto").(bool),
		BackupDuration:       int32(d.Get("backup_duration").(int)),
		BackupTime:           d.Get("backup_time").(string),
		ConfigId:             d.Get("config_id").(string),
		DatastoreType:        d.Get("engine_type").(string),
		DatastoreVersion:     d.Get("engine_version").(string),
		Name:                 d.Get("name").(string),
		NetIds:               []string{d.Get("subnet_id").(string)},
		PackageId:            d.Get("package_id").(string),
		Poc:                  d.Get("is_poc").(bool),
		PublicAccess:         d.Get("public_access").(bool),
		ReplicaSourceId:      d.Get("replica_source_id").(string),
		RedisPasswordEnabled: false,
		RedisPassword:        "",
	}

	if d.Get("replica_source_id").(string) == "" {
		createRequest.RedisPasswordEnabled = d.Get("redis_password_enabled").(bool)
		createRequest.RedisPassword = d.Get("redis_password").(string)
	}

	return createRequest
}

func secgroupRulesMemStoreUpdate(d *schema.ResourceData, m interface{}, allowedIP []interface{}) error {
	if len(allowedIP) > 0 {
		cli := m.(*client.Client)

		instanceID := d.Id()
		rules := createSecurityGroupRules(&allowedIP, d.Get("port").(int))
		reqBody, _ := json.Marshal(rules)
		_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.UpdateSecurityRules(context.TODO(), string(reqBody), instanceID)

		if err != nil {
			return err
		}

		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
	}

	return secgroupRulesMemStoreRead(d, m)
}

func secgroupRulesMemStoreRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	instanceID := d.Id()
	resp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.GetSecurityRules1(context.TODO(), instanceID)

	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if resp.Data == nil || len(resp.Data) == 0 {
		return nil
	}

	rules := make([]interface{}, len(resp.Data))
	for i, securityGroupRule := range resp.Data {
		rules[i] = securityGroupRule.RemoteIpPrefix
	}

	log.Printf("[DEBUG] reading rules %v\n", rules)
	d.Set("allowed_ip_prefix", rules)

	return nil
}

func resourceMemStoreDatabaseRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database read")

	cli := m.(*client.Client)

	dbResp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), d.Id())
	if err != nil {
		return fmt.Errorf("error when getting database info: %s", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if dbResp.Data == nil {
		d.SetId("")
		return nil
	}

	d.Set("backup_auto", dbResp.Data.BackupAuto)
	d.Set("backup_duration", dbResp.Data.BackupDuration)
	d.Set("backup_time", dbResp.Data.BackupTime)
	d.Set("engine_type", dbResp.Data.DatastoreType)
	d.Set("engine_version", dbResp.Data.DatastoreVersion)
	d.Set("free_backup_size", dbResp.Data.FreeBackupSize)
	d.Set("ip", dbResp.Data.Ip)
	d.Set("name", dbResp.Data.Name)
	d.Set("subnet_id", dbResp.Data.SubnetId)
	d.Set("package_id", dbResp.Data.QuotaPackageId)
	d.Set("port", dbResp.Data.Port)
	d.Set("public_access", dbResp.Data.PublicAccess)
	d.Set("ram", dbResp.Data.Ram)
	d.Set("replicas", dbResp.Data.Replicas)
	d.Set("status", dbResp.Data.Status)
	d.Set("cpu", dbResp.Data.Vcpus)
	d.Set("volume_size", dbResp.Data.VolumeSize)
	d.Set("volume_type", dbResp.Data.VolumeType)
	d.Set("volume_type_zone_id", dbResp.Data.VolumeTypeZoneId)
	d.Set("volume_used", dbResp.Data.VolumeUsed)
	d.Set("zone_uuid", dbResp.Data.ZoneUUID)

	if dbResp.Data.Configuration != nil {
		d.Set("config_id", dbResp.Data.Configuration.Id)
		d.Set("config_name", dbResp.Data.Configuration.Name)
	}

	return secgroupRulesMemStoreRead(d, m)
}

func resourceMemStoreDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")
	if d.HasChange("allowed_ip_prefix") {
		return secgroupRulesMemStoreUpdate(d, m, d.Get("allowed_ip_prefix").([]interface{}))
	}

	if d.HasChange("action") {
		switch d.Get("action").(string) {
		case "start":
			return resourceMemStoreDatabaseStart(d, m)
		case "stop":
			return resourceMemStoreDatabaseStop(d, m)
		case "reboot":
			return resourceMemStoreDatabaseReboot(d, m)
		}
	}

	if d.HasChange("config_id") {
		return resourceMemStoreUpdateConfigGroup(d, m)
	}

	if d.HasChange("redis_password_enabled") || (d.Get("redis_password_enabled").(bool) == true && d.HasChange("redis_password")) || d.HasChange("public_access") || d.HasChange("backup_auto") || (d.Get("backup_auto").(bool) == true && (d.HasChange("backup_duration") || d.HasChange("backup_time"))) {
		return resourceMemStoreUpdateSetting(d, m)
	}

	if d.HasChange("replica_source_id") && d.Get("replica_source_id").(string) == "" {
		return resourceMemStoreDatabasePromote(d, m)
	}

	if d.HasChange("package_id") {
		return resourceMemStoreResizeFlavor(d, m)
	}

	return nil
}

func resourceMemStoreUpdateConfigGroup(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateUpdateConfigGroupRequest(d)
	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.UpdateDatabaseConfigGroup(context.TODO(), string(reqBody), d.Id())

	if err != nil {
		return err
	}

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
		Pending:    databaseResizePending,
		Target:     databaseResizeTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseResizeTimeout,
		Delay:      databaseResizeDelay,
		MinTimeout: databaseResizeMinTimeout,
	}

	id, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be updated: %s", err)
	}
	idStr := id.(string)

	log.Println("[DEBUG]  Wait for state done, id: " + idStr)

	return resourceMemStoreDatabaseRead(d, m)
}

func resourceMemStoreUpdateSetting(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateUpdateSettingRequest(d)

	if d.HasChange("redis_password_enabled") || (d.Get("redis_password_enabled").(bool) == true && d.HasChange("redis_password")) {
		updateRequest.EditRedisPassword = true
		updateRequest.RedisPasswordEnabled = d.Get("redis_password_enabled").(bool)
		updateRequest.RedisPassword = d.Get("redis_password").(string)
	}

	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.UpdateDatabaseSetting(context.TODO(), string(reqBody), d.Id())

	if err != nil {
		return err
	}

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
		Pending:    databaseResizePending,
		Target:     databaseResizeTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseResizeTimeout,
		Delay:      databaseResizeDelay,
		MinTimeout: databaseResizeMinTimeout,
	}

	id, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be updated: %s", err)
	}
	idStr := id.(string)

	log.Println("[DEBUG]  Wait for state done, id: " + idStr)

	return resourceMemStoreDatabaseRead(d, m)
}

func resourceMemStoreResizeFlavor(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateResizeFlavorDatabaseRequest(d)
	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.ResizeInstance(context.TODO(), string(reqBody), d.Id(), nil)

	if err != nil {
		return err
	}

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
		Pending:    databaseResizePending,
		Target:     databaseResizeTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseResizeTimeout,
		Delay:      databaseResizeDelay,
		MinTimeout: databaseResizeMinTimeout,
	}

	id, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be resized: %s", err)
	}
	idStr := id.(string)

	log.Println("[DEBUG]  Wait for state done, id: " + idStr)

	d.Set("volume_size", updateRequest.DatabaseInstances[0].Config.VolumeSize)
	d.Set("volume_type", updateRequest.DatabaseInstances[0].Config.VolumeType)

	return nil
}

func resourceMemStoreDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database delete")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "delete")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.DeleteDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if err != nil {
		return fmt.Errorf("error when deleting database: %s", err)
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseDeletePending,
		Target:     databaseDeleteTarget,
		Refresh:    resourceMemStoreDatabaseDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseDeleteTimeout,
		Delay:      databaseDeleteDelay,
		MinTimeout: databaseDeleteMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be deleted: %s", err)
	}

	log.Println("[DEBUG]  Delete done")

	d.SetId("")
	return nil
}

func resourceMemStoreDatabaseDeleteStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), id)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vdbv2.DbInstanceInfo{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing config group: %s", GetResponseBody(httpResponse))
			}
		} else {
			if dbResp.Data == nil {
				return vdbv2.DbInstanceInfo{Status: "DELETED"}, "DELETED", nil
			}
		}
		respJSON, _ := json.Marshal(dbResp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		return dbResp.Data, dbResp.Data.Status, nil
	}
}

func resourceMemStoreDatabaseStart(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database start")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "start")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.StartDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if err != nil {
		return fmt.Errorf("error when starting database: %s", err)
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStartPending,
		Target:     databaseStartTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStartTimeout,
		Delay:      databaseStartDelay,
		MinTimeout: databaseStartMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be started: %s", err)
	}

	return resourceMemStoreDatabaseRead(d, m)
}

func resourceMemStoreDatabaseStop(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database stop")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "stop")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.StopDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if err != nil {
		return fmt.Errorf("error when stopping database: %s", err)
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStopPending,
		Target:     databaseStopTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStopTimeout,
		Delay:      databaseStopDelay,
		MinTimeout: databaseStopMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be stopped: %s", err)
	}

	return resourceMemStoreDatabaseRead(d, m)
}

func resourceMemStoreDatabaseReboot(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database reboot")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "reboot")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.RestartDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if err != nil {
		return fmt.Errorf("error when rebooting database: %s", err)
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseRebootPending,
		Target:     databaseRebootTarget,
		Refresh:    resourceMemStoreDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseRebootTimeout,
		Delay:      databaseRebootDelay,
		MinTimeout: databaseRebootMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be rebooted: %s", err)
	}

	return resourceMemStoreDatabaseRead(d, m)
}

func resourceMemStoreDatabasePromote(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database promote")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "detach_replica")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.DetachReplica(context.TODO(), string(reqBody), d.Id())
	if err != nil {
		return fmt.Errorf("error when promoting database: %s", err)
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	err = resource.Retry(databasePromoteTimeout, func() *resource.RetryError {
		dbResp, httpResponse, err := cli.Vdbv2Client.MemoryStoreDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), d.Id())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error when refreshing database state: %s", err))
		}
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return resource.NonRetryableError(fmt.Errorf("error when refreshing database state: %s", responseBody))
		}
		if dbResp.Data == nil {
			return resource.NonRetryableError(fmt.Errorf("error when refreshing database state: data is empty"))
		}

		if dbResp.Data.ReplicaSourceId != "" {
			return resource.RetryableError(fmt.Errorf("database is still not promoted"))
		}

		log.Println("[DEBUG]  Promote done")

		return nil
	})

	if err != nil {
		return fmt.Errorf("error when waiting for database to be promoted: %s", err)
	}

	return resourceMemStoreDatabaseRead(d, m)
}
