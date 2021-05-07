#bin/bash
export TF_LOG=debug
rm -rf .terraform
rm -r .terraform.lock.hcl
go build -o terraform-provider-vserver
OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
echo $OS_ARCH
rm -rf ~/.terraform.d/plugins/vngcloud.com.vn/terraform/vserver/0.2
mkdir -p  ~/.terraform.d/plugins/vngcloud.com.vn/terraform/vserver/0.2/$OS_ARCH
mv terraform-provider-vserver ~/.terraform.d/plugins/vngcloud.com.vn/terraform/vserver/0.2/$OS_ARCH
terraform init
terraform apply