---
organization: zigbang
category: ["saas"]
icon_url: "/images/plugins/zigbang/clickup.svg"
brand_color: "#7B68EE"
display_name: "ClickUp"
short_name: "clickup"
description: "Steampipe plugin for querying ClickUp Tasks, Lists and other resources."
og_description: Query ClickUp with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/zigbang/clickup-social-graphic.png"
---

# ClickUp + Turbot Steampipe

[ClickUp](https://clickup.com//) is a SaaS specialising in Project Management similar to Jira or Asana.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

## Documentation

- [Table definitions / examples](https://hub.steampipe.io/plugins/zigbang/clickup/tables)

## Getting Started

### Installation

```shell
steampipe plugin install zigbang/clickup
```

### Prerequisites

- ClickUp Account
- [ClickUp API Token](https://clickup.com/api/developer-portal/authentication#personal-token) 

### Configuration

The preferred option is to use Environment Variables for configuration.

However, you can configure in the `~/.steampipe/config/clickup.spc` (this will take precedence).

Environment Variables:
- `CLICKUP_TOKEN` for the API token (ex: `pk_t348t9v3UYFG535ti`)

Configuration File:

```hcl
connection "clickup" {
  plugin  = "zigbang/clickup"
  api_key = "pk_t348t9v3UYFG535ti"
}
```

### Testing

A quick test can be performed from your terminal with:

```shell
steampipe query "select * from clickup_task"
```
