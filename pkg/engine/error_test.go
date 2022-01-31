package engine

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

	fileadapter "github.com/bhojpur/policy/pkg/persist/file-adapter"
)

func TestPathError(t *testing.T) {
	_, err := NewEnforcer("hope_this_path_wont_exist", "")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}
}

func TestEnforcerParamError(t *testing.T) {
	_, err := NewEnforcer(1, 2, 3)
	if err == nil {
		t.Errorf("Should not be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}

	_, err2 := NewEnforcer(1, "2")
	if err2 == nil {
		t.Errorf("Should not be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err2.Error())
	}
}

func TestModelError(t *testing.T) {
	_, err := NewEnforcer("../../examples/error/error_model.conf", "../../examples/error/error_policy.csv")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}
}

//func TestPolicyError(t *testing.T) {
//	_, err := NewEnforcer("../../examples/basic_model.conf", "../../examples/error/error_policy.csv")
//	if err == nil {
//		t.Errorf("Should be error here.")
//	} else {
//		t.Log("Test on error: ")
//		t.Log(err.Error())
//	}
//}

func TestEnforceError(t *testing.T) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")
	_, err := e.Enforce("wrong", "wrong")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}

	e, _ = NewEnforcer("../../examples/abac_rule_model.conf")
	_, err = e.Enforce("wang", "wang", "wang")
	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}
}

func TestNoError(t *testing.T) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")

	err := e.LoadModel()
	if err != nil {
		t.Errorf("Should be no error here.")
		t.Log("Unexpected error: ")
		t.Log(err.Error())
	}

	err = e.LoadPolicy()
	if err != nil {
		t.Errorf("Should be no error here.")
		t.Log("Unexpected error: ")
		t.Log(err.Error())
	}

	err = e.SavePolicy()
	if err != nil {
		t.Errorf("Should be no error here.")
		t.Log("Unexpected error: ")
		t.Log(err.Error())
	}
}

func TestModelNoError(t *testing.T) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")

	e.modelPath = "hope_this_path_wont_exist"
	err := e.LoadModel()

	if err == nil {
		t.Errorf("Should be error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}
}

func TestMockAdapterErrors(t *testing.T) {
	adapter := fileadapter.NewAdapterMock("../../examples/rbac_with_domains_policy.csv")
	adapter.SetMockErr("mock error")

	e, _ := NewEnforcer("../../examples/rbac_with_domains_model.conf", adapter)

	added, err := e.AddPolicy("admin", "domain3", "data1", "read")
	if added {
		t.Errorf("added should be false")
	}

	if err == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}

	rules := [][]string{
		{"admin", "domain4", "data1", "read"},
	}
	added, err = e.AddPolicies(rules)
	if added {
		t.Errorf("added should be false")
	}

	if err == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}

	removed, err2 := e.RemoveFilteredPolicy(1, "domain1", "data1")
	if removed {
		t.Errorf("removed should be false")
	}

	if err2 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err2.Error())
	}

	removed, err3 := e.RemovePolicy("admin", "domain2", "data2", "read")
	if removed {
		t.Errorf("removed should be false")
	}

	if err3 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err3.Error())
	}

	rules = [][]string{
		{"admin", "domain1", "data1", "read"},
	}
	removed, err = e.RemovePolicies(rules)
	if removed {
		t.Errorf("removed should be false")
	}

	if err == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err.Error())
	}

	added, err4 := e.AddGroupingPolicy("bob", "admin2")
	if added {
		t.Errorf("added should be false")
	}

	if err4 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err4.Error())
	}

	added, err5 := e.AddNamedGroupingPolicy("g", []string{"eve", "admin2", "domain1"})
	if added {
		t.Errorf("added should be false")
	}

	if err5 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err5.Error())
	}

	added, err6 := e.AddNamedPolicy("p", []string{"admin2", "domain2", "data2", "write"})
	if added {
		t.Errorf("added should be false")
	}

	if err6 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err6.Error())
	}

	removed, err7 := e.RemoveGroupingPolicy("bob", "admin2")
	if removed {
		t.Errorf("removed should be false")
	}

	if err7 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err7.Error())
	}

	removed, err8 := e.RemoveFilteredGroupingPolicy(0, "bob")
	if removed {
		t.Errorf("removed should be false")
	}

	if err8 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err8.Error())
	}

	removed, err9 := e.RemoveNamedGroupingPolicy("g", []string{"alice", "admin", "domain1"})
	if removed {
		t.Errorf("removed should be false")
	}

	if err9 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err9.Error())
	}

	removed, err10 := e.RemoveFilteredNamedGroupingPolicy("g", 0, "eve")
	if removed {
		t.Errorf("removed should be false")
	}

	if err10 == nil {
		t.Errorf("Should be an error here.")
	} else {
		t.Log("Test on error: ")
		t.Log(err10.Error())
	}
}
