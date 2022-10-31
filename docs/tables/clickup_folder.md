# Table: clickup_folder

Obtain information on folders within your ClickUp environment, you must specify a `space_id` in the WHERE or JOIN clause.

## Examples

### List all folders for a specific space

```sql
select
  id,
  name,
  task_count
from
  clickup_folder
where
  space_id = 'myspace';
```
