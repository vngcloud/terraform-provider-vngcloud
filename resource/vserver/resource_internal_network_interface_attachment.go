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
	errorAddressPairInterfaceAttachmentRead     = "error reading Address Pair Interface Attachment (%s): %s"
	errorAddressPairInterfaceAttachmentCreate   = "error creating Address Pair Interface Attachment : %s"
	errorAddressPairInterfaceAttachmentDelete   = "error deleting Address Pair Interface Attachment (%s): %s"
	errorAddressPairInterfaceAttachmentNotFound = "[WARN] Address Pair Interface Attachment (%s) not found, removing from state"
)

func ResourceAddressPairInterfaceAttachment() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAddressPairInterfaceAttachmentRead,
		CreateContext: resourceAddressPairInterfaceAttachmentCreate,
		DeleteContext: resourceAddressPairInterfaceAttachmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceAddressPairInterfaceAttachmentImport,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_ip_address_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"internal_network_interface_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAddressPairInterfaceAttachmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ":")

	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return nil, fmt.Errorf(errorImportVip)
	}

	projectId := parts[0]
	vipId := parts[1]
	addressPairId := parts[2]

	if err := d.Set("project_id", projectId); err != nil {
		return nil, fmt.Errorf("error setting `project_id`: %s", err)
	}

	if err := d.Set("virtual_ip_address_id", vipId); err != nil {
		return nil, fmt.Errorf("error setting `virtual_ip_address_id`: %s", err)
	}

	d.SetId(addressPairId)
	return []*schema.ResourceData{d}, nil
}

func resourceAddressPairInterfaceAttachmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read Address Pair Interface Attachment")

	projectId := d.Get("project_id").(string)
	virtualIpAddressId := d.Get("virtual_ip_address_id").(string)
	addressPairId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.GetSpecificAddressPairUsingGET(ctx, projectId, virtualIpAddressId, addressPairId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorAddressPairInterfaceAttachmentNotFound, d.Id())
		d.SetId("")
		return nil
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf(errorAddressPairInterfaceAttachmentRead, addressPairId, responseBody)
		return diag.FromErr(err)
	}

	// set to state
	d.SetId(resp.Data.Uuid)
	d.Set("internal_network_interface_id", resp.Data.NetworkInterfaceId)
	d.Set("virtual_ip_address_id", resp.Data.VirtualIpAddressId)

	log.Printf("Read Address Pair Interface Attachment successfully")
	return nil
}

func resourceAddressPairInterfaceAttachmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Creating Address Pair Interface Attachment")

	var err error
	projectId := d.Get("project_id").(string)
	virtualIpAddressId := d.Get("virtual_ip_address_id").(string)

	cli := m.(*client.Client)
	req := vserver.CreateAddressPairRequest{
		InternalNetworkInterfaceId: d.Get("internal_network_interface_id").(string),
	}

	dataResp, httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.AddAddressPairUsingPOST(ctx, req, projectId, virtualIpAddressId)
	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorAddressPairInterfaceAttachmentCreate, responseBody)
		return diag.FromErr(errResponse)
	}

	d.SetId(dataResp.Data.Uuid)

	log.Printf("Creating Address Pair Interface Attachment successfully")
	return resourceAddressPairInterfaceAttachmentRead(ctx, d, m)
}

func resourceAddressPairInterfaceAttachmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting Address Pair Interface Attachment")

	projectId := d.Get("project_id").(string)

	cli := m.(*client.Client)
	httpResponse, err := cli.VserverClient.VirtualIpAddressRestControllerV2Api.RemoveAddressPairUsingDELETE(ctx, d.Id(), projectId)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf(errorAddressPairInterfaceAttachmentDelete, d.Id(), responseBody)
		return diag.FromErr(errorResponse)
	}

	d.SetId("")

	log.Printf("Deleted Address Pair Interface Attachment successfully")
	return nil
}
