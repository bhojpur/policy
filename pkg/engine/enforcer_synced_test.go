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
	"time"
)

func testEnforceSync(t *testing.T, e *SyncedEnforcer, sub string, obj interface{}, act string, res bool) {
	t.Helper()
	if myRes, _ := e.Enforce(sub, obj, act); myRes != res {
		t.Errorf("%s, %v, %s: %t, supposed to be %t", sub, obj, act, myRes, res)
	}
}

func TestSync(t *testing.T) {
	e, _ := NewSyncedEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")
	// Start reloading the policy every 200 ms.
	e.StartAutoLoadPolicy(time.Millisecond * 200)

	testEnforceSync(t, e, "alice", "data1", "read", true)
	testEnforceSync(t, e, "alice", "data1", "write", false)
	testEnforceSync(t, e, "alice", "data2", "read", false)
	testEnforceSync(t, e, "alice", "data2", "write", false)
	testEnforceSync(t, e, "bob", "data1", "read", false)
	testEnforceSync(t, e, "bob", "data1", "write", false)
	testEnforceSync(t, e, "bob", "data2", "read", false)
	testEnforceSync(t, e, "bob", "data2", "write", true)

	// Stop the reloading policy periodically.
	e.StopAutoLoadPolicy()
}

func TestStopAutoLoadPolicy(t *testing.T) {
	e, _ := NewSyncedEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv")
	e.StartAutoLoadPolicy(5 * time.Millisecond)
	if !e.IsAutoLoadingRunning() {
		t.Error("auto load is not running")
	}
	e.StopAutoLoadPolicy()
	// Need a moment, to exit goroutine
	time.Sleep(10 * time.Millisecond)
	if e.IsAutoLoadingRunning() {
		t.Error("auto load is still running")
	}
}
