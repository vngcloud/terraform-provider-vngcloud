resource "vngcloud_vserver_network" "network" {
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "name-network"
    cidr = "10.76.0.0/16"
    lifecycle {
        create_before_destroy = true
    }
}