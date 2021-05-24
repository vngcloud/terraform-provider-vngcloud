resource "vngcloud_vserver_volume" "volume" {
    name = "volume_name"
    size = 20
    volume_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    lifecycle {
        create_before_destroy = true
    } 
}