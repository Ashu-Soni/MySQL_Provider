provider "sql" {
  endpoint = "localhost:3306"
  username = "root"
  password = "AshuTosh07"
  
}


resource "sql_database" "app" {
  name = "ashutosh"
}

resource "sql_user" "user" {
  name = "darsh"
}

resource "sql_grant" "grant" {
  name = "darsh"
  dbname = "abc"
}

resource "sql_database" "bd2" {
  name = "abc"
}


