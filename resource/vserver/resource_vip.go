package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
	"net/http"
	"strings"
)

const (
	errorVipRead     = "error reading Virtual Ip Address (%s): %s"
	errorVipCreate   = "error creating Virtual Ip Address : %s"
	errorImportVip   = "import format error: to import a Virtual Ip Address use the format {projectId}:{vipId}"
	errorVipDelete   = "error deleting Virtual Ip Address (%s): %s"
	errorVipNotFound = "[WARN] Virtual Ip Address (%s) not found, removing from state"
)

func ResourceVip() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceVipRead,
		CreateContext: resourceVipCreate,
		UpdateContext: resourceVipUpdate,
		DeleteContext: resourceVipDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceVipImport,
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"vip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVipImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 2 || parts[0] == "" {
		return nil, fmt.Errorf(errorImportVip)
	}

	projectId := parts[0]
	vipId := parts[1]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	d.SetId(vipId)
	return []*schema.ResourceData{d}, nil
}

func resourceVipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read Virtual Ip Address")

	projectId := d.Get("project_id").(string)
	vipId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.GetVirtualIpAddressUsingGET(ctx, projectId, vipId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorVipNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf(errorVipRead, vipId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.Set("name", resp.Data.Name)
	d.Set("description", resp.Data.Description)
	d.Set("subnet_id", resp.Data.SubnetId)
	d.Set("vip_address", resp.Data.IpAddress)

	log.Printf("Read Virtual Ip Address successfully")
	return nil
}

func resourceVipCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Creating Virtual Ip Address")

	var err error
	projectId := d.Get("project_id").(string)

	cli := m.(*client.Client)
	req := vserver.CreateVirtualIpAddressRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		SubnetId:    d.Get("subnet_id").(string),
	}

	dataResp, httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.CreateVirtualIpAddressUsingPOST(ctx, req, projectId)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorVipCreate, responseBody)
		return diag.FromErr(errResponse)
	}

	d.SetId(dataResp.Data.Uuid)

	log.Printf("Creating Virtual Ip Address successfully")
	return resourceVipRead(ctx, d, m)
}

func resourceVipDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting Virtual Ip Address")

	projectId := d.Get("project_id").(string)
	vipId := d.Id()

	cli := m.(*client.Client)
	httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.DeleteVirtualIpAddressUsingDELETE(ctx, projectId, vipId)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf(errorVipDelete, vipId, responseBody)
		return diag.FromErr(errorResponse)
	}

	d.SetId("")

	log.Printf("Deleted Virtual Ip Address successfully")
	return nil
}

func resourceVipUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Updating virtual ip address")
	projectId := d.Get("project_id").(string)
	vipId := d.Id()

	if d.HasChange("name") || d.HasChange("description") {
		updateVipReq := vserver.UpdateVirtualIpAddressRequest{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		}
		cli := m.(*client.Client)
		_, httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.UpdateUsingPUT3(ctx, projectId, updateVipReq, vipId)

		if err != nil {
			log.Printf("[ERROR] %s\n", err)
		}
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return diag.FromErr(errResponse)
		}

	}

	log.Printf("Updated virtual ip address successfully")
	return resourceVipRead(ctx, d, m)
}
