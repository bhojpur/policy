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
	"testing"
)

func BenchmarkCachedRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawEnforce("alice", "data1", "read")
	}
}

func BenchmarkCachedBasicModel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkCachedRBACModel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data2", "read")
	}
}

func BenchmarkCachedRBACModelSmall(b *testing.B) {
	e, _ := NewCachedEnforcer("examples/rbac_model.conf", false)
	// 100 roles, 10 resources.
	for i := 0; i < 100; i++ {
		_, err := e.AddPolicy(fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i/10), "read")
		if err != nil {
			b.Fatal(err)
		}
	}
	// 1000 users.
	for i := 0; i < 1000; i++ {
		_, err := e.AddGroupingPolicy(fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i/10))
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("user501", "data9", "read")
	}
}

func BenchmarkCachedRBACModelMedium(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_model.conf", false)
	// 1000 roles, 100 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}

	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	// 10000 users.
	gPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		gPolicies = append(gPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i/10)})
	}

	_, err = e.AddGroupingPolicies(gPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("user5001", "data150", "read")
	}
}

func BenchmarkCachedRBACModelLarge(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_model.conf", false)

	// 10000 roles, 1000 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	// 100000 users.
	gPolicies := make([][]string, 0)
	for i := 0; i < 100000; i++ {
		gPolicies = append(gPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i/10)})
	}
	_, err = e.AddGroupingPolicies(gPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("user50001", "data1500", "read")
	}
}

func BenchmarkCachedRBACModelWithResourceRoles(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_with_resource_roles_model.conf", "../../examples/rbac_with_resource_roles_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkCachedRBACModelWithDomains(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_with_domains_model.conf", "../../examples/rbac_with_domains_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "domain1", "data1", "read")
	}
}

func BenchmarkCachedABACModel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/abac_model.conf", false)
	data1 := newTestResource("data1", "alice")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", data1, "read")
	}
}

func BenchmarkCachedKeyMatchModel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/keymatch_model.conf", "../../examples/keymatch_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "/alice_data/resource1", "GET")
	}
}

func BenchmarkCachedRBACModelWithDeny(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_with_deny_model.conf", "../../examples/rbac_with_deny_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkCachedPriorityModel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/priority_model.conf", "../../examples/priority_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkCachedRBACModelMediumParallel(b *testing.B) {
	e, _ := NewCachedEnforcer("../../examples/rbac_model.conf", false)

	// 10000 roles, 1000 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 10000; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}
	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	// 100000 users.
	gPolicies := make([][]string, 0)
	for i := 0; i < 100000; i++ {
		gPolicies = append(gPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i/10)})
	}
	_, err = e.AddGroupingPolicies(gPolicies)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = e.Enforce("user5001", "data150", "read")
		}
	})
}
