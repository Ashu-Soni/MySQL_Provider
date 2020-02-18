---
layout: "sql"
page_title: "MySQL: sql_database"
sidebar_current: "docs-sql-resource-database"
description: |-
  Creates and manages a database on a MySQL server.
---

# sql\_database

The ``sql_database`` resource creates and manages a database on a MySQL
server.

## Example Usage

```hcl
resource "sql_database" "first_database" {
  name = "my_first_database"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the database. This must be unique within
  a given MySQL server and may or may not be case-sensitive depending on
  the operating system on which the MySQL server is running.

## Attributes Reference

The following attributes are exported:

* `name` - The name of the database.

## Import

Databases can be imported using their name, e.g.

```
$ terraform import sql_database.example my-example-database
```
