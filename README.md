# ClickUp Plugin for Steampipe

## Installing and Testing the Plugin

To install the plugin, simple run the following command.

```
% make local
go build -o  ~/.steampipe/plugins/local/clickup/clickup.plugin *.go
```

Check your local plugin using the following command.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/clickup                                    | local   |             |
+--------------------------------------------------+---------+-------------+
```

Copy the sample `clickup.spc` file to `~/.steampipe/config` folder and change the name of the `plugin` from `clickup` to `local/clickup`.

```
% cat ~/.steampipe/config/clickup.spc
connection "clickup" {
    plugin = "local/clickup"

    token = "YOUR_API_TOKEN_HERE"
}
```

Check and see if you have a valid connection.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/clickup                                    | local   | clickup     |
+--------------------------------------------------+---------+-------------+
```

Let's test the plugin.

```
% % steampipe query "select count(*) from clickup_team" --timing
+-------+
| count |
+-------+
| 2     |
+-------+

Time: 1.73914125s
```

That's it.
