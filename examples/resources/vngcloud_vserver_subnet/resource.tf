resource "vngcloud_vserver_subnet" "subnet" {
  project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
  name       = "name-subnet"
  cidr       = "10.76.1.0/24"
  network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
  lifecycle {
    create_before_destroy = true
  }
}