package clickup

import (
    "context"

    "github.com/turbot/steampipe-plugin-sdk/v4/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
    p := &plugin.Plugin{
        Name: "steampipe-plugin-clickup",
        ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
            NewInstance: ConfigInstance,
            Schema:      ConfigSchema,
        },
        DefaultTransform: transform.FromGo(),
        DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
            TotalMaxConcurrency: 10,
        },
        TableMap: map[string]*plugin.Table{
            "clickup_folder":          tableClickupFolder(),
            "clickup_folderless_list": tableClickupFolderlessList(),
            "clickup_list":            tableClickupList(),
            "clickup_space":           tableClickupSpace(),
            "clickup_task":            tableClickupTask(),
            "clickup_team":            tableClickupTeam(),
            "clickup_team_member":     tableClickupTeamMember(),
        },
    }
    return p
}
