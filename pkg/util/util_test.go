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
	"testing"
)

func testEscapeAssertion(t *testing.T, s string, res string) {
	t.Helper()
	myRes := EscapeAssertion(s)
	t.Logf("%s: %s", s, myRes)

	if myRes != res {
		t.Errorf("%s: %s, supposed to be %s", s, myRes, res)
	}
}

func TestEscapeAssertion(t *testing.T) {
	testEscapeAssertion(t, "r.attr.value == p.attr", "r_attr.value == p_attr")
	testEscapeAssertion(t, "r.attp.value || p.attr", "r_attp.value || p_attr")
	testEscapeAssertion(t, "r.attp.value &&p.attr", "r_attp.value &&p_attr")
	testEscapeAssertion(t, "r.attp.value >p.attr", "r_attp.value >p_attr")
	testEscapeAssertion(t, "r.attp.value <p.attr", "r_attp.value <p_attr")
	testEscapeAssertion(t, "r.attp.value +p.attr", "r_attp.value +p_attr")
	testEscapeAssertion(t, "r.attp.value -p.attr", "r_attp.value -p_attr")
	testEscapeAssertion(t, "r.attp.value *p.attr", "r_attp.value *p_attr")
	testEscapeAssertion(t, "r.attp.value /p.attr", "r_attp.value /p_attr")
	testEscapeAssertion(t, "!r.attp.value /p.attr", "!r_attp.value /p_attr")
	testEscapeAssertion(t, "g(r.sub, p.sub) == p.attr", "g(r_sub, p_sub) == p_attr")
	testEscapeAssertion(t, "g(r.sub,p.sub) == p.attr", "g(r_sub,p_sub) == p_attr")
	testEscapeAssertion(t, "(r.attp.value || p.attr)p.u", "(r_attp.value || p_attr)p_u")
}

func testRemoveComments(t *testing.T, s string, res string) {
	t.Helper()
	myRes := RemoveComments(s)
	t.Logf("%s: %s", s, myRes)

	if myRes != res {
		t.Errorf("%s: %s, supposed to be %s", s, myRes, res)
	}
}

func TestRemoveComments(t *testing.T) {
	testRemoveComments(t, "r.act == p.act # comments", "r.act == p.act")
	testRemoveComments(t, "r.act == p.act#comments", "r.act == p.act")
	testRemoveComments(t, "r.act == p.act###", "r.act == p.act")
	testRemoveComments(t, "### comments", "")
	testRemoveComments(t, "r.act == p.act", "r.act == p.act")
}

func testArrayEquals(t *testing.T, a []string, b []string, res bool) {
	t.Helper()
	myRes := ArrayEquals(a, b)
	t.Logf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestArrayEquals(t *testing.T) {
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "b", "c"}, true)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "b"}, false)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{"a", "c", "b"}, false)
	testArrayEquals(t, []string{"a", "b", "c"}, []string{}, false)
}

func testArray2DEquals(t *testing.T, a [][]string, b [][]string, res bool) {
	t.Helper()
	myRes := Array2DEquals(a, b)
	t.Logf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestArray2DEquals(t *testing.T) {
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, true)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"a", "b", "c"}, {"1", "2"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{{"1", "2", "3"}, {"a", "b", "c"}}, false)
	testArray2DEquals(t, [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, [][]string{}, false)
}

func testSetEquals(t *testing.T, a []string, b []string, res bool) {
	t.Helper()
	myRes := SetEquals(a, b)
	t.Logf("%s == %s: %t", a, b, myRes)

	if myRes != res {
		t.Errorf("%s == %s: %t, supposed to be %t", a, b, myRes, res)
	}
}

func TestSetEquals(t *testing.T) {
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "b", "c"}, true)
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "b"}, false)
	testSetEquals(t, []string{"a", "b", "c"}, []string{"a", "c", "b"}, true)
	testSetEquals(t, []string{"a", "b", "c"}, []string{}, false)
}

func testContainEval(t *testing.T, s string, res bool) {
	t.Helper()
	myRes := HasEval(s)
	if myRes != res {
		t.Errorf("%s: %t, supposed to be %t", s, myRes, res)
	}
}
func TestContainEval(t *testing.T) {
	testContainEval(t, "eval() && a && b && c", true)
	testContainEval(t, "eval) && a && b && c", false)
	testContainEval(t, "eval)( && a && b && c", false)
	testContainEval(t, "eval(c * (a + b)) && a && b && c", true)
	testContainEval(t, "xeval() && a && b && c", false)
}

func testReplaceEval(t *testing.T, s string, rule string, res string) {
	t.Helper()
	myRes := ReplaceEval(s, rule)

	if myRes != res {
		t.Errorf("%s: %s supposed to be %s", s, myRes, res)
	}
}
func TestReplaceEval(t *testing.T) {
	testReplaceEval(t, "eval() && a && b && c", "a", "(a) && a && b && c")
	testReplaceEval(t, "eval() && a && b && c", "(a)", "((a)) && a && b && c")
}

func testGetEvalValue(t *testing.T, s string, res []string) {
	t.Helper()
	myRes := GetEvalValue(s)

	if !ArrayEquals(myRes, res) {
		t.Errorf("%s: %s supposed to be %s", s, myRes, res)
	}
}

func TestGetEvalValue(t *testing.T) {
	testGetEvalValue(t, "eval(a) && a && b && c", []string{"a"})
	testGetEvalValue(t, "a && eval(a) && b && c", []string{"a"})
	testGetEvalValue(t, "eval(a) && eval(b) && a && b && c", []string{"a", "b"})
	testGetEvalValue(t, "a && eval(a) && eval(b) && b && c", []string{"a", "b"})
}

func testReplaceEvalWithMap(t *testing.T, s string, sets map[string]string, res string) {
	t.Helper()
	myRes := ReplaceEvalWithMap(s, sets)

	if myRes != res {
		t.Errorf("%s: %s supposed to be %s", s, myRes, res)
	}
}

func TestReplaceEvalWithMap(t *testing.T) {
	testReplaceEvalWithMap(t, "eval(rule1)", map[string]string{"rule1": "a == b"}, "a == b")
	testReplaceEvalWithMap(t, "eval(rule1) && c && d", map[string]string{"rule1": "a == b"}, "a == b && c && d")
	testReplaceEvalWithMap(t, "eval(rule1)", nil, "eval(rule1)")
	testReplaceEvalWithMap(t, "eval(rule1) && c && d", nil, "eval(rule1) && c && d")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2)", map[string]string{"rule1": "a == b", "rule2": "a == c"}, "a == b || a == c")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2) && c && d", map[string]string{"rule1": "a == b", "rule2": "a == c"}, "a == b || a == c && c && d")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2)", map[string]string{"rule1": "a == b"}, "a == b || eval(rule2)")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2) && c && d", map[string]string{"rule1": "a == b"}, "a == b || eval(rule2) && c && d")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2)", map[string]string{"rule2": "a == b"}, "eval(rule1) || a == b")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2) && c && d", map[string]string{"rule2": "a == b"}, "eval(rule1) || a == b && c && d")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2)", nil, "eval(rule1) || eval(rule2)")
	testReplaceEvalWithMap(t, "eval(rule1) || eval(rule2) && c && d", nil, "eval(rule1) || eval(rule2) && c && d")
}
