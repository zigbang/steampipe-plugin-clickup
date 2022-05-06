package clickup

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableClickupTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_team",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "color", Type: proto.ColumnType_STRING},
			{Name: "avatar", Type: proto.ColumnType_STRING},
		},
	}
}

func listTeams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	client, _ := connect(ctx, d)
	teams, _, _ := client.Teams.GetTeams(ctx)

	for _, t := range teams {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}
