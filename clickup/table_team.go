package clickup

import (
    "context"
    "fmt"

    "github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableClickupTeam() *plugin.Table {
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
    client, err := connect(ctx, d)
    if err != nil {
        return nil, fmt.Errorf("unable to establish a connection: %v", err)
    }

    teams, _, err := client.Teams.GetTeams(ctx)
    if err != nil {
        return nil, fmt.Errorf("unable to obtain teams: %v", err)
    }

    for _, t := range teams {
        d.StreamListItem(ctx, t)
    }

    return nil, nil
}
