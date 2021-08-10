resource "vngcloud_vserver_server" "server" {
  project_id        = "pro-462803f3-6858-466f-bf05-df2b33faa360"
  name              = "name-server"
  encryption_volume = false
  flavor_id         = "flav-437f64a6-f55e-4f42-b861-65bcc62420de"
  image_id          = "img-d0abd877-4e4e-4682-8ab6-9fe6d3474d40"
  network_id        = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
  root_disk_size    = 20
  root_disk_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
  ssh_key           = "ssh-dd96293a-49a4-4f8d-ac9a-033a23809e1d"
  security_group    = ["secg-2c2be529-58fb-4fd5-9e6e-bef82f7b4663"]
  subnet_id         = "sub-a5e4f8e9-e99e-498c-b99a-6b4720cf5c6f"
  lifecycle {
    create_before_destroy = true
  }
}