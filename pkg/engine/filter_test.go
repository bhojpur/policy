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
	"github.com/bhojpur/policy/pkg/util"
)

func TestInitFilteredAdapter(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_domains_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)

	// policy should not be loaded yet
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, false)
}

func TestLoadFilteredPolicy(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_domains_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)
	if err := e.LoadPolicy(); err != nil {
		t.Errorf("unexpected error in LoadPolicy: %v", err)
	}

	// validate initial conditions
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, true)
	testHasPolicy(t, e, []string{"admin", "domain2", "data2", "read"}, true)

	if err := e.LoadFilteredPolicy(&fileadapter.Filter{
		P: []string{"", "domain1"},
		G: []string{"", "", "domain1"},
	}); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}
	if !e.IsFiltered() {
		t.Errorf("adapter did not set the filtered flag correctly")
	}

	// only policies for domain1 should be loaded
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, true)
	testHasPolicy(t, e, []string{"admin", "domain2", "data2", "read"}, false)

	if err := e.SavePolicy(); err == nil {
		t.Errorf("enforcer did not prevent saving filtered policy")
	}
	if err := e.GetAdapter().SavePolicy(e.GetModel()); err == nil {
		t.Errorf("adapter did not prevent saving filtered policy")
	}
}

func TestLoadMoreTypeFilteredPolicy(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_pattern_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_pattern_model.conf", adapter)
	if err := e.LoadPolicy(); err != nil {
		t.Errorf("unexpected error in LoadPolicy: %v", err)
	}
	e.AddNamedMatchingFunc("g2", "matching func", util.KeyMatch2)
	_ = e.BuildRoleLinks()

	testEnforce(t, e, "alice", "/book/1", "GET", true)

	// validate initial conditions
	testHasPolicy(t, e, []string{"book_admin", "book_group", "GET"}, true)
	testHasPolicy(t, e, []string{"pen_admin", "pen_group", "GET"}, true)

	if err := e.LoadFilteredPolicy(&fileadapter.Filter{
		P:  []string{"book_admin"},
		G:  []string{"alice"},
		G2: []string{"", "book_group"},
	}); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}
	if !e.IsFiltered() {
		t.Errorf("adapter did not set the filtered flag correctly")
	}

	testHasPolicy(t, e, []string{"alice", "/pen/1", "GET"}, false)
	testHasPolicy(t, e, []string{"alice", "/pen2/1", "GET"}, false)
	testHasPolicy(t, e, []string{"pen_admin", "pen_group", "GET"}, false)
	testHasGroupingPolicy(t, e, []string{"alice", "book_admin"}, true)
	testHasGroupingPolicy(t, e, []string{"bob", "pen_admin"}, false)
	testHasGroupingPolicy(t, e, []string{"cathy", "pen_admin"}, false)
	testHasGroupingPolicy(t, e, []string{"cathy", "/book/1/2/3/4/5"}, false)

	testEnforce(t, e, "alice", "/book/1", "GET", true)
	testEnforce(t, e, "alice", "/pen/1", "GET", false)
}

func TestAppendFilteredPolicy(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_domains_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)
	if err := e.LoadPolicy(); err != nil {
		t.Errorf("unexpected error in LoadPolicy: %v", err)
	}

	// validate initial conditions
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, true)
	testHasPolicy(t, e, []string{"admin", "domain2", "data2", "read"}, true)

	if err := e.LoadFilteredPolicy(&fileadapter.Filter{
		P: []string{"", "domain1"},
		G: []string{"", "", "domain1"},
	}); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}
	if !e.IsFiltered() {
		t.Errorf("adapter did not set the filtered flag correctly")
	}

	// only policies for domain1 should be loaded
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, true)
	testHasPolicy(t, e, []string{"admin", "domain2", "data2", "read"}, false)

	//disable clear policy and load second domain
	if err := e.LoadIncrementalFilteredPolicy(&fileadapter.Filter{
		P: []string{"", "domain2"},
		G: []string{"", "", "domain2"},
	}); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}

	//both domain policies should be loaded
	testHasPolicy(t, e, []string{"admin", "domain1", "data1", "read"}, true)
	testHasPolicy(t, e, []string{"admin", "domain2", "data2", "read"}, true)
}

func TestFilteredPolicyInvalidFilter(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_domains_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)

	if err := e.LoadFilteredPolicy([]string{"", "domain1"}); err == nil {
		t.Errorf("expected error in LoadFilteredPolicy, but got nil")
	}
}

func TestFilteredPolicyEmptyFilter(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/rbac_with_domains_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)

	if err := e.LoadFilteredPolicy(nil); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}
	if e.IsFiltered() {
		t.Errorf("adapter did not reset the filtered flag correctly")
	}
	if err := e.SavePolicy(); err != nil {
		t.Errorf("unexpected error in SavePolicy: %v", err)
	}
}

func TestUnsupportedFilteredPolicy(t *testing.T) {
	e, _ := NewEnforcer("../../examples/rbac_with_domains_model.conf", "../../examples/rbac_with_domains_policy.csv")

	err := e.LoadFilteredPolicy(&fileadapter.Filter{
		P: []string{"", "domain1"},
		G: []string{"", "", "domain1"},
	})
	if err == nil {
		t.Errorf("encorcer should have reported incompatibility error")
	}
}

func TestFilteredAdapterEmptyFilepath(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)

	if err := e.LoadFilteredPolicy(nil); err != nil {
		t.Errorf("unexpected error in LoadFilteredPolicy: %v", err)
	}
}

func TestFilteredAdapterInvalidFilepath(t *testing.T) {
	e, _ := NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("../../examples/does_not_exist_policy.csv")
	_ = e.InitWithAdapter("../../examples/rbac_with_domains_model.conf", adapter)

	if err := e.LoadFilteredPolicy(nil); err == nil {
		t.Errorf("expected error in LoadFilteredPolicy, but got nil")
	}
}
