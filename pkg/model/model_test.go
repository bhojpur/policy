package model

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
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bhojpur/policy/pkg/config"
)

var (
	basicExample = filepath.Join("../..", "examples", "basic_model.conf")
	basicConfig  = &MockConfig{
		data: map[string]string{
			"request_definition::r": "sub, obj, act",
			"policy_definition::p":  "sub, obj, act",
			"policy_effect::e":      "some(where (p.eft == allow))",
			"matchers::m":           "r.sub == p.sub && r.obj == p.obj && r.act == p.act",
		},
	}
)

type MockConfig struct {
	data map[string]string
	config.ConfigInterface
}

func (mc *MockConfig) String(key string) string {
	return mc.data[key]
}

func TestNewModel(t *testing.T) {
	m := NewModel()
	if m == nil {
		t.Error("new model should not be nil")
	}
}

func TestNewModelFromFile(t *testing.T) {
	m, err := NewModelFromFile(basicExample)
	if err != nil {
		t.Errorf("model failed to load from file: %s", err)
	}
	if m == nil {
		t.Error("model should not be nil")
	}
}

func TestNewModelFromString(t *testing.T) {
	modelBytes, _ := ioutil.ReadFile(basicExample)
	modelString := string(modelBytes)
	m, err := NewModelFromString(modelString)
	if err != nil {
		t.Errorf("model faild to load from string: %s", err)
	}
	if m == nil {
		t.Error("model should not be nil")
	}
}

func TestLoadModelFromConfig(t *testing.T) {
	m := NewModel()
	err := m.loadModelFromConfig(basicConfig)
	if err != nil {
		t.Error("basic config should not return an error")
	}
	m = NewModel()
	err = m.loadModelFromConfig(&MockConfig{})
	if err == nil {
		t.Error("empty config should return error")
	} else {
		// check for missing sections in message
		for _, rs := range requiredSections {
			if !strings.Contains(err.Error(), sectionNameMap[rs]) {
				t.Errorf("section name: %s should be in message", sectionNameMap[rs])
			}
		}
	}
}

func TestHasSection(t *testing.T) {
	m := NewModel()
	_ = m.loadModelFromConfig(basicConfig)
	for _, sec := range requiredSections {
		if !m.hasSection(sec) {
			t.Errorf("%s section was expected in model", sec)
		}
	}
	m = NewModel()
	_ = m.loadModelFromConfig(&MockConfig{})
	for _, sec := range requiredSections {
		if m.hasSection(sec) {
			t.Errorf("%s section was not expected in model", sec)
		}
	}
}

func TestModel_AddDef(t *testing.T) {
	m := NewModel()
	s := "r"
	v := "sub, obj, act"
	ok := m.AddDef(s, s, v)
	if !ok {
		t.Errorf("non empty assertion should be added")
	}
	ok = m.AddDef(s, s, "")
	if ok {
		t.Errorf("empty assertion value should not be added")
	}
}

func TestModelToTest(t *testing.T) {
	testModelToText(t, "r.sub == p.sub && r.obj == p.obj && r_func(r.act, p.act) && testr_func(r.act, p.act)", "r_sub == p_sub && r_obj == p_obj && r_func(r_act, p_act) && testr_func(r_act, p_act)")
	testModelToText(t, "r.sub == p.sub && r.obj == p.obj && p_func(r.act, p.act) && testp_func(r.act, p.act)", "r_sub == p_sub && r_obj == p_obj && p_func(r_act, p_act) && testp_func(r_act, p_act)")
}

func testModelToText(t *testing.T, mData, mExpected string) {
	m := NewModel()
	data := map[string]string{
		"r": "sub, obj, act",
		"p": "sub, obj, act",
		"e": "some(where (p.eft == allow))",
		"m": mData,
	}
	expected := map[string]string{
		"r": "sub, obj, act",
		"p": "sub, obj, act",
		"e": "some(where (p_eft == allow))",
		"m": mExpected,
	}
	addData := func(ptype string) {
		m.AddDef(ptype, ptype, data[ptype])
	}
	for ptype := range data {
		addData(ptype)
	}
	newM := NewModel()
	print(m.ToText())
	_ = newM.LoadModelFromText(m.ToText())
	for ptype := range data {
		if newM[ptype][ptype].Value != expected[ptype] {
			t.Errorf("\"%s\" assertion value changed, current value: %s, it should be: %s", ptype, newM[ptype][ptype].Value, expected[ptype])
		}
	}
}
