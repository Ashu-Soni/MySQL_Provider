---
layout: "sql"
page_title: "MySQL: sql_grant"
sidebar_current: "docs-sql-resource-grant"
description: |-
  Creates and manages privileges given to a user on a MySQL server
---

# sql\_grant

The ``sql_grant`` resource creates and manages privileges given to
a user on a MySQL server.

## Granting Privileges to a User

```hcl
resource "sql_database" "first_database" {
  name = "my_first_database"
}

resource "sql_user" "first_user" {
  name        = "john"
}

resource "sql_grant" "jdoe" {
  name       = "${sql_user.first_user.user}"
  dbname 	 =  "${sql_database.first_database.name}"
}
```

## Granting Privileges to a Role

```hcl
resource "sql_role" "role1" {
  name = "admin"
}

resource "sql_grant" "developer" {
  name       = "${sql_user.first_user.name}"
  role       = "${sql_role.role1.name}"
  dbname     = "${sql_database.first_database.name}"
}
```


## Argument Reference

~> **Note:** MySQL removed the `REQUIRE` option from `GRANT` in version 8. `tls_option` is ignored in MySQL 8 and above.
~> **Note:** `sql_grant` will grant all the privileges to specific user and database.

The following arguments are supported:

* `name` - (Required) The name of the user. Conflicts with `role`.
* `role` - (Optional) The role to grant `privileges` to. Conflicts with `user` and `host`.
* `dbname` - (Required) The database to grant privileges on.

## Attributes Reference

No further attributes are exported.


