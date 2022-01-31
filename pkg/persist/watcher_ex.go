package persist

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

import "github.com/bhojpur/policy/pkg/model"

// WatcherEx is the strengthen for Bhojpur Policy watchers.
type WatcherEx interface {
	Watcher
	// UpdateForAddPolicy calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.AddPolicy()
	UpdateForAddPolicy(sec, ptype string, params ...string) error
	// UpdateForRemovePolicy calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.RemovePolicy()
	UpdateForRemovePolicy(sec, ptype string, params ...string) error
	// UpdateForRemoveFilteredPolicy calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.RemoveFilteredNamedGroupingPolicy()
	UpdateForRemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error
	// UpdateForSavePolicy calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.RemoveFilteredNamedGroupingPolicy()
	UpdateForSavePolicy(model model.Model) error
	// UpdateForAddPolicies calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.AddPolicies()
	UpdateForAddPolicies(sec string, ptype string, rules ...[]string) error
	// UpdateForRemovePolicies calls the update callback of other instances to synchronize their policy.
	// It is called after Enforcer.RemovePolicies()
	UpdateForRemovePolicies(sec string, ptype string, rules ...[]string) error
}
