---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "VPC (Virtual Private Cloud)"
page_title: "vngcloud_vserver_network Resource - terraform-provider-vngcloud"
description: |-
  Provides a VNG Cloud VServer Network resource. This can be used to import, create, modify, and delete.
---

# vngcloud_vserver_network (Resource)



## Example Usage

```terraform
resource "vngcloud_vserver_network" "network" {
    project_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    name = "example-network"
    cidr = "10.76.0.0/16"
    zone_id = "HCM03-1A"
}
```

## Argument Reference

The following arguments are supported:

- **cidr** (String, Required) Classless Inter-Domain Routing of Network`(\16)`. Please choose below options:
  - 10.0.0.0 - 10.255.0.0
  - 172.16.0.0 - 172.24.0.0
  - 192.168.0.0
- **name** (String, Required) Name of Network.
- **project_id** (String, Required) ID of project.
- **route_table_default_id** (String, Optional) ID of default route table.
- **zone_id** (String, Optional) ID of Zone. If not specified, the default zone will be used

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
- **id** (String) The ID of this Network.

