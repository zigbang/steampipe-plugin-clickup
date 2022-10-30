package clickup

import (
    "context"
    "fmt"

    "github.com/raksul/go-clickup/clickup"
    "github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupTask() *plugin.Table {
    return &plugin.Table{
        Name: "clickup_task",
        List: &plugin.ListConfig{
            KeyColumns: plugin.SingleColumn("list_id"),
            Hydrate:    listTasks,
        },
        Columns: taskColumns(),
    }
}

func listTasks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    client, err := connect(ctx, d)
    if err != nil {
        return nil, fmt.Errorf("unable to establish a connection: %v", err)
    }

    listId := d.KeyColumnQuals["list_id"].GetStringValue()

    opts := &clickup.GetTasksOptions{
        Page:          0,
        Archived:      true,
        IncludeClosed: true,
    }

    for {
        items, _, err := client.Tasks.GetTasks(ctx, listId, opts)
        if err != nil {
            return nil, fmt.Errorf("unable to obtain tasks for list id '%s': %v", listId, err)
        }

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

func taskColumns() []*plugin.Column {
    return []*plugin.Column{
        {
            Name:      "list_id",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromQual("list_id"),
        },
        {
            Name: "id",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "custom_id",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "name",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "text_content",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "description",
            Type: proto.ColumnType_STRING,
        },
        {
            Name:      "status",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Status.Status"),
        },
        {
            Name:      "order_index",
            Type:      proto.ColumnType_INT,
            Transform: transform.FromField("Orderindex"),
        },
        {
            Name: "date_created",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "date_updated",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "date_closed",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "creator",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "assignees",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "checklists",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "tags",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "parent",
            Type: proto.ColumnType_STRING,
        },
        {
            Name:      "priority",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Priority.Priority"),
        },
        {
            Name: "due_date",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "start_date",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "points",
            Type: proto.ColumnType_INT,
        },
        {
            Name: "time_estimate",
            Type: proto.ColumnType_INT,
        },
        {
            Name: "team_id",
            Type: proto.ColumnType_STRING,
        },
        {
            Name:      "url",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("URL"),
        },
        {
            Name: "permission_level",
            Type: proto.ColumnType_STRING,
        },
    }
}
