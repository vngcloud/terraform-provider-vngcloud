# VNG Cloud Terraform Examples

Hướng dẫn chạy Terraform để triển khai tài nguyên trên VNG Cloud.

## Cấu trúc thư mục

```
examples/
├── main.tf                        # Root module: provider, backend, và gọi các module
├── variable.tf                    # Biến cho root module (client_id, client_secret)
├── terraform.tfvars               # Giá trị credentials (KHÔNG commit lên Git)
├── .vngcloud_credentials          # Credentials cho S3 backend (KHÔNG commit lên Git)
└── modules/
    ├── vng-cloud-vserver/         # Module tạo VM, Volume, Server Group
    ├── vng-cloud-vlb/             # Module tạo Load Balancer, Pool, Listener
    ├── vng-cloud-k8s/             # Module tạo Kubernetes cluster
    └── vng-cloud-vdb/             # Module tạo Database (MySQL, Redis)
```

## Yêu cầu

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.3
- Tài khoản VNG Cloud với IAM credentials (Client ID & Client Secret)
- (Tuỳ chọn) vStorage bucket nếu dùng S3 backend để lưu state

## Cấu hình Credentials

Hai file dưới đây chứa thông tin nhạy cảm và **không được commit lên Git**. Bạn cần tự tạo chúng trên máy local.

---

### 1. `terraform.tfvars` — IAM Credentials (bắt buộc)

File này cung cấp `client_id` và `client_secret` cho VNG Cloud provider.

```bash
cat > terraform.tfvars << 'EOF'
# Thông tin nhạy cảm: không commit vào Git
client_id     = "<your-client-id>"
client_secret = "<your-client-secret>"
EOF
```

> Lấy `client_id` và `client_secret` tại [VNG Cloud IAM Portal](https://iam.console.vngcloud.vn/) > **Service Account** > **Create Service Account** > copy Client ID & Client Secret.

---

### 2. `.vngcloud_credentials` — S3 Backend Credentials (nếu dùng remote state)

File này chứa Access Key / Secret Key để Terraform đọc/ghi state file trên vStorage.

```bash
cat > .vngcloud_credentials << 'EOF'
# Thông tin nhạy cảm: không commit vào Git
export AWS_ACCESS_KEY_ID="<vstorage-access-key>"
export AWS_SECRET_ACCESS_KEY="<vstorage-secret-key>"
EOF
```

> Lấy Access Key / Secret Key tại [VNG Cloud vStorage Portal](https://vstorage.console.vngcloud.vn/) > **S3 Keys** > **Create S3 Key**.

**Để tắt S3 backend** (dùng local state thay thế), comment hoặc xóa block `backend "s3"` trong [main.tf](main.tf).

## Chạy Terraform

### Bước 1: Init

```bash
cd examples/

# Nếu dùng S3 backend, export credentials trước
source .vngcloud_credentials

terraform init
```

### Bước 2: Chọn module muốn chạy

Mặc định [main.tf](main.tf) đang bật module `vserver`. Bật/tắt module bằng cách comment/uncomment trong `main.tf`:

```hcl
module "vserver" {
  source = "./modules/vng-cloud-vserver"
}

# module "vlb" {
#   source = "./modules/vng-cloud-vlb"
# }
```

### Bước 3: Plan

```bash
terraform plan
```

### Bước 4: Apply

```bash
terraform apply
```

Xác nhận bằng cách nhập `yes` khi được hỏi.

### Bước 5: Destroy (dọn dẹp tài nguyên)

```bash
terraform destroy
```

> **Lưu ý:** Terraform tự động load `terraform.tfvars` nên không cần chỉ định `-var-file`. Flag này chỉ cần khi dùng file tên khác, ví dụ: `terraform apply -var-file="prod.tfvars"`.

## Mô tả các Module

### vng-cloud-vserver

Tạo tài nguyên vServer tại region HCM-3:

| Resource | Mô tả |
|---|---|
| `vngcloud_vserver_server_group` | Server Group với policy Soft Affinity / Soft Anti-Affinity |
| `vngcloud_vserver_server` | Các VM (mặc định: `terra-26`, `terra-56`) chạy Ubuntu 24.04 |
| `vngcloud_vserver_volume` | Data disk 50GB đính kèm mỗi VM |
| `vngcloud_vserver_volume_attach` | Gắn volume vào server |

Trước khi chạy, mở [modules/vng-cloud-vserver/variable.tf](modules/vng-cloud-vserver/variable.tf) và cập nhật các biến sau:

**Bắt buộc thay đổi theo môi trường của bạn:**

| Biến | Lấy ở đâu | Định dạng |
|---|---|---|
| `project_id` | Portal > góc trên phải > Project ID | `pro-xxxxxxxx-...` |
| `network_id` | vServer Portal > Network > VPC > copy ID | `net-xxxxxxxx-...` |
| `subnet_id` | vServer Portal > Network > VPC > tab Subnets > copy ID | `sub-xxxxxxxx-...` |
| `ssh_key_id` | vServer Portal > SSH Keys > copy ID | `ssh-xxxxxxxx-...` |
| `security_group_id_list` | vServer Portal > Security Groups > copy ID | `["secg-xxxxxxxx-..."]` |

**Có thể điều chỉnh tuỳ nhu cầu:**

| Biến | Mô tả | Mặc định |
|---|---|---|
| `server_names` | Danh sách tên VM cần tạo | `["terra-26", "terra-56"]` |
| `root_disk_size` | Dung lượng root disk (GB) | `20` |
| `data_disk_size` | Dung lượng data disk đính kèm (GB) | `50` |

### vng-cloud-vlb

Tạo Load Balancer Layer 7:

| Resource | Mô tả |
|---|---|
| `vngcloud_vlb_load_balancer` | Internet-facing Layer 7 LB |
| `vngcloud_vlb_pool` | Pool HTTP với health check và 3 member |
| `vngcloud_vlb_listener` | Listener HTTP port 81 |
| `vngcloud_vlb_l7policy` | L7 policy redirect theo path |

### vng-cloud-k8s

Tạo Kubernetes cluster (VKS):

| Resource | Mô tả |
|---|---|
| `vngcloud_vserver_k8s` | K8S cluster với auto-scaling (1-3 node) |
| `vngcloud_vserver_cluster_node_group` | Node group bổ sung |
| `vngcloud_vserver_change_cluster_sec_group` | Gán Security Group cho master/minion |

### vng-cloud-vdb

Tạo Database instance:

| Resource | Mô tả |
|---|---|
| `vngcloud_vdb_relational_database` | MySQL 8 instance |
| `vngcloud_vdb_memstore_database` | Redis instance |
| `vngcloud_vdb_relational_backup` | Backup cho MySQL |
| `vngcloud_vdb_memstore_backup` | Backup cho Redis |
| `vngcloud_vdb_relational_config_group` | Config group cho MySQL |
| `vngcloud_vdb_memstore_config_group` | Config group cho Redis |

## Lưu ý bảo mật

- **KHÔNG** commit `terraform.tfvars` và `.vngcloud_credentials` lên Git vì chứa thông tin nhạy cảm.
- Đảm bảo các file này đã được thêm vào `.gitignore`.
- Sử dụng biến môi trường hoặc Vault để quản lý secrets trong môi trường CI/CD.
