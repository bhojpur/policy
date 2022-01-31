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
	"encoding/json"
	"testing"
)

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func TestBhojpurJsGetPermissionForUserOld(t *testing.T) {
	e, err := NewEnforcer("examples/rbac_model.conf", "examples/rbac_policy.csv")
	if err != nil {
		panic(err)
	}
	target_str, _ := BhojpurJsGetPermissionForUserOld(e, "alice")
	t.Log("GetPermissionForUser Alice", string(target_str))
	alice_target := make(map[string][]string)
	err = json.Unmarshal(target_str, &alice_target)
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	perm, ok := alice_target["read"]
	if !ok {
		t.Errorf("Test error: Alice doesn't have read permission")
	}
	if !contains(perm, "data1") {
		t.Errorf("Test error: Alice cannot read data1")
	}
	if !contains(perm, "data2") {
		t.Errorf("Test error: Alice cannot read data2")
	}
	perm, ok = alice_target["write"]
	if !ok {
		t.Errorf("Test error: Alice doesn't have write permission")
	}
	if contains(perm, "data1") {
		t.Errorf("Test error: Alice can write data1")
	}
	if !contains(perm, "data2") {
		t.Errorf("Test error: Alice cannot write data2")
	}

	target_str, _ = BhojpurJsGetPermissionForUserOld(e, "bob")
	t.Log("GetPermissionForUser Bob", string(target_str))
	bob_target := make(map[string][]string)
	err = json.Unmarshal(target_str, &bob_target)
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	_, ok = bob_target["read"]
	if ok {
		t.Errorf("Test error: Bob has read permission")
	}
	perm, ok = bob_target["write"]
	if !ok {
		t.Errorf("Test error: Bob doesn't have permission")
	}
	if !contains(perm, "data2") {
		t.Errorf("Test error: Bob cannot write data2")
	}
	if contains(perm, "data1") {
		t.Errorf("Test error: Bob can write data1")
	}
	if contains(perm, "data_not_exist") {
		t.Errorf("Test error: Bob can access a non-existing data")
	}

	_, ok = bob_target["rm_rf"]
	if ok {
		t.Errorf("Someone can have a non-existing action (rm -rf)")
	}
}
