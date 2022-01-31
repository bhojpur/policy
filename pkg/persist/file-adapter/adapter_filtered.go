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
	"os"
	"strings"

	"github.com/bhojpur/policy/pkg/model"
	"github.com/bhojpur/policy/pkg/persist"
)

// FilteredAdapter is the filtered file adapter for Bhojpur Policy. It can load policy
// from file or save policy to file and supports loading of filtered policies.
type FilteredAdapter struct {
	*Adapter
	filtered bool
}

// Filter defines the filtering rules for a FilteredAdapter's policy. Empty values
// are ignored, but all others must match the filter.
type Filter struct {
	P  []string
	G  []string
	G1 []string
	G2 []string
	G3 []string
	G4 []string
	G5 []string
}

// NewFilteredAdapter is the constructor for FilteredAdapter.
func NewFilteredAdapter(filePath string) *FilteredAdapter {
	a := FilteredAdapter{}
	a.filtered = true
	a.Adapter = NewAdapter(filePath)
	return &a
}

// LoadPolicy loads all policy rules from the storage.
func (a *FilteredAdapter) LoadPolicy(model model.Model) error {
	a.filtered = false
	return a.Adapter.LoadPolicy(model)
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *FilteredAdapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	if filter == nil {
		return a.LoadPolicy(model)
	}
	if a.filePath == "" {
		return errors.New("invalid file path, file path cannot be empty")
	}

	filterValue, ok := filter.(*Filter)
	if !ok {
		return errors.New("invalid filter type")
	}
	err := a.loadFilteredPolicyFile(model, filterValue, persist.LoadPolicyLine)
	if err == nil {
		a.filtered = true
	}
	return err
}

func (a *FilteredAdapter) loadFilteredPolicyFile(model model.Model, filter *Filter, handler func(string, model.Model)) error {
	f, err := os.Open(a.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if filterLine(line, filter) {
			continue
		}

		handler(line, model)
	}
	return scanner.Err()
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *FilteredAdapter) IsFiltered() bool {
	return a.filtered
}

// SavePolicy saves all policy rules to the storage.
func (a *FilteredAdapter) SavePolicy(model model.Model) error {
	if a.filtered {
		return errors.New("cannot save a filtered policy")
	}
	return a.Adapter.SavePolicy(model)
}

func filterLine(line string, filter *Filter) bool {
	if filter == nil {
		return false
	}
	p := strings.Split(line, ",")
	if len(p) == 0 {
		return true
	}
	var filterSlice []string
	switch strings.TrimSpace(p[0]) {
	case "p":
		filterSlice = filter.P
	case "g":
		filterSlice = filter.G
	case "g1":
		filterSlice = filter.G1
	case "g2":
		filterSlice = filter.G2
	case "g3":
		filterSlice = filter.G3
	case "g4":
		filterSlice = filter.G4
	case "g5":
		filterSlice = filter.G5
	}
	return filterWords(p, filterSlice)
}

func filterWords(line []string, filter []string) bool {
	if len(line) < len(filter)+1 {
		return true
	}
	var skipLine bool
	for i, v := range filter {
		if len(v) > 0 && strings.TrimSpace(v) != strings.TrimSpace(line[i+1]) {
			skipLine = true
			break
		}
	}
	return skipLine
}
