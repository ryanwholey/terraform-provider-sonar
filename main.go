package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/ryanwholey/terraform-provider-sonar/sonar"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sonar.Provider})
}
