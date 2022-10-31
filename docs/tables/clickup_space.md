# Table: clickup_space

Obtain information on spaces within your ClickUp environment, you must specify a `team_id` in the WHERE or JOIN clause.

## Examples

### List all spaces for a specific team

```sql
select
  id,
  name,
  private,
  archived
from
  clickup_space
where
  team_id = 'myteam';
```
