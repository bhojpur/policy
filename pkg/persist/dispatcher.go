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

// Dispatcher is the interface for Bhojpur Policy dispatcher
type Dispatcher interface {
	// AddPolicies adds policies rule to all instance.
	AddPolicies(sec string, ptype string, rules [][]string) error
	// RemovePolicies removes policies rule from all instance.
	RemovePolicies(sec string, ptype string, rules [][]string) error
	// RemoveFilteredPolicy removes policy rules that match the filter from all instance.
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
	// ClearPolicy clears all current policy in all instances
	ClearPolicy() error
	// UpdatePolicy updates policy rule from all instance.
	UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error
	// UpdatePolicies updates some policy rules from all instance
	UpdatePolicies(sec string, ptype string, oldrules, newRules [][]string) error
	// UpdateFilteredPolicies deletes old rules and adds new rules.
	UpdateFilteredPolicies(sec string, ptype string, oldRules [][]string, newRules [][]string) error
}
