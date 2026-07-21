---
page_title: "Release Notes"
description: |-
  What changed in each release of the VNG Cloud Terraform provider.
---

# Release Notes

## v1.3.16

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

- `taint` diffing is now order-independent and content-based: reordering `taint` entries in
  config no longer produces a plan diff, and two entries with identical `key`/`value`/`effect`
  are treated as the same taint (duplicates are not sent to the API twice). Previously, taint
  order mattered and could cause spurious in-place updates.
