cp ~/go/bin/terraform-provider-sonar ~/.terraform.d/plugins/darwin_amd64/
rm -rf .terraform
rm -rf terraform.tfstate*
terraform init
