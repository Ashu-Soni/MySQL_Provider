---
layout: "sql"
page_title: "MySQL: sql_role"
sidebar_current: "docs-sql-resource-role"
description: |-
  Creates and manages a role  on a MySQL server.
---

# mysql\_role

The ``sql_role`` resource creates and manages a user on a MySQL
server.

~> **Note:** MySQL introduced roles in version 8. They do not work on MySQL 5 and lower.


## Example Usage

```hcl
resource "sql_role" "role1" {
  name = "admin"
}
```
## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the role.

## Attributes Reference

No further attributes are exported.

