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
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkHasPolicySmall(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 100 roles, 10 resources.
	for i := 0; i < 100; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.HasPolicy(fmt.Sprintf("user%d", rand.Intn(100)), fmt.Sprintf("data%d", rand.Intn(100)/10), "read")
	}
}

func BenchmarkHasPolicyMedium(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 1000 roles, 100 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.HasPolicy(fmt.Sprintf("user%d", rand.Intn(1000)), fmt.Sprintf("data%d", rand.Intn(1000)/10), "read")
	}
}

func BenchmarkHasPolicyLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 10000 roles, 1000 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.HasPolicy(fmt.Sprintf("user%d", rand.Intn(10000)), fmt.Sprintf("data%d", rand.Intn(10000)/10), "read")
	}
}

func BenchmarkAddPolicySmall(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 100 roles, 10 resources.
	for i := 0; i < 100; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", rand.Intn(100)+100), fmt.Sprintf("data%d", (rand.Intn(100)+100)/10), "read")
	}
}

func BenchmarkAddPolicyMedium(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 1000 roles, 100 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", rand.Intn(1000)+1000), fmt.Sprintf("data%d", (rand.Intn(1000)+1000)/10), "read")
	}
}

func BenchmarkAddPolicyLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 10000 roles, 1000 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", rand.Intn(10000)+10000), fmt.Sprintf("data%d", (rand.Intn(10000)+10000)/10), "read")
	}
}

func BenchmarkRemovePolicySmall(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 100 roles, 10 resources.
	for i := 0; i < 100; i++ {
		_, _ = e.AddPolicy(fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.RemovePolicy(fmt.Sprintf("user%d", rand.Intn(100)), fmt.Sprintf("data%d", rand.Intn(100)/10), "read")
	}
}

func BenchmarkRemovePolicyMedium(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 1000 roles, 100 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.RemovePolicy(fmt.Sprintf("user%d", rand.Intn(1000)), fmt.Sprintf("data%d", rand.Intn(1000)/10), "read")
	}
}

func BenchmarkRemovePolicyLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", false)

	// 10000 roles, 1000 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.RemovePolicy(fmt.Sprintf("user%d", rand.Intn(10000)), fmt.Sprintf("data%d", rand.Intn(10000)/10), "read")
	}
}
