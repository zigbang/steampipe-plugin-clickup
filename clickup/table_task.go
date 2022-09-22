package clickup

import (
	"context"

	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableClickupTask(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_task",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("list_id"),
			Hydrate:    listTasks,
		},
		Columns: []*plugin.Column{
			{Name: "list_id", Type: proto.ColumnType_STRING, Hydrate: listId, Transform: transform.FromValue()},
			{Name: "id", Type: proto.ColumnType_STRING},
			// custom_id
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "text_content", Type: proto.ColumnType_STRING},
			{Name: "description", Type: proto.ColumnType_STRING},
			// status
			{Name: "overrideindex", Type: proto.ColumnType_STRING},
			{Name: "date_created", Type: proto.ColumnType_STRING},
			{Name: "date_updated", Type: proto.ColumnType_STRING},
			// date_closed
			// creator
			// assignees
			// checklists
			// tags
			// parent
			// priority
			// due_date
			// start_date
			// time_estimate
			// time_spent
			// custom_fields
			// list
			// folder
			// space
			{Name: "url", Type: proto.ColumnType_STRING},
		},
	}
}

func listTasks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)
	client, _ := connect(ctx, d)

	list_id := d.KeyColumnQuals["list_id"].GetStringValue()

	opts := &clickup.GetTasksOptions{
		Page:          0,
		Archived:      true,
		IncludeClosed: true,
	}

	for {
		items, _, _ := client.Tasks.GetTasks(ctx, list_id, opts)
		for _, t := range items {
			d.StreamListItem(ctx, t)
		}

		if len(items) < 100 {
			break
		}

		opts.Page = opts.Page + 1
	}

	return nil, nil
}

func listId(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["list_id"].GetStringValue()
	return q, nil
}
