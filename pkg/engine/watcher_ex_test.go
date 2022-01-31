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

	"github.com/bhojpur/policy/pkg/model"
)

type SampleWatcherEx struct {
	SampleWatcher
}

func (w SampleWatcherEx) UpdateForAddPolicy(sec, ptype string, params ...string) error {
	return nil
}
func (w SampleWatcherEx) UpdateForRemovePolicy(sec, ptype string, params ...string) error {
	return nil
}

func (w SampleWatcherEx) UpdateForRemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}

func (w SampleWatcherEx) UpdateForSavePolicy(model model.Model) error {
	return nil
}

func (w SampleWatcherEx) UpdateForAddPolicies(sec string, ptype string, rules ...[]string) error {
	return nil
}

func (w SampleWatcherEx) UpdateForRemovePolicies(sec string, ptype string, rules ...[]string) error {
	return nil
}

func TestSetWatcherEx(t *testing.T) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_policy.csv")

	sampleWatcherEx := SampleWatcherEx{}
	err := e.SetWatcher(sampleWatcherEx)
	if err != nil {
		t.Fatal(err)
	}

	_ = e.SavePolicy()                              // calls watcherEx.UpdateForSavePolicy()
	_, _ = e.AddPolicy("admin", "data1", "read")    // calls watcherEx.UpdateForAddPolicy()
	_, _ = e.RemovePolicy("admin", "data1", "read") // calls watcherEx.UpdateForRemovePolicy()
	_, _ = e.RemoveFilteredPolicy(1, "data1")       // calls watcherEx.UpdateForRemoveFilteredPolicy()
	_, _ = e.RemovePolicy("admin", "data1", "read") // calls watcherEx.UpdateForRemovePolicy()
	_, _ = e.AddGroupingPolicy("g:admin", "data1")
	_, _ = e.RemoveGroupingPolicy("g:admin", "data1")
	_, _ = e.AddGroupingPolicy("g:admin", "data1")
	_, _ = e.RemoveFilteredGroupingPolicy(1, "data1")
	_, _ = e.AddPolicies([][]string{{"admin", "data1", "read"}, {"admin", "data2", "read"}})    // calls watcherEx.UpdateForAddPolicies()
	_, _ = e.RemovePolicies([][]string{{"admin", "data1", "read"}, {"admin", "data2", "read"}}) // calls watcherEx.UpdateForRemovePolicies()
}
