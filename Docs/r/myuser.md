---
layout: "sql"
page_title: "MySQL: sql_user"
sidebar_current: "docs-sql-resource-user"
description: |-
  Creates and manages a user on a MySQL server.
---

# sql\_user

The ``sql_user`` resource creates and manages a user on a MySQL
server.

## Example Usage

```hcl
resource "sql_user" "first_user" {
  name               = "john"

}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user.

## Attributes Reference

The following attributes are exported:

* `name` - The name of the user.

## Attributes Reference

No further attributes are exported.
