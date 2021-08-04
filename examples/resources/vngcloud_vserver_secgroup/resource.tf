resource "vngcloud_vserver_secgroup" "secgroup" {
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "name-secgroup"
    lifecycle {
        create_before_destroy = true
    }
}