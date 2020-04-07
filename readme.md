# terraform-provider-sonar

A Terraform provider to manage state for the SONAR deployment server
 
## Setup

Supply the provider with an endpoint and credentials using environment variables:

```sh
export SONAR_API_URL=http://localhost:7734/
export SONAR_ACCESS_TOKEN=
```

Build, install, and apply

```sh
make build
cp $GOPATH/bin/terraform-provider-sonar ~/.terraform.d/plugins
cd example
terraform init
terraform apply
```
