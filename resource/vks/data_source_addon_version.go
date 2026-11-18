package vks

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
)

// DataSourceAddonVersion picks a version of a catalog add-on for a Kubernetes
// version (like aws_eks_addon_version).
func DataSourceAddonVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAddonVersionRead,
		Schema: map[string]*schema.Schema{
			"addon_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(addonNameRegexp, "must be 2-63 lowercase alphanumeric characters or hyphens"),
			},
			"kubernetes_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kubernetes version of the cluster, e.g. 1.32 or v1.32.9.",
			},
			"most_recent": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Return the newest compatible version instead of the default one.",
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requires_configuration": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"type": {
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
		},
	}
}

// kubernetesMinor turns "v1.32.9" / "1.32" into "1.32".
func kubernetesMinor(version string) string {
	parts := strings.SplitN(strings.TrimPrefix(version, "v"), ".", 3)
	if len(parts) < 2 {
		return strings.TrimPrefix(version, "v")
	}
	return parts[0] + "." + parts[1]
}

// selectAddonVersion returns the default version for minor or, with mostRecent,
// the first compatible one (the API lists versions newest first).
func selectAddonVersion(versions []vks.AddonVersionInfoDto, minor string, mostRecent bool) *vks.AddonVersionInfoDto {
	for i := range versions {
		v := &versions[i]
		for _, c := range v.Compatibilities {
			if c.ClusterVersion != minor {
				continue
			}
			if mostRecent || c.DefaultVersion {
				return v
			}
		}
	}
	return nil
}

func dataSourceAddonVersionRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	addonName := d.Get("addon_name").(string)
	minor := kubernetesMinor(d.Get("kubernetes_version").(string))

	resp, httpResponse, err := cli.VksClient.V1AddonCatalogControllerApi.V1AddonVersionsGet(context.TODO(), &vks.V1AddonCatalogControllerApiV1AddonVersionsGetOpts{
		AddonName:         optional.NewString(addonName),
		KubernetesVersion: optional.NewString(minor),
	})
	if err != nil || httpResponse == nil || CheckErrorResponse(httpResponse) {
		return addonAPIError(fmt.Sprintf("listing versions of add-on %s", addonName), httpResponse, err)
	}

	var info *vks.AddonInfoDto
	for i := range resp.Addons {
		if resp.Addons[i].AddonName == addonName {
			info = &resp.Addons[i]
			break
		}
	}
	if info == nil {
		return fmt.Errorf("add-on %s not found in the catalog", addonName)
	}
	version := selectAddonVersion(info.AddonVersions, minor, d.Get("most_recent").(bool))
	if version == nil {
		return fmt.Errorf("no version of add-on %s is compatible with Kubernetes %s (or none is marked default; try most_recent = true)", addonName, minor)
	}

	namespace := version.DefaultNamespace
	if namespace == "" {
		namespace = info.DefaultNamespace
	}
	d.SetId(addonName)
	d.Set("version", version.AddonVersion)
	d.Set("default_namespace", namespace)
	d.Set("requires_configuration", version.RequiresConfiguration)
	d.Set("type", info.Type_)
	d.Set("owner", info.Owner)
	d.Set("publisher", info.Publisher)
	return nil
}
