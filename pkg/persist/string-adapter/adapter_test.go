package string_adapter

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"testing"

	plcsvr "github.com/bhojpur/policy/pkg/engine"
	"github.com/bhojpur/policy/pkg/model"
)

func Test_KeyMatchRbac(t *testing.T) {
	conf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub)  && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
	line := `
p, alice, /alice_data/*, (GET)|(POST)
p, alice, /alice_data/resource1, POST
p, data_group_admin, /admin/*, POST
p, data_group_admin, /bob_data/*, POST
g, alice, data_group_admin
`
	sa := NewAdapter(line)
	md := model.NewModel()
	err := md.LoadModelFromText(conf)
	if err != nil {
		t.Error(err)
		return
	}
	e, _ := plcsvr.NewEnforcer(md, sa)
	sub := "alice"
	obj := "/alice_data/login"
	act := "POST"
	if res, _ := e.Enforce(sub, obj, act); res != true {
		t.Error("**error**")
	}
}

func Test_StringRbac(t *testing.T) {
	conf := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	line := `
p, alice, data1, read
p, data_group_admin, data3, read
p, data_group_admin, data3, write
g, alice, data_group_admin
`
	sa := NewAdapter(line)
	md := model.NewModel()
	err := md.LoadModelFromText(conf)
	if err != nil {
		t.Error(err)
		return
	}
	e, _ := plcsvr.NewEnforcer(md, sa)
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	if res, _ := e.Enforce(sub, obj, act); res != true {
		t.Error("**error**")
	}
}
