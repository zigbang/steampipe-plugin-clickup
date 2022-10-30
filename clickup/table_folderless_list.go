package clickup

import (
    "context"
    "fmt"

    "github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupFolderlessList() *plugin.Table {
    return &plugin.Table{
        Name: "clickup_folderless_list",
        List: &plugin.ListConfig{
            KeyColumns: plugin.SingleColumn("space_id"),
            Hydrate:    listFolderlessLists,
        },
        Columns: folderlessColumns(),
    }
}

func listFolderlessLists(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    spaceId := d.KeyColumnQuals["space_id"].GetStringValue()

    client, err := connect(ctx, d)
    if err != nil {
        return nil, fmt.Errorf("unable to establish a connection: %v", err)
    }

    items, _, err := client.Lists.GetFolderlessLists(ctx, spaceId, false)
    if err != nil {
        return nil, fmt.Errorf("unable to obtain folderless lists for space id '%s': %v", spaceId, err)
    }

    for _, t := range items {
        d.StreamListItem(ctx, t)
    }

    return nil, nil
}

func folderlessColumns() []*plugin.Column {
    return []*plugin.Column{
        {
            Name:      "space_id",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromQual("space_id"),
        },
        {
            Name: "id",
            Type: proto.ColumnType_STRING,
        },
        {
            Name: "name",
            Type: proto.ColumnType_STRING,
        },
        {
            Name:      "order_index",
            Type:      proto.ColumnType_INT,
            Transform: transform.FromField("Orderindex"),
        },
        {
            Name: "content",
            Type: proto.ColumnType_STRING,
        },
        {
            Name:      "status",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Status.Status"),
        },
        {
            Name:      "priority",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Priority.Priority"),
        },
        {
            Name:      "assignee",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Assignee.Username"),
        },
        {
            Name:      "assignee_email",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Assignee.Email"),
        },
        {
            Name: "task_count",
            Type: proto.ColumnType_STRING,
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
            Name:      "folder_id",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Folder.ID"),
        },
        {
            Name:      "folder",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Folder.Name"),
        },
        {
            Name:      "space",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromField("Space.Name"),
        },
        {
            Name: "archived",
            Type: proto.ColumnType_BOOL,
        },
        {
            Name: "permission_level",
            Type: proto.ColumnType_STRING,
        },
    }
}
