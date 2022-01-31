package util

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
	"regexp"
	"sort"
	"strings"
)

var evalReg *regexp.Regexp = regexp.MustCompile(`\beval\((?P<rule>[^)]*)\)`)

// EscapeAssertion escapes the dots in the assertion, because the expression evaluation doesn't support such variable names.
func EscapeAssertion(s string) string {
	//Replace the first dot, because it can't be recognized by the regexp.
	if strings.HasPrefix(s, "r") || strings.HasPrefix(s, "p") {
		s = strings.Replace(s, ".", "_", 1)
	}
	var regex = regexp.MustCompile(`(\|| |=|\)|\(|&|<|>|,|\+|-|!|\*|\/)((r|p)[0-9]*)\.`)
	s = regex.ReplaceAllStringFunc(s, func(m string) string {
		return strings.Replace(m, ".", "_", 1)
	})
	return s
}

// RemoveComments removes the comments starting with # in the text.
func RemoveComments(s string) string {
	pos := strings.Index(s, "#")
	if pos == -1 {
		return s
	}
	return strings.TrimSpace(s[0:pos])
}

// ArrayEquals determines whether two string arrays are identical.
func ArrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Array2DEquals determines whether two 2-dimensional string arrays are identical.
func Array2DEquals(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !ArrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

// ArrayRemoveDuplicates removes any duplicated elements in a string array.
func ArrayRemoveDuplicates(s *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *s {
		if !found[x] {
			found[x] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

// ArrayToString gets a printable string for a string array.
func ArrayToString(s []string) string {
	return strings.Join(s, ", ")
}

// ParamsToString gets a printable string for variable number of parameters.
func ParamsToString(s ...string) string {
	return strings.Join(s, ", ")
}

// SetEquals determines whether two string sets are identical.
func SetEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// JoinSlice joins a string and a slice into a new slice.
func JoinSlice(a string, b ...string) []string {
	res := make([]string, 0, len(b)+1)

	res = append(res, a)
	res = append(res, b...)

	return res
}

// JoinSliceAny joins a string and a slice into a new interface{} slice.
func JoinSliceAny(a string, b ...string) []interface{} {
	res := make([]interface{}, 0, len(b)+1)

	res = append(res, a)
	for _, s := range b {
		res = append(res, s)
	}

	return res
}

// SetSubtract returns the elements in `a` that aren't in `b`.
func SetSubtract(a []string, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// HasEval determine whether matcher contains function eval
func HasEval(s string) bool {
	return evalReg.MatchString(s)
}

// ReplaceEval replace function eval with the value of its parameters
func ReplaceEval(s string, rule string) string {
	return evalReg.ReplaceAllString(s, "("+rule+")")
}

// ReplaceEvalWithMap replace function eval with the value of its parameters via given sets.
func ReplaceEvalWithMap(src string, sets map[string]string) string {
	return evalReg.ReplaceAllStringFunc(src, func(s string) string {
		subs := evalReg.FindStringSubmatch(s)
		if subs == nil {
			return s
		}
		key := subs[1]
		value, found := sets[key]
		if !found {
			return s
		}
		return evalReg.ReplaceAllString(s, value)
	})
}

// GetEvalValue returns the parameters of function eval
func GetEvalValue(s string) []string {
	subMatch := evalReg.FindAllStringSubmatch(s, -1)
	var rules []string
	for _, rule := range subMatch {
		rules = append(rules, rule[1])
	}
	return rules
}

func RemoveDuplicateElement(s []string) []string {
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
