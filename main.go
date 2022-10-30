package main

import (
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/zigbang/steampipe-plugin-clickup/clickup"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{PluginFunc: clickup.Plugin})
}
