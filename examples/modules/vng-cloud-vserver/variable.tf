
# =============================================================================
# CÁC THÔNG SỐ CẦN THAY ĐỔI THEO MÔI TRƯỜNG CỦA BẠN
# Xem hướng dẫn lấy ID tại: https://docs.vngcloud.vn
# =============================================================================

# [BẮT BUỘC] Project ID của bạn trên VNG Cloud
# Lấy tại: https://hcmportal.vngcloud.vn > góc trên phải > Project ID
# Định dạng: pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
variable "project_id" {
  type    = string
  default = "pro-462803f3-6858-466f-bf05-df2b33faa360"
}

# [BẮT BUỘC] Network ID của VPC muốn đặt VM vào
# Lấy tại: vServer Portal > Network > VPC > chọn VPC > copy ID
# Định dạng: net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
variable "network_id" {
  type    = string
  default = "net-4ca7ce61-b2cd-47cf-9c91-b62e4380562e"
}

# [BẮT BUỘC] Subnet ID nằm trong VPC trên
# Lấy tại: vServer Portal > Network > VPC > chọn VPC > tab Subnets > copy ID
# Định dạng: sub-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
variable "subnet_id" {
  type    = string
  default = "sub-2d6a2e2b-8e00-49b2-b8de-38f391d04c42"
}

# [BẮT BUỘC] SSH Key ID để truy cập VM sau khi tạo
# Lấy tại: vServer Portal > SSH Keys > tạo mới hoặc upload key > copy ID
# Định dạng: ssh-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
variable "ssh_key_id" {
  type    = string
  default = "ssh-c853d8fb-54a2-4990-9932-e2591369fc83"
}

# [BẮT BUỘC] Danh sách Security Group ID áp dụng cho VM
# Lấy tại: vServer Portal > Security Groups > chọn group > copy ID
# Định dạng: secg-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
# Có thể thêm nhiều security group: ["secg-aaa...", "secg-bbb..."]
variable "security_group_id_list" {
  type    = list(string)
  default = [
    "secg-b8ca3404-b25d-4fcc-b1ec-2aff86494cc4"
  ]
}

# =============================================================================
# CÁC THÔNG SỐ CÓ THỂ GIỮ NGUYÊN HOẶC ĐIỀU CHỈNH
# =============================================================================

# Tên các VM sẽ được tạo (mỗi phần tử = 1 VM)
variable "server_names" {
  description = "List of server names to create"
  type        = list(string)
  default     = ["terra-26", "terra-56"]
}

# Dung lượng root disk (GB) — tối thiểu 20GB
variable "root_disk_size" {
  type    = number
  default = 20
}

# Dung lượng data disk đính kèm mỗi VM (GB)
variable "data_disk_size" {
  type    = number
  default = 50
}

# =============================================================================
# THÔNG SỐ HẠ TẦNG — thường không cần đổi nếu dùng cùng region HCM-3
# Lấy ID tại: vServer Portal > Flavors / Images / Volume Types
# =============================================================================

# Flavor: s-general-4vCPU-8GB RAM tại zone 1C
variable "s_general_4x8_zone1c" {
  type    = string
  default = "flav-d1e5634e-0565-11f0-a0a4-ec2a72332f83"
}

# Image: Ubuntu 24.04
variable "ubuntu_24_04" {
  type    = string
  default = "img-34440a82-92fb-40bc-b79c-b1a2b49b93de"
}

# Volume type: NVMe 3000 IOPS tại zone 1C
variable "volume_type_3000_1c" {
  type    = string
  default = "vtype-e782f8e1-0569-11f0-a0a4-ec2a72332f83"
}

# Server Group policy — giữ nguyên tên, không cần đổi
variable "soft_affinity" {
  type    = string
  default = "SOFT AFFINITY"
}

variable "soft_anti_affinity" {
  type    = string
  default = "SOFT ANTI AFFINITY"
}

# Có encode user_data sang base64 không (false = plain text)
variable "user_data_base64_encode" {
  type    = bool
  default = false
}
