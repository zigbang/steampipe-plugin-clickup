package main

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/ygpark80/steampipe-plugin-clickup/clickup"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: clickup.Plugin})
}
