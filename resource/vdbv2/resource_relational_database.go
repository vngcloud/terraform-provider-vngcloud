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

func ResourceRelationalDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceRelationalDatabaseCreate,
		Read:   resourceRelationalDatabaseRead,
		Update: resourceRelationalDatabaseUpdate,
		Delete: resourceRelationalDatabaseDelete,
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
				Optional: true,
			},
			"backup_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"password": {
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
			"replica_source_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceRelationalDatabaseStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG]  State refresh")

		dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById1(context.TODO(), id)
		//if err != nil {
		//	return nil, "", fmt.Errorf("error when refreshing database state: %s", err)
		//}
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

func resourceRelationalDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database create")

	cli := m.(*client.Client)

	var instanceId string
	if d.Get("backup_id").(string) == "" {
		createRequest := generateRelationalCreateDatabaseRequest(d)
		reqBody, _ := json.Marshal(createRequest)
		log.Println("[DEBUG]  Body: " + string(reqBody))

		if d.Get("replica_source_id").(string) == "" {
			createDbResult, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.CreateRelationalDatabaseInstance(context.TODO(), string(reqBody), nil)
			//if err != nil {
			//	return fmt.Errorf("error when creating database: %s", err)
			//}
			if CheckErrorResponse(httpResponse) {
				responseBody := GetResponseBody(httpResponse)
				errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
				return errorResponse
			}
			log.Println("[DEBUG]  Create db result, id: " + createDbResult.Data[0].ResourceId)
			instanceId = createDbResult.Data[0].ResourceId
		} else {
			createDbResult, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.CreateRelationalDatabaseInstanceReplica(context.TODO(), string(reqBody), d.Get("replica_source_id").(string), nil)
			//if err != nil {
			//	return fmt.Errorf("error when creating database replica: %s", err)
			//}
			if CheckErrorResponse(httpResponse) {
				responseBody := GetResponseBody(httpResponse)
				errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
				return errorResponse
			}
			log.Println("[DEBUG]  Create db replica result, id: " + createDbResult.Data[0].ResourceId)
			instanceId = createDbResult.Data[0].ResourceId
		}
	} else {
		restoreRequest := generateRelationalRestoreBackupRequest(d)
		reqBody, _ := json.Marshal(restoreRequest)
		log.Println("[DEBUG]  Body: " + string(reqBody))

		createDbResult, httpResponse, _ := cli.Vdbv2Client.RelationalBackupAPIApi.RestoreBackup1(context.TODO(), string(reqBody), d.Get("backup_id").(string), nil)
		//if err != nil {
		//	return fmt.Errorf("error when creating database: %s", err)
		//}
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
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, instanceId),
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

	err = resourceRelationalDatabaseRead(d, m)
	if err != nil {
		return fmt.Errorf("error when getting database info after create: %s", err)
	}

	if allowedIP != nil && len(allowedIP) > 0 {
		err := secgroupRulesRelationalUpdate(d, m, allowedIP)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRelationalCreateDatabaseRequest(d *schema.ResourceData) vdb.CreateDbInstanceRequest {
	createRequest := vdb.CreateDbInstanceRequest{
		BackupAuto:       d.Get("backup_auto").(bool),
		BackupDuration:   int32(d.Get("backup_duration").(int)),
		BackupTime:       d.Get("backup_time").(string),
		ConfigId:         d.Get("config_id").(string),
		Databases:        nil,
		DatastoreType:    d.Get("engine_type").(string),
		DatastoreVersion: d.Get("engine_version").(string),
		Name:             d.Get("name").(string),
		NetIds:           []string{d.Get("subnet_id").(string)},
		PackageId:        d.Get("package_id").(string),
		Poc:              d.Get("is_poc").(bool),
		PublicAccess:     d.Get("public_access").(bool),
		ReplicaSourceId:  d.Get("replica_source_id").(string),
		User:             nil,
		VolumeSize:       int32(d.Get("volume_size").(int)),
		VolumeType:       d.Get("volume_type").(string),
	}

	if d.Get("replica_source_id").(string) == "" {
		createRequest.User = &vdb.UserRequest{
			Databases: []vdb.Database{{Name: d.Get("db_name").(string)}},
			Name:      d.Get("username").(string),
			Password:  d.Get("password").(string),
		}

		createRequest.Databases = []vdb.DatabaseRequest{{
			CharacterSet: "utf8",
			Collate:      "utf8_general_ci",
			Name:         d.Get("db_name").(string),
		}}
	}

	return createRequest
}

func secgroupRulesRelationalUpdate(d *schema.ResourceData, m interface{}, allowedIP []interface{}) error {
	cli := m.(*client.Client)

	instanceID := d.Id()
	rules := []vdbv2.SecurityGroupRuleEntity{}
	if len(allowedIP) > 0 {
		rules = createSecurityGroupRules(&allowedIP, d.Get("port").(int))
	}
	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.UpdateSecurityRules(context.TODO(), rules, instanceID)

	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	return secgroupRulesRelationalRead(d, m)
}

func createSecurityGroupRules(allowedIP *[]interface{}, port int) []vdbv2.SecurityGroupRuleEntity {
	rules := make([]vdbv2.SecurityGroupRuleEntity, len(*allowedIP))
	for i, rule := range *allowedIP {
		rules[i].RemoteIpPrefix = rule.(string)
		rules[i].PortRangeMin = int32(port)
		rules[i].PortRangeMax = int32(port)
	}
	return rules
}

func secgroupRulesRelationalRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	instanceID := d.Id()
	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetSecurityRules(context.TODO(), instanceID)

	//if err != nil {
	//	return err
	//}

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

func resourceRelationalDatabaseRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database read")

	cli := m.(*client.Client)

	dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById1(context.TODO(), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when getting database info: %s", err)
	//}

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
	//d.Set("redis_password_enabled", dbResp.Data.RedisPasswordEnabled)
	d.Set("replica_source_id", dbResp.Data.ReplicaSourceId)
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

	return secgroupRulesRelationalRead(d, m)
}

func resourceRelationalDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")
	if d.HasChange("allowed_ip_prefix") {
		return secgroupRulesRelationalUpdate(d, m, d.Get("allowed_ip_prefix").([]interface{}))
	}

	if d.HasChange("action") {
		switch d.Get("action").(string) {
		case "start":
			return resourceRelationalDatabaseStart(d, m)
		case "stop":
			return resourceRelationalDatabaseStop(d, m)
		case "reboot":
			return resourceRelationalDatabaseReboot(d, m)
		}
	}

	if d.HasChange("config_id") {
		return resourceRelationalUpdateConfigGroup(d, m)
	}

	if d.HasChange("password") || d.HasChange("public_access") || d.HasChange("backup_auto") || (d.Get("backup_auto").(bool) == true && (d.HasChange("backup_duration") || d.HasChange("backup_time"))) {
		return resourceRelationalUpdateSetting(d, m)
	}

	if d.HasChange("replica_source_id") && d.Get("replica_source_id").(string) == "" {
		return resourceRelationalDatabasePromote(d, m)
	}

	if d.HasChange("volume_size") || d.HasChange("volume_type") {
		return resourceRelationalResizeVolume(d, m)
	}

	if d.HasChange("package_id") {
		return resourceRelationalResizeFlavor(d, m)
	}

	return nil
}

func resourceRelationalResizeVolume(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateResizeVolumeDatabaseRequest(d)
	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.ResizeStorage(context.TODO(), string(reqBody), d.Id(), nil)

	//if err != nil {
	//	return err
	//}

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
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
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

	d.Set("package_id", updateRequest.DatabaseInstances[0].Config.PackageID)

	return nil
}

func resourceRelationalUpdateConfigGroup(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateUpdateConfigGroupRequest(d)
	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.UpdateDatabaseConfigGroup1(context.TODO(), string(reqBody), d.Id())

	//if err != nil {
	//	return err
	//}

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
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
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

	return resourceRelationalDatabaseRead(d, m)
}

func resourceRelationalUpdateSetting(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateUpdateSettingRequest(d)

	if d.HasChange("password") {
		updateRequest.Password = d.Get("password").(string)
	}
	if d.HasChange("public_access") {
		publicAccess := d.Get("public_access").(bool)
		updateRequest.PublicAccess = &publicAccess
	}

	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.UpdateDatabaseSetting1(context.TODO(), string(reqBody), d.Id())

	//if err != nil {
	//	return err
	//}

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
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
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

	return resourceRelationalDatabaseRead(d, m)
}

func resourceRelationalResizeFlavor(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database update")

	cli := m.(*client.Client)

	updateRequest := generateResizeFlavorDatabaseRequest(d)
	reqBody, _ := json.Marshal(updateRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.ResizeInstance1(context.TODO(), string(reqBody), d.Id(), nil)

	//if err != nil {
	//	return err
	//}

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
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
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

func resourceRelationalDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Database delete")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "delete")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.DeleteDatabaseInstances1(context.TODO(), string(reqBody), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when deleting database: %s", err)
	//}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseDeletePending,
		Target:     databaseDeleteTarget,
		Refresh:    resourceRelationalDatabaseDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseDeleteTimeout,
		Delay:      databaseDeleteDelay,
		MinTimeout: databaseDeleteMinTimeout,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be deleted: %s", err)
	}

	log.Println("[DEBUG]  Delete done")

	d.SetId("")
	return nil
}

func resourceRelationalDatabaseDeleteStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById1(context.TODO(), id)
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

func resourceRelationalDatabaseStart(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database start")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "start")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.StartDatabaseInstances1(context.TODO(), string(reqBody), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when starting database: %s", err)
	//}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStartPending,
		Target:     databaseStartTarget,
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStartTimeout,
		Delay:      databaseStartDelay,
		MinTimeout: databaseStartMinTimeout,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be started: %s", err)
	}

	return resourceRelationalDatabaseRead(d, m)
}

func resourceRelationalDatabaseStop(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database stop")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "stop")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.StopDatabaseInstances1(context.TODO(), string(reqBody), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when stopping database: %s", err)
	//}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStopPending,
		Target:     databaseStopTarget,
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStopTimeout,
		Delay:      databaseStopDelay,
		MinTimeout: databaseStopMinTimeout,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be stopped: %s", err)
	}

	return resourceRelationalDatabaseRead(d, m)
}

func resourceRelationalDatabaseReboot(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database reboot")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "reboot")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.RestartDatabaseInstances1(context.TODO(), string(reqBody), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when rebooting database: %s", err)
	//}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseRebootPending,
		Target:     databaseRebootTarget,
		Refresh:    resourceRelationalDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseRebootTimeout,
		Delay:      databaseRebootDelay,
		MinTimeout: databaseRebootMinTimeout,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error when waiting for database to be rebooted: %s", err)
	}

	return resourceRelationalDatabaseRead(d, m)
}

func resourceRelationalDatabasePromote(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Database promote")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "detach_replica")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.DetachReplica1(context.TODO(), string(reqBody), d.Id())
	//if err != nil {
	//	return fmt.Errorf("error when promoting database: %s", err)
	//}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	err := resource.Retry(databasePromoteTimeout, func() *resource.RetryError {
		dbResp, httpResponse, err := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById1(context.TODO(), d.Id())
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

	return resourceRelationalDatabaseRead(d, m)
}

func generateResizeVolumeDatabaseRequest(d *schema.ResourceData) ResizeRequest {
	config := ResizeConfig{
		IsPoc:      d.Get("is_poc").(bool),
		VolumeSize: d.Get("volume_size").(int),
		VolumeType: d.Get("volume_type").(string),
	}

	instance := ResizeInstance{
		InstancesID: d.Id(),
		Config:      config,
	}
	instances := make([]ResizeInstance, 1)
	instances[0] = instance
	resizeVolumeRequest := ResizeRequest{
		ResourceType:      "dbaas",
		Action:            "resize",
		DatabaseInstances: instances,
	}

	return resizeVolumeRequest
}

func generateResizeFlavorDatabaseRequest(d *schema.ResourceData) ResizeRequest {
	config := ResizeConfig{
		IsPoc:     d.Get("is_poc").(bool),
		PackageID: d.Get("package_id").(string),
	}

	instance := ResizeInstance{
		InstancesID: d.Id(),
		Config:      config,
	}
	instances := make([]ResizeInstance, 1)
	instances[0] = instance
	resizeFlavorRequest := ResizeRequest{
		ResourceType:      "dbaas",
		Action:            "resize",
		DatabaseInstances: instances,
	}

	return resizeFlavorRequest
}

func generateUpdateConfigGroupRequest(d *schema.ResourceData) UpdateDBRequest {
	request := UpdateDBRequest{
		DbInstanceID: d.Id(),
		ConfigID:     d.Get("config_id").(string),
	}

	return request
}

func generateUpdateSettingRequest(d *schema.ResourceData) UpdateDBRequest {
	publicAccess := d.Get("public_access").(bool)
	backupAuto := d.Get("backup_auto").(bool)
	request := UpdateDBRequest{
		DbInstanceID:   d.Id(),
		Password:       "",
		PublicAccess:   &publicAccess,
		BackupAuto:     &backupAuto,
		BackupDuration: d.Get("backup_duration").(int),
		BackupTime:     d.Get("backup_time").(string),

		RedisPassword: "",
	}

	return request
}

func generateActionRequest(d *schema.ResourceData, action string) ActionRequest {
	instance := Instance{
		InstancesID: d.Id(),
	}
	instances := make([]Instance, 1)
	instances[0] = instance
	actionRequest := ActionRequest{
		ResourceType:      "dbaas",
		Action:            action,
		DatabaseInstances: instances,
	}

	return actionRequest
}

type ResizeConfig struct {
	VolumeType string `json:"volumeType,omitempty"`
	VolumeSize int    `json:"volumeSize,omitempty"`
	PackageID  string `json:"packageId,omitempty"`
	IsPoc      bool   `json:"poc,omitempty"`
}

type ResizeInstance struct {
	InstancesID string       `json:"instancesId,omitempty"`
	Config      ResizeConfig `json:"config,omitempty"`
}

type ResizeRequest struct {
	ResourceType      string           `json:"resourceType,omitempty"`
	Action            string           `json:"action,omitempty"`
	DatabaseInstances []ResizeInstance `json:"databaseInstances,omitempty"`
}

type UpdateDBRequest struct {
	DbInstanceID string `json:"dbInstanceId,omitempty"`

	ConfigID string `json:"configId,omitempty"`

	Password       string `json:"password,omitempty"`
	PublicAccess   *bool  `json:"publicAccess,omitempty"`
	BackupAuto     *bool  `json:"backupAuto,omitempty"`
	BackupDuration int    `json:"backupDuration,omitempty"`
	BackupTime     string `json:"backupTime,omitempty"`

	RedisPassword        string `json:"redisPassword,omitempty"`
	RedisPasswordEnabled *bool  `json:"redisPasswordEnabled,omitempty"`
	EditRedisPassword    *bool  `json:"editRedisPassword,omitempty"`
}

type Instance struct {
	InstancesID string `json:"instancesId,omitempty"`
}

type ActionRequest struct {
	ResourceType      string     `json:"resType,omitempty"`
	Action            string     `json:"action,omitempty"`
	DatabaseInstances []Instance `json:"databaseInstances,omitempty"`
}
