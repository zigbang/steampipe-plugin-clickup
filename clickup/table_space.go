package clickup

import (
    "context"
    "fmt"

    "github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableClickupSpace() *plugin.Table {
    return &plugin.Table{
        Name: "clickup_space",
        List: &plugin.ListConfig{
            KeyColumns: plugin.SingleColumn("team_id"),
            Hydrate:    listSpaces,
        },
        Columns: spaceColumns(),
    }
}

func listSpaces(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    teamId := d.KeyColumnQuals["team_id"].GetStringValue()

    client, err := connect(ctx, d)
    if err != nil {
        return nil, fmt.Errorf("unable to establish a connection: %v", err)
    }

    spaces, _, err := client.Spaces.GetSpaces(ctx, teamId)
    if err != nil {
        return nil, fmt.Errorf("unable to obtain spaces for team id '%s': %v", teamId, err)
    }

    for _, t := range spaces {
        d.StreamListItem(ctx, t)
    }

    return nil, nil
}

func spaceColumns() []*plugin.Column {
    return []*plugin.Column{
        {
            Name:      "team_id",
            Type:      proto.ColumnType_STRING,
            Transform: transform.FromQual("team_id"),
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
            Name: "private",
            Type: proto.ColumnType_BOOL,
        },
        {
            Name: "statuses",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "multiple_assignees",
            Type: proto.ColumnType_BOOL,
        },
        {
            Name: "features",
            Type: proto.ColumnType_JSON,
        },
        {
            Name: "archived",
            Type: proto.ColumnType_BOOL,
        },
    }
}
