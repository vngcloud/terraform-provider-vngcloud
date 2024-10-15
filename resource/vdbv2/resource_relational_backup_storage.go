package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"log"
	"time"
)

func ResourceRelationalBackupStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceRelationalBackupStorageCreate,
		Read:   resourceRelationalBackupStorageRead,
		Update: resourceRelationalBackupStorageUpdate,
		Delete: resourceRelationalBackupStorageDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"backup_storage_package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"quota": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRelationalBackupStorageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage read")

	cli := m.(*client.Client)

	backupStorageResult, httpResponse, err := cli.Vdbv2Client.RelationalBackupStorageAPIApi.GetListBackupStorage1(context.TODO(), nil)
	if err != nil {
		return err
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	if backupStorageResult.Data == nil || len(backupStorageResult.Data) == 0 {
		d.SetId("")
		return nil
	}

	backupStorageInfo := backupStorageResult.Data[0]

	d.Set("backup_storage_package_id", backupStorageInfo.BackupPackageId)
	d.Set("name", backupStorageInfo.Name)
	d.Set("quota", backupStorageInfo.Quota)
	d.Set("usage", backupStorageInfo.Usage)
	d.SetId(backupStorageInfo.Id)

	return nil
}

func resourceRelationalBackupStorageCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage create")

	cli := m.(*client.Client)

	createBackupStorageRequest := generateRelationalCreateBackupStorageRequest(d)
	reqBody, _ := json.Marshal(createBackupStorageRequest)
	log.Println("[DEBUG] Request body " + string(reqBody))

	createBackupStorageResult, httpResponse, err := cli.Vdbv2Client.RelationalBackupStorageAPIApi.CreateRelationalBackUpStorage(context.TODO(), string(reqBody), nil)
	if err != nil {
		return err
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	idStr := createBackupStorageResult.Data[0].ResourceId

	time.Sleep(10 * time.Second)

	d.SetId(idStr)
	return resourceRelationalBackupStorageRead(d, m)
}

func generateRelationalCreateBackupStorageRequest(d *schema.ResourceData) CreateBackupStorageRequest {
	createRequest := CreateBackupStorageRequest{
		BackupPackageID: d.Get("backup_storage_package_id").(string),
		IsPoc:           d.Get("is_poc").(bool),
		Period:          1,
	}
	return createRequest
}

func resourceRelationalBackupStorageUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRelationalBackupStorageResizeQuota(d, m)
}

func resourceRelationalBackupStorageResizeQuota(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage resize quota")

	cli := m.(*client.Client)

	instance := BackupStorageInstance{
		InstancesID: d.Id(),
		Config: BackupStorageConfig{
			BackupPackageID: d.Get("backup_storage_package_id").(string),
			IsPoc:           d.Get("is_poc").(bool),
		},
	}
	instances := make([]BackupStorageInstance, 1)
	instances[0] = instance
	resizeQuotaRequest := BackupStorageRequest{
		ResourceType:      "dbaas-backup-storage",
		Action:            "resize",
		DatabaseInstances: instances,
	}
	reqBody, _ := json.Marshal(resizeQuotaRequest)
	log.Println("[DEBUG] Resize quota request body " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.RelationalBackupStorageAPIApi.ResizeBackupStorage1(context.TODO(), string(reqBody), nil)
	if err != nil {
		return err
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	time.Sleep(10 * time.Second)

	return nil
}

func resourceRelationalBackupStorageDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage delete")

	cli := m.(*client.Client)

	instance := BackupStorageInstance{
		InstancesID: d.Id(),
	}
	instances := make([]BackupStorageInstance, 1)
	instances[0] = instance
	deleteBackupStorageRequest := BackupStorageNoPaymentRequest{
		ResourceType:      "dbaas-backup-storage",
		Action:            "delete",
		DatabaseInstances: instances,
	}
	reqBody, _ := json.Marshal(deleteBackupStorageRequest)
	log.Println("[DEBUG] Delete request body " + string(reqBody))

	_, httpResponse, err := cli.Vdbv2Client.RelationalBackupStorageAPIApi.DeleteBackupStorage1(context.TODO(), string(reqBody))
	if err != nil {
		return err
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}

	time.Sleep(10 * time.Second)

	d.SetId("")
	return nil
}

type BackupStorageConfig struct {
	BackupPackageID string `json:"backupPackageId,omitempty"`
	IsPoc           bool   `json:"poc,omitempty"`
	EngineGroup     int    `json:"engineGroup,omitempty"`
}

type BackupStorageInstance struct {
	InstancesID string              `json:"instancesId,omitempty"`
	Config      BackupStorageConfig `json:"config,omitempty"`
}

type BackupStorageRequest struct {
	ResourceType      string                  `json:"resourceType,omitempty"`
	Action            string                  `json:"action,omitempty"`
	DatabaseInstances []BackupStorageInstance `json:"databaseInstances,omitempty"`
}

type BackupStorageNoPaymentRequest struct {
	ResourceType      string                  `json:"resType,omitempty"`
	Action            string                  `json:"action,omitempty"`
	DatabaseInstances []BackupStorageInstance `json:"databaseInstances,omitempty"`
}

type CreateBackupStorageRequest struct {
	BackupPackageID string `json:"backupPackageId,omitempty"`
	IsPoc           bool   `json:"poc,omitempty"`
	Period          int    `json:"period,omitempty"`
}
