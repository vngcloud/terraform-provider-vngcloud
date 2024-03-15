---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "VPC (Virtual Private Cloud)"
page_title: "vngcloud_vserver_route_table Resource - terraform-provider-vngcloud"
description: |-
  Provides a VNG Cloud VServer Route Table resource. This can be used to import, create, modify, and delete.
---

# vngcloud_vserver_route_table (Resource)



## Example Usage

```terraform
resource "vngcloud_vserver_route_table" "route_table" {
  project_id = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  name                    = "example-route-table"
  network_id               = "net-00fae478-aad2-4acb-bb10-32a4b337b45b"

  route {
    destination_cidr_block = "10.0.0.0/24"
    target                 = "10.0.0.1"
  }

  route {
    destination_cidr_block = "192.168.0.0/16"
    target                 = "192.168.0.1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `project_id` -  (String, Required) The ID of the project.
* `name` - (String, Required) The name of the route table.
* `network_id` -  (String, Required) The ID of the network associated with the route table.
* `route` - (Set of Maps, Optional)) A set of route definitions. Each route has the following attributes:
  * `destination_cidr_block` -  (String, Required) The destination CIDR block of the route.
  * `target` -  (String, Required) The target of the route.



## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `id` - (String) The ID of this route table.
* `status` - (String) The status of this route table.

## Import

Route table  can be imported using their unique identifier, e.g.
The unique identifier is the ID of the project and the ID of the route table, separated by a colon.
Example: `pro-26151c78-0470-4b4c-88a1-6ec41ef29492:rt-dcca77e6-4f5e-455b-9d9c-d0a0387a5ae2`
```
$ terraform import vngcloud_vserver_route_table.route_table pro-26151c78-0470-4b4c-88a1-6ec41ef29492:rt-dcca77e6-4f5e-455b-9d9c-d0a0387a5ae2
```
## IAM Policy
### Create:
In order to **create Route Table**, user must have been granted permissions below:
- CreateRouteTable
- GetDetailRouteTable

### Update
In order to **update Route Table**, user must have been granted permissions below:
-  EditListRouteFromRouteTable

### Delete
In order to **delete Route Table**, user must have been granted permissions below:
- DeleteRouteTable
