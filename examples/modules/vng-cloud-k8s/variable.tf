variable "project_id" {
  type    = string
  default = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "s_general_4x8" {
  type    = string
  default = "flav-05f97524-0410-46a4-87a8-af92aa759231"
}
variable "image_id" {
  type    = string
  default = "img-1c29f7df-fa23-4dd2-bcfb-9de14dee72e7"
}
variable "volume_type_name" {
  type    = string
  default = "3000"
}
variable "root_disk_size" {
  type    = number
  default = 20
}
variable "data_disk_size" {
  type    = number
  default = 20
}
variable "network_id" {
  type    = string
  default = "net--xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "subnet_id" {
  type    = string
  default = "sub--xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "ssh_key_id" {
  type    = string
  default = "ssh--xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "security_group_id_list" {
  type    = list(string)
  default = [
    "secg-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  ]
}