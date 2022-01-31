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

import (
	"encoding/csv"
	"strings"

	"github.com/bhojpur/policy/pkg/model"
)

// LoadPolicyLine loads a text line as a policy rule to model.
func LoadPolicyLine(line string, m model.Model) {
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}

	r := csv.NewReader(strings.NewReader(line))
	r.Comma = ','
	r.Comment = '#'
	r.TrimLeadingSpace = true

	tokens, err := r.Read()
	if err != nil {
		return
	}

	LoadPolicyArray(tokens, m)
}

// LoadPolicyArray loads a policy rule to model.
func LoadPolicyArray(rule []string, m model.Model) {
	key := rule[0]
	sec := key[:1]
	m[sec][key].Policy = append(m[sec][key].Policy, rule[1:])
	m[sec][key].PolicyMap[strings.Join(rule[1:], model.DefaultSep)] = len(m[sec][key].Policy) - 1
}

// Adapter is the interface for Bhojpur Policy adapters.
type Adapter interface {
	// LoadPolicy loads all policy rules from the storage.
	LoadPolicy(model model.Model) error
	// SavePolicy saves all policy rules to the storage.
	SavePolicy(model model.Model) error

	// AddPolicy adds a policy rule to the storage.
	// This is part of the Auto-Save feature.
	AddPolicy(sec string, ptype string, rule []string) error
	// RemovePolicy removes a policy rule from the storage.
	// This is part of the Auto-Save feature.
	RemovePolicy(sec string, ptype string, rule []string) error
	// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
	// This is part of the Auto-Save feature.
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
}
