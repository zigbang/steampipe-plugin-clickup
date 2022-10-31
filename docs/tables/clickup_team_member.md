# Table: clickup_team_member

Obtain information on team members from teams within your ClickUp environment.

## Examples

### List all team members

```sql
select
  team_id,
  team_name,
  user_id,
  username,
  email,
  color,
  profile_picture,
  initials,
  last_active
from
  clickup_team_member;
```
