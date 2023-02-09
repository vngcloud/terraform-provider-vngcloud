variable "client_id" {
  type = string
  default = "9e734abc-7171-4b91-aab1-f652d2147f83"
}
variable "client_secret" {
  type = string
  default = "29ee24e8-c1b8-4b18-addb-f0c65689a2a8"
}
variable "project_id" {
  type = string
  default = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
}

variable "image_id" {
  type = string
  default = "img-2aba43f5-281a-4978-81ee-824e68cb0fb4"
}
variable "flavor_id" {
  type = string
  default = "flav-b9b44c1b-5823-4069-a0a2-ffd5acbeccf7"
}

variable "volume_type_name" {
  type = string
  default = "200"
}
variable "root_disk_size" {
  type = number
  default = 20
}
variable "data_disk_size" {
  type = number
  default = 20
}
variable "network_id" {
  type = string
  default = "net-910ea9c5-cbb5-45af-860e-25d4b6550d53"
}
variable "subnet_id" {
  type = string
  default = "sub-0c71ba69-5b55-44db-95ce-46ddab94054c"
}
variable "ssh_key_id" {
  type = string
  default = "ssh-d07ddb73-19e6-43d0-8d13-2f291a793443"
}
variable "security_group_id_list" {
  type = list(string)
  default = ["secg-5c1e8e96-d106-4a91-8047-a76d1ae5ba9b", "secg-28aee6f4-a4f7-4a8c-a3f5-4f739ea14dcb"]
}

variable "server_count" {
  type = number
  default = 1
}

