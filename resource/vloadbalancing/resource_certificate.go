package vloadbalancing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vloadbalancing"
	"log"
	"net/http"
	"time"
)

const (
	errorCARead      = "error reading CA (%s): %s"
	errorCAImport    = "error importing CA: %s"
	errorCADelete    = "error deleting CA (%s): %s"
	errorParsingDate = "error parsing date (%s): %s"
	errorCANotFound  = "[WARN] CA (%s) not found, removing from state"
)

func ResourceCA() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceCARead,
		CreateContext: resourceCAImporting,
		DeleteContext: resourceCADelete,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"certificate_chain": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"expiration": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCARead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Read Certificate Authority")

	projectId := d.Get("project_id").(string)
	caId := d.Id()

	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VlbClient.CaRestControllerV2Api.GetCertificateAuthorityByIdUsingGET(ctx, caId, projectId)

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		log.Printf(errorCANotFound, d.Id())
		d.SetId("")
		return nil
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		err := fmt.Errorf(errorCARead, caId, responseBody)
		return diag.FromErr(err)
	}

	log.Printf("Read Certificate Authority successfully")
	return nil
}

func resourceCAImporting(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Importing Certificate Authority")

	var expirationTime time.Time
	var err error
	projectId := d.Get("project_id").(string)

	if expiration, ok := d.Get("expiration").(string); ok {
		expirationTime, err = time.Parse(time.RFC3339, expiration)
		if err != nil {
			log.Printf(errorParsingDate, expiration, err)
		}
	}

	cli := m.(*client.Client)
	req := vloadbalancing.ImportCaRequest{
		Name:             d.Get("name").(string),
		Certificate:      d.Get("certificate").(string),
		CertificateChain: d.Get("certificate_chain").(string),
		Expiration:       expirationTime,
		Passphrase:       d.Get("passphrase").(string),
		PrivateKey:       d.Get("private_key").(string),
		Type_:            d.Get("type").(string),
	}

	ca, httpResponse, err := cli.VlbClient.CaRestControllerV2Api.ImportCAUsingPOST(ctx, req, projectId)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		errResponse := fmt.Errorf(errorCAImport, responseBody)
		return diag.FromErr(errResponse)
	}

	d.SetId(ca.Data.Uuid)

	log.Printf("Imported Certificate Authority successfully")
	return nil
}

func resourceCADelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("Deleting Certificate Authority")

	projectId := d.Get("project_id").(string)
	caId := d.Id()

	cli := m.(*client.Client)
	httpResponse, err := cli.VlbClient.CaRestControllerV2Api.DeleteCertificateAuthorityUsingDELETE(ctx, caId, projectId)

	if err != nil {
		log.Printf("[ERROR] %s\n", err)
	}

	if checkErrorResponse(httpResponse) {
		responseBody := getResponseBody(httpResponse)
		errorResponse := fmt.Errorf(errorCADelete, caId, responseBody)
		return diag.FromErr(errorResponse)
	}

	d.SetId("")

	log.Printf("Deleted Certificate Authority successfully")
	return nil
}
