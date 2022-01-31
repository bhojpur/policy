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
	"sync"

	"github.com/Knetic/govaluate"
	"github.com/bhojpur/policy/pkg/util"
)

// FunctionMap represents the collection of Function.
type FunctionMap struct {
	fns *sync.Map
}

// [string]govaluate.ExpressionFunction

// AddFunction adds an expression function.
func (fm *FunctionMap) AddFunction(name string, function govaluate.ExpressionFunction) {
	fm.fns.LoadOrStore(name, function)
}

// LoadFunctionMap loads an initial function map.
func LoadFunctionMap() FunctionMap {
	fm := &FunctionMap{}
	fm.fns = &sync.Map{}

	fm.AddFunction("keyMatch", util.KeyMatchFunc)
	fm.AddFunction("keyGet", util.KeyGetFunc)
	fm.AddFunction("keyMatch2", util.KeyMatch2Func)
	fm.AddFunction("keyGet2", util.KeyGet2Func)
	fm.AddFunction("keyMatch3", util.KeyMatch3Func)
	fm.AddFunction("keyMatch4", util.KeyMatch4Func)
	fm.AddFunction("keyMatch5", util.KeyMatch5Func)
	fm.AddFunction("regexMatch", util.RegexMatchFunc)
	fm.AddFunction("ipMatch", util.IPMatchFunc)
	fm.AddFunction("globMatch", util.GlobMatchFunc)

	return *fm
}

// GetFunctions return a map with all the functions
func (fm *FunctionMap) GetFunctions() map[string]govaluate.ExpressionFunction {
	ret := make(map[string]govaluate.ExpressionFunction)

	fm.fns.Range(func(k interface{}, v interface{}) bool {
		ret[k.(string)] = v.(govaluate.ExpressionFunction)
		return true
	})

	return ret
}
