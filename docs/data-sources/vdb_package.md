---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vngcloud_vdb_package Data Source - terraform-provider-vngcloud"
description: |-
  Provides a VNG Cloud VServer Flavor zone.
---

# vngcloud_vdb_package (Data Source)



## Example Usage

```terraform
data "vngcloud_vdb_package" "package" {
    engine_type = "MySQL"
    engine_version = "5.7"
    name = "v1.small1x1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **engine_type** (String)
- **engine_version** (String)
- **name** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **flavor_id** (String)
- **package_id** (String)
- **price_key** (String)
- **ram** (Number)
- **vcpus** (Number)
- **zone_id** (Number)



