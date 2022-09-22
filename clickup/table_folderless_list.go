package clickup

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableClickupFolderlessList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_folderless_list",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("space_id"),
			Hydrate:    listFolderlessLists,
		},
		Columns: []*plugin.Column{
			{Name: "space_id", Type: proto.ColumnType_STRING, Hydrate: tableClickupFolderlessListKey, Transform: transform.FromValue()},
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "overrideindex", Type: proto.ColumnType_INT},
			{Name: "content", Type: proto.ColumnType_STRING},
			// status
			// priority
			// assignee
			{Name: "task_count", Type: proto.ColumnType_STRING},
			// due_date
			// start_date
			// folder
			// space
			{Name: "archived", Type: proto.ColumnType_BOOL},
			{Name: "override_statuses", Type: proto.ColumnType_BOOL},
			{Name: "permission_level", Type: proto.ColumnType_STRING},
		},
	}
}

func listFolderlessLists(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	space_id := d.KeyColumnQuals["space_id"].GetStringValue()

	client, _ := connect(ctx, d)
	items, _, _ := client.Lists.GetFolderlessLists(ctx, space_id, false)

	for _, t := range items {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}

func tableClickupFolderlessListKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["space_id"].GetStringValue()
	return q, nil
}
