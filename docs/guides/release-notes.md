---
page_title: "Release Notes"
description: |-
  What changed in each release of the VNG Cloud Terraform provider.
---

# Release Notes

## v1.3.19

### Enhancements

- `vngcloud_vks_cluster` now accepts `list_subnet_ids` for both `SINGLE` and `MULTI` `az_strategy`
  (previously `MULTI`-only). `subnet_id` still works for backward compatibility — set one, not both.
  A `node_group` without its own `subnet_id` now correctly inherits the cluster's subnet either way.
- `auto_scale_config` update now distinguishes three cases based on the change relative to state:
  no `auto_scale_config` before and still none → unchanged, nothing sent; existing config removed
  from HCL → explicit disable sent; block present with changes → `min_size`/`max_size` updated.

### Bug Fixes

- Fixed spurious "force replacement" on `vngcloud_vks_cluster` configured with only
  `list_subnet_ids` (no `subnet_id`).
- Fixed cluster delete failing when node groups were still attached — they're now removed first.
- Fixed `auto_healing_config` not clearing in state when the API reports it as unset.
- `secondary_subnets` (`CILIUM_NATIVE_ROUTING`) is now backend-owned/computed, avoiding a stale-value
  force-replace.

## v1.3.18

### Enhancements

- `taint` on `vngcloud_vks_cluster_node_group` (and the inline `node_group` block inside
  `vngcloud_vks_cluster`) can now also be written as a list attribute, in addition to the
  existing repeated `taint { ... }` block syntax (Terraform supports both forms for this kind
  of field — see [Attributes as Blocks](https://developer.hashicorp.com/terraform/language/attr-as-blocks)).
  **Existing configs using `taint { ... }` blocks are unaffected and require no changes.**

  The new attribute form makes `taint = []` valid to explicitly clear all taints — something
  block syntax can never express, since omitting all `taint { ... }` blocks is indistinguishable
  from never having configured taints at all:

  ```hcl
  taint = [
    {
      key    = "key1"
      value  = "value1"
      effect = "PreferNoSchedule"
    },
    {
      key    = "key2"
      value  = "value2"
      effect = "NoExecute"
    }
  ]
  ```

  See "Clearing all taints" in
  [vks_cluster_node_group](https://registry.terraform.io/providers/vngcloud/vngcloud/latest/docs/resources/vks_cluster_node_group#clearing-all-taints)
  for details. Note: a single resource cannot mix both forms for the same `taint` argument —
  use either `taint { ... }` blocks or a `taint = [...]` list, not both.
