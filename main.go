package main

import (
	"terraform-provider-cds/dyl"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dyl.Provider})

}
