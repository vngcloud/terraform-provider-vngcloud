---
subcategory: "Kubernetes Service"
page_title: "vngcloud_vks_addon_version Data Source - terraform-provider-vngcloud"
description: |-
  Finds a version of a VKS add-on for a Kubernetes version.
---

# vngcloud_vks_addon_version (Data Source)

Finds the version of a VKS add-on to install on a cluster of a given Kubernetes version.
Use it with [`vngcloud_vks_cluster_addon`](../resources/vks_cluster_addon.md).

## Example Usage

```terraform
data "vngcloud_vks_addon_version" "csi_default" {
  addon_name         = "block-storage-csi"
  kubernetes_version = "1.32"
}

data "vngcloud_vks_addon_version" "csi_latest" {
  addon_name         = "block-storage-csi"
  kubernetes_version = vngcloud_vks_cluster.primary.version
  most_recent        = true
}
```

## Argument Reference

- **addon_name** (String, Required): Name of the add-on.
- **kubernetes_version** (String, Required): Kubernetes version of the cluster, for example
  `1.32` or `v1.32.9`. Only the minor version is used.
- **most_recent** (Bool, Optional): When `true`, return the newest version compatible with the
  Kubernetes version. When `false` (default), return the version VKS recommends as the default
  for it.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String): Name of the add-on.
- **version** (String): The add-on version, for example `1.2.0-vks.1`.
- **default_namespace** (String): Namespace the version installs into by default.
- **requires_configuration** (Bool): When `true`, the add-on cannot be created without
  `configuration_values`.
- **type** (String): Type of the add-on.
- **owner** (String): Owner of the add-on.
- **publisher** (String): Publisher of the add-on.
