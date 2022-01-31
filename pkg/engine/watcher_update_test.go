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
)

type SampleWatcherUpdatable struct {
	SampleWatcher
}

func (w SampleWatcherUpdatable) UpdateForUpdatePolicy(params ...string) error {
	return nil
}

func TestSetWatcherUpdatable(t *testing.T) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_policy.csv")

	sampleWatcherEx := SampleWatcherUpdatable{}
	err := e.SetWatcher(sampleWatcherEx)
	if err != nil {
		t.Fatal(err)
	}

	_ = e.SavePolicy()                                                                            // calls watcherEx.UpdateForSavePolicy()
	_, _ = e.UpdatePolicy([]string{"admin", "data1", "read"}, []string{"admin", "data2", "read"}) // calls watcherEx.UpdateForUpdatePolicy()

}
