# Table: clickup_folderless_list

Obtain information on folderless lists within your ClickUp environment, you must specify a `space_id` in the WHERE or JOIN clause.

## Examples

### List all folderless lists for a specific space

```sql
select
  id,
  name,
  task_count,
  status,
  priority,
  assignee,
  assignee_email,
  due_date
from
  clickup_folderless_list
where
  space_id = 'myspace';
```
