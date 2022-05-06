package clickup

import (
	"context"

	"github.com/raksul/go-clickup/clickup"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableClickupTeamMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "clickup_team_member",
		List: &plugin.ListConfig{
			Hydrate: listTeamMembers,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.ID")},
			{Name: "team_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.Name")},
			{Name: "user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID")},
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Username")},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Email")},
			{Name: "color", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Color")},
			{Name: "profile_picture", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ProfilePicture")},
			{Name: "initials", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Initials")},
			{Name: "role", Type: proto.ColumnType_INT, Transform: transform.FromField("User.Role")},
			{Name: "custom_role", Type: proto.ColumnType_JSON, Transform: transform.FromField("User.CustomRole")},
			{Name: "last_active", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.LastActive")},
			{Name: "date_joined", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.DateJoined")},
			{Name: "date_invited", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.DateInvited")},
		},
	}
}

func listTeamMembers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	client, _ := connect(ctx, d)
	teams, _, _ := client.Teams.GetTeams(ctx)

	for _, t := range teams {
		for _, m := range t.Members {
			item := TeamMember{
				Team: t, User: m.User,
			}
			d.StreamListItem(ctx, item)
		}
	}

	return nil, nil
}

type TeamMember struct {
	Team clickup.Team
	User clickup.TeamUser
}
