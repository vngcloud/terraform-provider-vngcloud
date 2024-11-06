package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdb"
)

func ResourceMemStoreBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemStoreBackupCreate,
		Read:   resourceMemStoreBackupRead,
		Delete: resourceMemStoreBackupDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"net_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"net_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"backup_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMemStoreBackupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup read")

	cli := m.(*client.Client)

	dbResp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreBackupAPIApi.GetDetailBackupById(context.TODO(), d.Id())
	//if err != nil {
	//	return err
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

	d.Set("name", dbResp.Data.Name)
	d.Set("instance_id", dbResp.Data.DbInstanceId)
	d.Set("instance_name", dbResp.Data.InstanceName)
	d.Set("type", dbResp.Data.Type_)
	d.Set("parent_id", dbResp.Data.Parent)
	d.Set("parent_name", dbResp.Data.ParentName)
	d.Set("created", dbResp.Data.Created)
	d.Set("backup_type", dbResp.Data.BackupType)
	d.Set("status", dbResp.Data.Status)
	d.Set("engine_type", dbResp.Data.DatastoreType)
	d.Set("engine_version", dbResp.Data.DatastoreVersion)
	d.Set("storage_type", dbResp.Data.StorageType)
	d.Set("storage_size", dbResp.Data.StorageSize)
	d.Set("ram", dbResp.Data.Ram)
	d.Set("cpu", dbResp.Data.Vcpu)
	d.Set("config_id", dbResp.Data.ConfigId)
	d.Set("config_name", dbResp.Data.ConfigName)
	d.Set("username", dbResp.Data.Username)
	d.Set("net_ids", dbResp.Data.NetIds)
	d.Set("net_name", dbResp.Data.NetName)
	d.Set("size", dbResp.Data.Size)
	d.Set("backup_tier", dbResp.Data.BackupTier)

	return nil
}

func resourceMemStoreBackupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup create")

	cli := m.(*client.Client)

	createRequest := generateMemStoreCreateBackupRequest(d)
	reqBody, _ := json.Marshal(createRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	createDbResult, httpResponse, _ := cli.Vdbv2Client.MemoryStoreBackupAPIApi.CreateBackups(context.TODO(), string(reqBody))
	//if err != nil {
	//	return err
	//}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	log.Println("[DEBUG] Create backup result, id: " + createDbResult.Data.BackupId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"NEW", "BUILDING", "SAVING"},
		Target:     []string{"COMPLETED"},
		Refresh:    resourceMemStoreBackupStateRefreshFunc(cli, createDbResult.Data.BackupId),
		Timeout:    30 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	id, _ := stateConf.WaitForStateContext(context.TODO())
	idStr := id.(string)

	log.Println("[DEBUG] Wait for state done, id: " + idStr)

	d.SetId(idStr)

	return resourceMemStoreBackupRead(d, m)
}

func generateMemStoreCreateBackupRequest(d *schema.ResourceData) vdb.BackupRequest {
	createRequest := vdb.BackupRequest{
		DbInstanceId: d.Get("instance_id").(string),
		Name:         d.Get("name").(string),
		Description:  "Created by Terraform",
		BackupType:   "FULL",
	}

	return createRequest
}

func resourceMemStoreBackupStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG] State refresh")

		dbResp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreBackupAPIApi.GetDetailBackupById(context.TODO(), id)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return nil, "", fmt.Errorf("error when refreshing backup state: %s", responseBody)
		}
		if dbResp.Data == nil {
			return nil, "", fmt.Errorf("error when refreshing backup state: data is empty")
		}
		log.Println("[DEBUG] Backup status: " + dbResp.Data.Status)

		return dbResp.Data.Id, dbResp.Data.Status, nil
	}
}

func resourceMemStoreBackupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup delete")

	cli := m.(*client.Client)
	backupId := d.Id()

	deleteRequest := BackupDeleteRequest{
		{backupId},
	}
	reqBody, _ := json.Marshal(deleteRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreBackupAPIApi.DeleteBackups(context.TODO(), string(reqBody))

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
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    resourceMemStoreBackupDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for delete backup (%s) : %s", d.Id(), err)
	}
	d.SetId("")

	return nil
}

func generateMemStoreRestoreBackupRequest(d *schema.ResourceData) RestoreBackupRequest {
	config := RestoreBackupConfig{
		BackupAuto:           d.Get("backup_auto").(bool),
		BackupDuration:       d.Get("backup_duration").(int),
		BackupID:             d.Get("backup_id").(string),
		BackupTime:           d.Get("backup_time").(string),
		ConfigID:             d.Get("config_id").(string),
		DatastoreType:        d.Get("engine_type").(string),
		DatastoreVersion:     d.Get("engine_version").(string),
		Name:                 d.Get("name").(string),
		NetIds:               []string{d.Get("subnet_id").(string)},
		PackageID:            d.Get("package_id").(string),
		IsPoc:                d.Get("is_poc").(bool),
		PublicAccess:         d.Get("public_access").(bool),
		RedisPassword:        d.Get("redis_password").(string),
		RedisPasswordEnabled: d.Get("redis_password_enabled").(bool),
	}

	instance := RestoreBackupInstance{
		InstancesID: d.Id(),
		Config:      config,
	}
	instances := make([]RestoreBackupInstance, 1)
	instances[0] = instance
	restoreBackupRequest := RestoreBackupRequest{
		ResourceType:      "dbaas-backup",
		Action:            "restore_backup",
		DatabaseInstances: instances,
	}

	return restoreBackupRequest
}

func resourceMemStoreBackupDeleteStateRefreshFunc(cli *client.Client, backupId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.MemoryStoreBackupAPIApi.GetDetailBackupById(context.TODO(), backupId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vdbv2.BackupInfo{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing backup: %s", GetResponseBody(httpResponse))
			}
		} else {
			if dbResp.Data == nil {
				return vdbv2.BackupInfo{Status: "DELETED"}, "DELETED", nil
			}
		}
		respJSON, _ := json.Marshal(dbResp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		return vdbv2.BackupInfo{Status: "ACTIVE"}, "ACTIVE", nil
	}
}
