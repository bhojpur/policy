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

import "testing"

func testEnforceCache(t *testing.T, e *CachedEnforcer, sub string, obj interface{}, act string, res bool) {
	t.Helper()
	if myRes, _ := e.Enforce(sub, obj, act); myRes != res {
		t.Errorf("%s, %v, %s: %t, supposed to be %t", sub, obj, act, myRes, res)
	}
}

func TestCache(t *testing.T) {
	e, _ := NewCachedEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")
	// The cache is enabled by default for NewCachedEnforcer.

	testEnforceCache(t, e, "alice", "data1", "read", true)
	testEnforceCache(t, e, "alice", "data1", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", false)
	testEnforceCache(t, e, "alice", "data2", "write", false)

	// The cache is enabled, calling RemovePolicy, LoadPolicy or RemovePolicies will
	// also operate cached items.
	_, _ = e.RemovePolicy("alice", "data1", "read")

	testEnforceCache(t, e, "alice", "data1", "read", false)
	testEnforceCache(t, e, "alice", "data1", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", false)
	testEnforceCache(t, e, "alice", "data2", "write", false)

	e, _ = NewCachedEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_policy.csv")

	testEnforceCache(t, e, "alice", "data1", "read", true)
	testEnforceCache(t, e, "bob", "data2", "write", true)
	testEnforceCache(t, e, "alice", "data2", "read", true)
	testEnforceCache(t, e, "alice", "data2", "write", true)

	_, _ = e.RemovePolicies([][]string{
		{"alice", "data1", "read"},
		{"bob", "data2", "write"},
	})

	testEnforceCache(t, e, "alice", "data1", "read", false)
	testEnforceCache(t, e, "bob", "data2", "write", false)
	testEnforceCache(t, e, "alice", "data2", "read", true)
	testEnforceCache(t, e, "alice", "data2", "write", true)
}
