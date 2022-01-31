package fileadapter

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
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/bhojpur/policy/pkg/model"
	"github.com/bhojpur/policy/pkg/persist"
)

// AdapterMock is the file adapter for Bhojpur Policy.
// It can load policy from file or save policy to file.
type AdapterMock struct {
	filePath   string
	errorValue string
}

// NewAdapterMock is the constructor for AdapterMock.
func NewAdapterMock(filePath string) *AdapterMock {
	a := AdapterMock{}
	a.filePath = filePath
	return &a
}

// LoadPolicy loads all policy rules from the storage.
func (a *AdapterMock) LoadPolicy(model model.Model) error {
	err := a.loadPolicyFile(model, persist.LoadPolicyLine)
	return err
}

// SavePolicy saves all policy rules to the storage.
func (a *AdapterMock) SavePolicy(model model.Model) error {
	return nil
}

func (a *AdapterMock) loadPolicyFile(model model.Model, handler func(string, model.Model)) error {
	f, err := os.Open(a.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line, model)
		if err != nil {
			if err == io.EOF {
				return nil
			}
		}
	}
}

// SetMockErr sets string to be returned by of the mock during testing
func (a *AdapterMock) SetMockErr(errorToSet string) {
	a.errorValue = errorToSet
}

// GetMockErr returns a mock error or nil
func (a *AdapterMock) GetMockErr() error {
	var returnError error
	if a.errorValue != "" {
		returnError = errors.New(a.errorValue)
	}
	return returnError
}

// AddPolicy adds a policy rule to the storage.
func (a *AdapterMock) AddPolicy(sec string, ptype string, rule []string) error {
	return a.GetMockErr()
}

// AddPolicies removes policy rules from the storage.
func (a *AdapterMock) AddPolicies(sec string, ptype string, rules [][]string) error {
	return a.GetMockErr()
}

// RemovePolicy removes a policy rule from the storage.
func (a *AdapterMock) RemovePolicy(sec string, ptype string, rule []string) error {
	return a.GetMockErr()
}

// RemovePolicies removes policy rules from the storage.
func (a *AdapterMock) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return a.GetMockErr()
}

// UpdatePolicy removes a policy rule from the storage.
func (a *AdapterMock) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {
	return a.GetMockErr()
}

func (a *AdapterMock) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	return a.GetMockErr()
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *AdapterMock) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return a.GetMockErr()
}
