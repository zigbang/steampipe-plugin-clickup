# Table: clickup_list

Obtain information on lists within your ClickUp environment, you must specify a `folder_id` in the WHERE or JOIN clause.

## Examples

### List all lists for a specific folder

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
  clickup_list
where
  folder_id = 'myfolder';
```
