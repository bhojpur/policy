# Bhojpur Policy - ORM Adapter

With this library, the [Bhojpur Policy](https://github.com/bhojpur/policy) can load policy from ORM supported databases or save policy to it.

Based on the ORM Drivers Support, rhe current supported databases are:

- MySQL: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- MyMySQL: [github.com/ziutek/mymysql/godrv](https://github.com/ziutek/mymysql/godrv)
- PostgreSQL: [github.com/lib/pq](https://github.com/lib/pq)
- TiDB: [github.com/pingcap/tidb](https://github.com/pingcap/tidb)
- SQLite: [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
- MS-SQL: [github.com/denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb)
- Oracle: [github.com/mattn/go-oci8](https://github.com/mattn/go-oci8) (experiment)

## Installation

    go get github.com/bhojpur/policy

## Simple MySQL Example

```go
package main

import (
	plcsvr "github.com/bhojpur/policy/pkg/engine"
	_ "github.com/go-sql-driver/mysql"

	ormadapter "github.com/bhojpur/policy/pkg/persist/orm-adapter"
)

func main() {
	// Initialize an ORM Adapter and use it in a Bhojpur Policy enforcer:
	// The adapter will use the MySQL database named "bhojpur".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := ormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source. 

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "bhojpur_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := ormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	e, _ := plcsvr.NewEnforcer("../../examples/rbac_model.conf", a)
	
	// Load the policy from DB.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Simple Postgres Example

```go
package main

import (
	plcsvr "github.com/bhojpur/policy/pkg/engine"
	_ "github.com/lib/pq"

	ormadapter "github.com/bhojpur/policy/pkg/persist/orm-adapter"
)

func main() {
	// Initialize an ORM Adapter and use it in a Bhojpur Policy enforcer:
	// The adapter will use the PostgreSQL database named "bhojpur".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := ormadapter.NewAdapter("postgres", "user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "bhojpur_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := ormadapter.NewAdapter("postgres", "dbname=abc user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable", true)

	e, _ := plcsvr.NewEnforcer("../../examples/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Getting Help

- [Bhojpur Policy](https://github.com/bhojpur/policy)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.