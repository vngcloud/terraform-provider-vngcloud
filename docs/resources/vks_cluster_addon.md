---
subcategory: "Kubernetes Service"
description: |-
  Manages an add-on of a VKS cluster.
---

# vngcloud_vks_cluster_addon

Manages an add-on (for example `block-storage-csi`) installed on a VNGCloud Kubernetes
Service (VKS) cluster. VKS installs the add-on as a Helm release and keeps it in sync
with the version and configuration values set here.

Use the [`vngcloud_vks_addon_version`](../data-sources/vks_addon_version.md) data source
to find the version to install for a Kubernetes version.

## Example Usage

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "my-vks-cluster"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
}

data "vngcloud_vks_addon_version" "csi" {
  addon_name         = "block-storage-csi"
  kubernetes_version = vngcloud_vks_cluster.primary.version
}

resource "vngcloud_vks_cluster_addon" "csi" {
  cluster_id    = vngcloud_vks_cluster.primary.id
  addon_name    = "block-storage-csi"
  addon_version = data.vngcloud_vks_addon_version.csi.version
}
```

## Example Usage - custom configuration values

```hcl
resource "vngcloud_vks_cluster_addon" "csi" {
  cluster_id                  = vngcloud_vks_cluster.primary.id
  addon_name                  = "block-storage-csi"
  addon_version               = "1.2.0-vks.1"
  resolve_conflicts_on_update = "OVERWRITE"

  configuration_values = jsonencode({
    controller = {
      replicaCount = 2
    }
  })
}
```

The values must match the JSON Schema of the add-on version. Print it with the CLI:
`grn vks get-addon-configuration-schema --addon-name block-storage-csi --addon-version 1.2.0-vks.1`.

## Argument Reference

* `cluster_id` - (Required, ForceNew) ID of the cluster.
* `addon_name` - (Required, ForceNew) Name of the add-on: 2-63 lowercase letters, digits or hyphens.
* `addon_version` - (Optional) Version of the add-on, for example `1.2.0-vks.1`. When not set,
  VKS installs the default version for the cluster's Kubernetes minor version. Changing it
  updates the add-on in place.
* `configuration_values` - (Optional) Custom configuration values as a JSON or YAML string
  (at most 64 KiB).
  Changing it updates the add-on in place; removing it clears the custom values.
* `namespace` - (Optional, ForceNew) Namespace to install the add-on into. Defaults to the
  default namespace of the add-on version.
* `resolve_conflicts_on_create` - (Optional) How to handle in-cluster resources that already
  exist when the add-on is created: `NONE`, `OVERWRITE` or `PRESERVE`.
* `resolve_conflicts_on_update` - (Optional) How to handle fields changed outside VKS when the
  add-on is updated: `NONE`, `OVERWRITE` or `PRESERVE`.
* `preserve` - (Optional) When `true`, destroying this resource stops VKS from managing the
  add-on but leaves its Helm release and workloads in the cluster. Default `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - `<cluster_id>:<addon_name>`.
* `status` - Status of the add-on: `CREATING`, `ACTIVE`, `DEGRADED`, `UPDATING`, `DELETING`,
  `CREATE_FAILED`, `UPDATE_FAILED` or `DELETE_FAILED`.
* `owner` - Owner of the add-on.
* `publisher` - Publisher of the add-on.
* `health_issues` - Issues reported for the add-on. Each item has `code`, `message` and
  `resource_ids`.
* `created_at` - Creation time (RFC 3339).
* `modified_at` - Last modification time (RFC 3339).

## Notes

* Create and update wait until the add-on is `ACTIVE` or `DEGRADED`. `DEGRADED` means the
  add-on is installed but its workloads are not ready yet — for example while the cluster
  has no node group. It turns `ACTIVE` on its own once the workloads are running, so
  Terraform does not wait for it.
* An update is asynchronous: Terraform waits for the add-on update to finish and fails with
  its errors if the update is `Failed` or `Cancelled`.
* `resolve_conflicts_on_create`, `resolve_conflicts_on_update` and `preserve` are not read
  back from the API; changing them alone does not call the API.
* The `enabled_load_balancer_plugin` and `enabled_block_store_csi_plugin` arguments of
  `vngcloud_vks_cluster` still work. Do not manage the same plugin with both that argument
  and this resource.

## Timeouts

* `create` - Default 20 minutes.
* `update` - Default 20 minutes.
* `delete` - Default 20 minutes.

## Import

An add-on can be imported with `<cluster_id>:<addon_name>`:

```shell
terraform import vngcloud_vks_cluster_addon.csi k8s-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx:block-storage-csi
```
