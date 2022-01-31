# Bhojpur Policy - String Adapter

The string adapter for [Bhojpur Policy](https://github.com/bhojpur/policy)

## Installation

go get github.com/bhojpur/policy


### Simple Example


```go
package main

import (
	"fmt"

	sadap "github.com/bhojpur/policy/pkg/persist/string-adapter"

	plcsvr "github.com/bhojpur/policy/pkg/engine"
	"github.com/bhojpur/policy/pkg/model"
)

func main() {

	modelText := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act`

	m := model.Model{}

	m.LoadModelFromText(modelText)

	line := `
p, alice, data1, read
p, bob, data2, write
p, data_group_admin, data_group, write

g, alice, data_group_admin
g2, data1, data_group
g2, data2, data_group
`
	sa := sadap.NewAdapter(line)

	// Initialize a Gorm adapter and use it in a Bhojpur Policy enforcer:
	// The adapter will use the MySQL database named "bhojpur".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing ORM instance with ormadapter.NewAdapterByDB(ormInstance)
	//a, _ := ormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	//	e, _ := plcsvr.NewEnforcer("../../examples/rbac_model.conf", a)
	e, _ := plcsvr.NewEnforcer(m, sa)

	// Or, you can use an existing DB "abc" like this:
	// The adapter will use the table named "bhojpur_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := ormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	if res, _ := e.Enforce("alice", "data1", "read"); res {
		fmt.Println("permitted")
	} else {
		fmt.Println("rejected")
	}

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}

```