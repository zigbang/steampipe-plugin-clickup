package clickup

import (
	"context"
	"errors"
	"os"

	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*clickup.Client, error) {
	token := os.Getenv("CLICKUP_TOKEN")

	clickupConfig := GetConfig(d.Connection)
	if clickupConfig.Token != nil {
		token = *clickupConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	api := clickup.NewClient(nil, token)
	return api, nil
}
