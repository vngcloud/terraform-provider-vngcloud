---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vngcloud_vserver_flavor_zone Data Source - terraform-provider-vngcloud"
subcategory: ""
description: |-
  
---

# vngcloud_vserver_flavor_zone (Data Source)



## Example Usage

```terraform
data "vngcloud_vserver_flavor_zone" "flavor_zone" {
    name = "DEV General v1 Instances"
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Name of flavor zone.
- **project_id** (String) ID of project

### Optional

- **id** (String) The ID of this resource.

