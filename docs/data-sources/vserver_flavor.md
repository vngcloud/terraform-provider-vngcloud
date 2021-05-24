---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vngcloud_vserver_flavor Data Source - terraform-provider-vngcloud"
subcategory: ""
description: |-
  
---

# vngcloud_vserver_flavor (Data Source)



## Example Usage

```terraform
data "vngcloud_vserver_flavor" "flavor" {
    name = "dev.v1.small1x1"
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    flavor_zone_id =  "9BEC387B-0DC1-41D1-B177-5FDC72ED7B0C"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **flavor_zone_id** (String)
- **name** (String)
- **project_id** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **bandwidth** (Number)
- **bandwidth_unit** (String)
- **cpu** (Number)
- **gpu** (Number)
- **memory** (Number)

