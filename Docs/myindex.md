---
layout: "sql"
page_title: "Provider: sql"
sidebar_current: "docs-sql-index"
description: |-
  A provider for MySQL Server.
---

# MySQL Provider

[MySQL](http://www.mysql.com) is a relational database server. The MySQL
provider exposes resources used to manage the configuration of resources
in a MySQL server.

Use the navigation to the left to read about the available resources.

## Example Usage

The following is a minimal example:

```hcl
# Configure the MySQL provider
provider "sql" {
  endpoint = "localhost:3306"
  username = "app-user"
  password = "app-password"
}

# Create a Database
resource "sql_database" "first_database" {
  name = "my_first_database"
}
```
## Argument Reference

The following arguments are supported:

* `endpoint` - (Required) The address of the MySQL server to use. Most often a "hostname:port" pair, but may also be an absolute path to a Unix socket when the host OS is Unix-compatible. Can also be sourced from the `MYSQL_ENDPOINT` environment variable.
* `username` - (Required) Username to use to authenticate with the server, can also be sourced from the `MYSQL_USERNAME` environment variable.
* `password` - (Optional) Password for the given user, if that user has a password, can also be sourced from the `MYSQL_PASSWORD` environment variable.

## Note

For detailed implementation of mysql terraform provider, refer [this link](https://github.com/terraform-providers/terraform-provider-mysql).
