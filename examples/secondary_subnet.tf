resource "vngcloud_vserver_subnet" "secondary_demo" {
    project_id = var.project_id
    name       = "tf-secondary-demo"
    cidr       = "10.76.20.0/24"
    network_id = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    zone_id    = "HCM03-1A"

    secondary_subnet {
        name = "sec-a"
        cidr = "10.76.21.0/24"
    }
}
