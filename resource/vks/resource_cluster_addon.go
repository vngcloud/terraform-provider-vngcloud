package vks

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
)

// Add-on name / version / namespace rules, mirroring vks-api AddonConst.
var (
	addonNameRegexp      = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`)
	addonVersionRegexp   = regexp.MustCompile(`^v?[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.]+)?-vks\.[0-9]+$`)
	addonNamespaceRegexp = regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
)

const addonMaxConfigurationValuesLength = 65536

// ResourceClusterAddon manages one add-on of a VKS cluster (like aws_eks_addon).
func ResourceClusterAddon() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterAddonCreate,
		Read:   resourceClusterAddonRead,
		Update: resourceClusterAddonUpdate,
		Delete: resourceClusterAddonDelete,
		Importer: &schema.ResourceImporter{
			State: resourceClusterAddonImport,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the cluster.",
			},
			"addon_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(addonNameRegexp, "must be 2-63 lowercase alphanumeric characters or hyphens"),
				Description:  "Name of the add-on, e.g. block-storage-csi. See the vngcloud_vks_addon_version data source.",
			},
			"addon_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringMatch(addonVersionRegexp, "must be a VKS add-on build such as 1.2.0-vks.1"),
				Description:  "Version of the add-on. Defaults to the default version for the cluster's Kubernetes minor.",
			},
			"configuration_values": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, addonMaxConfigurationValuesLength),
				Description:  "Custom configuration values as a JSON or YAML string; must match the configuration schema of the version. Removing it clears the values.",
			},
			"namespace": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.All(validation.StringLenBetween(1, 63), validation.StringMatch(addonNamespaceRegexp, "must be a valid Kubernetes namespace name")),
				Description:  "Namespace to install the add-on into. Defaults to the version's default namespace.",
			},
			"resolve_conflicts_on_create": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OVERWRITE", "PRESERVE"}, false),
				Description:  "How to resolve conflicts with existing in-cluster resources on create: NONE, OVERWRITE or PRESERVE.",
			},
			"resolve_conflicts_on_update": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "OVERWRITE", "PRESERVE"}, false),
				Description:  "How to resolve fields changed outside VKS on update: NONE, OVERWRITE or PRESERVE.",
			},
			"preserve": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "On destroy, stop managing the add-on but keep its Helm release in the cluster.",
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"publisher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_issues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code":         {Type: schema.TypeString, Computed: true},
						"message":      {Type: schema.TypeString, Computed: true},
						"resource_ids": {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
					},
				},
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// clusterAddonID / parseClusterAddonID: the resource ID is "<cluster_id>:<addon_name>".
func clusterAddonID(clusterID, addonName string) string {
	return clusterID + ":" + addonName
}

func parseClusterAddonID(id string) (string, string, error) {
	parts := strings.Split(id, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("unexpected ID %q, expected <cluster_id>:<addon_name>", id)
	}
	return parts[0], parts[1], nil
}

// addonAPIError formats a failed call. The generated client closes the
// response body and saves its contents in GenericSwaggerError.
func addonAPIError(action string, httpResponse *http.Response, err error) error {
	if httpResponse == nil || !CheckErrorResponse(httpResponse) {
		return fmt.Errorf("%s: %w", action, err)
	}
	var apiErr vks.GenericSwaggerError
	if httpResponse.StatusCode != http.StatusForbidden && errors.As(err, &apiErr) {
		return fmt.Errorf("%s: Status Code: %d, %s", action, httpResponse.StatusCode, apiErr.Body())
	}
	return fmt.Errorf("%s: %s", action, GetResponseBody(httpResponse))
}

func formatAddonIssues(addon *vks.AddonDto) string {
	if addon == nil || addon.Health == nil || len(addon.Health.Issues) == 0 {
		return ""
	}
	var parts []string
	for _, i := range addon.Health.Issues {
		parts = append(parts, fmt.Sprintf("%s: %s", i.Code, i.Message))
	}
	return " (" + strings.Join(parts, "; ") + ")"
}

func resourceClusterAddonCreate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)
	addonName := d.Get("addon_name").(string)

	request := vks.CreateAddonRequestDto{
		AddonName:           addonName,
		AddonVersion:        d.Get("addon_version").(string),
		ConfigurationValues: d.Get("configuration_values").(string),
		ResolveConflicts:    d.Get("resolve_conflicts_on_create").(string),
		ClientRequestToken:  resource.PrefixedUniqueId("tf-"),
	}
	if ns, ok := d.GetOk("namespace"); ok {
		request.NamespaceConfig = &vks.NamespaceConfigDto{Namespace: ns.(string)}
	}
	log.Printf("[DEBUG] Creating add-on %s on cluster %s", addonName, clusterID)

	_, httpResponse, err := cli.VksClient.V1ClusterAddonControllerApi.V1ClustersClusterIdAddonsPost(context.TODO(), request, clusterID)
	if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
		return addonAPIError(fmt.Sprintf("creating add-on %s", addonName), httpResponse, err)
	}
	d.SetId(clusterAddonID(clusterID, addonName))

	if err := waitClusterAddonStatus(cli, clusterID, addonName, addonCreatePending, addonCreateTarget, d.Timeout(schema.TimeoutCreate)); err != nil {
		return fmt.Errorf("error waiting for add-on %s to be created: %s", d.Id(), err)
	}
	return resourceClusterAddonRead(d, m)
}

func resourceClusterAddonRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	clusterID, addonName, err := parseClusterAddonID(d.Id())
	if err != nil {
		return err
	}
	resp, httpResponse, err := cli.VksClient.V1ClusterAddonControllerApi.V1ClustersClusterIdAddonsAddonNameGet(context.TODO(), clusterID, addonName)
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		log.Printf("[WARN] Add-on %s not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
		return addonAPIError(fmt.Sprintf("reading add-on %s", d.Id()), httpResponse, err)
	}
	addon := resp.Addon
	if addon == nil {
		return fmt.Errorf("reading add-on %s: empty response", d.Id())
	}

	d.Set("cluster_id", clusterID)
	d.Set("addon_name", addonName)
	d.Set("addon_version", addon.AddonVersion)
	d.Set("configuration_values", addon.ConfigurationValues)
	if addon.NamespaceConfig != nil {
		d.Set("namespace", addon.NamespaceConfig.Namespace)
	}
	d.Set("status", addon.Status)
	d.Set("owner", addon.Owner)
	d.Set("publisher", addon.Publisher)
	var issues []map[string]interface{}
	if addon.Health != nil {
		for _, i := range addon.Health.Issues {
			issues = append(issues, map[string]interface{}{
				"code":         i.Code,
				"message":      i.Message,
				"resource_ids": i.ResourceIds,
			})
		}
	}
	d.Set("health_issues", issues)
	if !addon.CreatedAt.IsZero() {
		d.Set("created_at", addon.CreatedAt.Format(time.RFC3339))
	}
	if !addon.ModifiedAt.IsZero() {
		d.Set("modified_at", addon.ModifiedAt.Format(time.RFC3339))
	}
	return nil
}

func resourceClusterAddonUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	clusterID, addonName, err := parseClusterAddonID(d.Id())
	if err != nil {
		return err
	}

	// resolve_conflicts_* and preserve only affect later calls: nothing to send.
	if !d.HasChange("addon_version") && !d.HasChange("configuration_values") {
		return resourceClusterAddonRead(d, m)
	}

	request := vks.UpdateAddonRequestDto{
		ResolveConflicts:   d.Get("resolve_conflicts_on_update").(string),
		ClientRequestToken: resource.PrefixedUniqueId("tf-"),
	}
	if d.HasChange("addon_version") {
		request.AddonVersion = d.Get("addon_version").(string)
	}
	if d.HasChange("configuration_values") {
		// "" (attribute removed) clears the values.
		v := d.Get("configuration_values").(string)
		request.ConfigurationValues = &v
	}
	if request.AddonVersion == "" && request.ConfigurationValues == nil {
		return resourceClusterAddonRead(d, m)
	}
	log.Printf("[DEBUG] Updating add-on %s", d.Id())

	resp, httpResponse, err := cli.VksClient.V1ClusterAddonControllerApi.V1ClustersClusterIdAddonsAddonNameUpdatePost(context.TODO(), request, clusterID, addonName)
	if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
		return addonAPIError(fmt.Sprintf("updating add-on %s", d.Id()), httpResponse, err)
	}
	if resp.Update == nil || resp.Update.Id == "" {
		return fmt.Errorf("updating add-on %s: response has no update", d.Id())
	}

	if err := waitClusterAddonUpdate(cli, clusterID, addonName, resp.Update.Id, d.Timeout(schema.TimeoutUpdate)); err != nil {
		return fmt.Errorf("error waiting for update %s of add-on %s: %s", resp.Update.Id, d.Id(), err)
	}
	return resourceClusterAddonRead(d, m)
}

func resourceClusterAddonDelete(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	clusterID, addonName, err := parseClusterAddonID(d.Id())
	if err != nil {
		return err
	}
	opts := &vks.V1ClusterAddonControllerApiV1ClustersClusterIdAddonsAddonNameDeleteOpts{
		Preserve: optional.NewBool(d.Get("preserve").(bool)),
	}
	_, httpResponse, err := cli.VksClient.V1ClusterAddonControllerApi.V1ClustersClusterIdAddonsAddonNameDelete(context.TODO(), clusterID, addonName, opts)
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
		return addonAPIError(fmt.Sprintf("deleting add-on %s", d.Id()), httpResponse, err)
	}
	if err := waitClusterAddonStatus(cli, clusterID, addonName, addonDeletePending, DELETED, d.Timeout(schema.TimeoutDelete)); err != nil {
		return fmt.Errorf("error waiting for add-on %s to be deleted: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func resourceClusterAddonImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	clusterID, addonName, err := parseClusterAddonID(d.Id())
	if err != nil {
		return nil, err
	}
	d.Set("cluster_id", clusterID)
	d.Set("addon_name", addonName)
	if err := resourceClusterAddonRead(d, m); err != nil {
		return nil, err
	}
	if d.Id() == "" {
		return nil, fmt.Errorf("add-on %s:%s not found", clusterID, addonName)
	}
	return []*schema.ResourceData{d}, nil
}

// clusterAddonStateRefreshFunc reports the add-on status; a 404 is DELETED.
// A *_FAILED status is returned as an error so the wait stops at once.
func clusterAddonStateRefreshFunc(cli *client.Client, clusterID, addonName string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, err := cli.VksClient.V1ClusterAddonControllerApi.V1ClustersClusterIdAddonsAddonNameGet(context.TODO(), clusterID, addonName)
		if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
			return vks.AddonDto{Status: "DELETED"}, "DELETED", nil
		}
		if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
			return nil, "", addonAPIError("describing add-on", httpResponse, err)
		}
		if resp.Addon == nil {
			return nil, "", fmt.Errorf("describing add-on %s: empty response", addonName)
		}
		status := resp.Addon.Status
		log.Printf("[DEBUG] Add-on %s:%s status %s", clusterID, addonName, status)
		if CheckContainString(addonFailed, status) {
			return resp.Addon, status, fmt.Errorf("add-on reached %s%s", status, formatAddonIssues(resp.Addon))
		}
		return resp.Addon, status, nil
	}
}

func waitClusterAddonStatus(cli *client.Client, clusterID, addonName string, pending, target []string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:    pending,
		Target:     target,
		Refresh:    clusterAddonStateRefreshFunc(cli, clusterID, addonName),
		Timeout:    timeout,
		Delay:      addonDelay,
		MinTimeout: addonMinTimeout,
	}
	_, err := stateConf.WaitForState()
	return err
}

// waitClusterAddonUpdate polls an add-on update until Successful; Failed or
// Cancelled stop the wait with the update errors.
func waitClusterAddonUpdate(cli *client.Client, clusterID, addonName, updateID string, timeout time.Duration) error {
	opts := &vks.V1ClusterUpdateControllerApiV1ClustersClusterIdUpdatesUpdateIdGetOpts{
		AddonName: optional.NewString(addonName),
	}
	stateConf := &resource.StateChangeConf{
		Pending: addonUpdateInProgress,
		Target:  addonUpdateSuccessful,
		Refresh: func() (interface{}, string, error) {
			resp, httpResponse, err := cli.VksClient.V1ClusterUpdateControllerApi.V1ClustersClusterIdUpdatesUpdateIdGet(context.TODO(), clusterID, updateID, opts)
			if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
				return nil, "", addonAPIError("describing update", httpResponse, err)
			}
			if resp.Update == nil {
				return nil, "", fmt.Errorf("describing update %s: empty response", updateID)
			}
			u := resp.Update
			if u.Status == "Failed" || u.Status == "Cancelled" {
				var msgs []string
				for _, e := range u.Errors {
					msgs = append(msgs, fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage))
				}
				return u, u.Status, fmt.Errorf("update %s: %s", u.Status, strings.Join(msgs, "; "))
			}
			return u, u.Status, nil
		},
		Timeout:    timeout,
		Delay:      addonDelay,
		MinTimeout: addonMinTimeout,
	}
	_, err := stateConf.WaitForState()
	return err
}
