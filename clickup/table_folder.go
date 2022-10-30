package clickup

import (
    "context"
    "fmt"

    "github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupFolder() *plugin.Table {
    return &plugin.Table{
        Name: "clickup_folder",
        List: &plugin.ListConfig{
            KeyColumns: plugin.SingleColumn("space_id"),
            Hydrate:    listFolders,
        },
        Columns: folderColumns(),
    }
}

func listFolders(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    spaceId := d.KeyColumnQuals["space_id"].GetStringValue()

    client, err := connect(ctx, d)
    if err != nil {
        return nil, fmt.Errorf("unable to establish a connection: %v", err)
    }

    items, _, err := client.Folders.GetFolders(ctx, spaceId, false)
    if err != nil {
        return nil, fmt.Errorf("unable to obtain folders for space id '%s': %v", spaceId, err)
    }

    for _, t := range items {
        d.StreamListItem(ctx, t)
    }

    return nil, nil
}

func folderColumns() []*plugin.Column {
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
            Name: "override_statuses",
            Type: proto.ColumnType_BOOL,
        },
        {
            Name: "hidden",
            Type: proto.ColumnType_BOOL,
        },
        {
            Name: "space",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "task_count",
            Type: proto.ColumnType_STRING,
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
