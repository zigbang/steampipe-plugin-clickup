package clickup

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			"clickup_folder":          tableClickupFolder(ctx),
			"clickup_folderless_list": tableClickupFolderlessList(ctx),
			"clickup_list":            tableClickupList(ctx),
			"clickup_space":           tableClickupSpace(ctx),
			"clickup_task":            tableClickupTask(ctx),
			"clickup_team":            tableClickupTeam(ctx),
			"clickup_team_member":     tableClickupTeamMember(ctx),
		},
	}
	return p
}
