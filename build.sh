#bin/bash
#export TF_LOG=debug
rm -rf .terraform
rm -rf .terraform.*
#rm -rf terraform.*
go build -o terraform-provider-vngcloud
#OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
OS_ARCH="darwin_amd64"
echo $OS_ARCH
#rm -rf ~/.terraform.d/plugins/vngcloud.vn/terraform/vngcloud/0.2
#mkdir -p  ~/.terraform.d/plugins/vngcloud.vn/terraform/vngcloud/0.2/$OS_ARCH
#mv terraform-provider-vngcloud ~/.terraform.d/plugins/vngcloud.vn/terraform/vngcloud/0.2/$OS_ARCH
rm -rf ~/.terraform.d/plugins/registry.terraform.tunm4/vngcloud/vngcloud/1.0.0
mkdir -p  ~/.terraform.d/plugins/registry.terraform.tunm4/vngcloud/vngcloud/1.0.0/$OS_ARCH
mv terraform-provider-vngcloud ~/.terraform.d/plugins/registry.terraform.tunm4/vngcloud/vngcloud/1.0.0/$OS_ARCH
terraform init
#terraform apply