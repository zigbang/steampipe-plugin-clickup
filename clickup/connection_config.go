package clickup

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type clickupConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {Type: schema.TypeString},
}

func ConfigInstance() interface{} {
	return &clickupConfig{}
}

func GetConfig(connection *plugin.Connection) clickupConfig {
	if connection == nil || connection.Config == nil {
		return clickupConfig{}
	}
	config, _ := connection.Config.(clickupConfig)
	return config
}
