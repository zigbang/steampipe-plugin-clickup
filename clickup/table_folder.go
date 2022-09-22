package clickup

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableClickupFolder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_folder",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("space_id"),
			Hydrate:    listFolders,
		},
		Columns: []*plugin.Column{
			{Name: "space_id", Type: proto.ColumnType_STRING, Hydrate: spaceId, Transform: transform.FromValue()},
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "override_statuses", Type: proto.ColumnType_BOOL},
			{Name: "hidden", Type: proto.ColumnType_BOOL},
			// spapce
			{Name: "task_count", Type: proto.ColumnType_STRING},
			// lists
		},
	}
}

func listFolders(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	space_id := d.KeyColumnQuals["space_id"].GetStringValue()

	client, _ := connect(ctx, d)
	items, _, _ := client.Folders.GetFolders(ctx, space_id, false)

	for _, t := range items {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}

func spaceId(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["space_id"].GetStringValue()
	return q, nil
}
