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
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func TestBhojpurJsGetPermissionForUser(t *testing.T) {
	e, err := NewSyncedEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_with_hierarchy_policy.csv")
	if err != nil {
		panic(err)
	}
	receivedString, err := BhojpurJsGetPermissionForUser(e, "alice") // make sure BhojpurJsGetPermissionForUser can be used with a SyncedEnforcer.
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	received := map[string]interface{}{}
	err = json.Unmarshal([]byte(receivedString), &received)
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	expectedModel, err := ioutil.ReadFile("../../examples/rbac_model.conf")
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	expectedModelStr := regexp.MustCompile("\n+").ReplaceAllString(string(expectedModel), "\n")
	if strings.TrimSpace(received["m"].(string)) != expectedModelStr {
		t.Errorf("%s supposed to be %s", strings.TrimSpace(received["m"].(string)), expectedModelStr)
	}

	expectedPolicies, err := ioutil.ReadFile("../../examples/rbac_with_hierarchy_policy.csv")
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
	expectedPoliciesItem := regexp.MustCompile(",|\n").Split(string(expectedPolicies), -1)
	i := 0
	for _, sArr := range received["p"].([]interface{}) {
		for _, s := range sArr.([]interface{}) {
			if strings.TrimSpace(s.(string)) != strings.TrimSpace(expectedPoliciesItem[i]) {
				t.Errorf("%s supposed to be %s", strings.TrimSpace(s.(string)), strings.TrimSpace(expectedPoliciesItem[i]))
			}
			i++
		}
	}
}
