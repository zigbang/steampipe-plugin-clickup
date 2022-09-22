package clickup

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableClickupSpace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_space",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("team_id"),
			Hydrate:    listSpaces,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING, Hydrate: teamId, Transform: transform.FromValue()},
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "private", Type: proto.ColumnType_BOOL},
			// statuses
			{Name: "multiple_assignees", Type: proto.ColumnType_BOOL},
			// features
		},
	}
}

func listSpaces(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	team_id := d.KeyColumnQuals["team_id"].GetStringValue()

	client, _ := connect(ctx, d)
	spaces, _, _ := client.Spaces.GetSpaces(ctx, team_id)

	for _, t := range spaces {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}

func teamId(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["team_id"].GetStringValue()
	return q, nil
}
