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

	"github.com/bhojpur/policy/pkg/util"
)

func rawEnforce(sub string, obj string, act string) bool {
	policy := [2][3]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}}
	for _, rule := range policy {
		if sub == rule[0] && obj == rule[1] && act == rule[2] {
			return true
		}
	}
	return false
}

func BenchmarkRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawEnforce("alice", "data1", "read")
	}
}

func BenchmarkBasicModel(b *testing.B) {
	e, _ := NewEnforcer("../../examples/basic_model.conf", "../../examples/basic_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkRBACModel(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", "../../examples/rbac_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data2", "read")
	}
}

func BenchmarkRBACModelSmall(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", false)

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

func BenchmarkRBACModelMedium(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", false)

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
		_, _ = e.Enforce("user5001", "data99", "read")
	}
}

func BenchmarkRBACModelLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", false)

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
		_, _ = e.Enforce("user50001", "data999", "read")
	}
}

func BenchmarkRBACModelWithResourceRoles(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_with_resource_roles_model.conf", "../../examples/rbac_with_resource_roles_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkRBACModelWithDomains(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_with_domains_model.conf", "../../examples/rbac_with_domains_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "domain1", "data1", "read")
	}
}

func BenchmarkABACModel(b *testing.B) {
	e, _ := NewEnforcer("../../examples/abac_model.conf", false)
	data1 := newTestResource("data1", "alice")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", data1, "read")
	}
}

func BenchmarkKeyMatchModel(b *testing.B) {
	e, _ := NewEnforcer("../../examples/keymatch_model.conf", "../../examples/keymatch_policy.csv", false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "/alice_data/resource1", "GET")
	}
}

func BenchmarkRBACModelWithDeny(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_with_deny_model.conf", "../../examples/rbac_with_deny_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkPriorityModel(b *testing.B) {
	e, _ := NewEnforcer("../../examples/priority_model.conf", "../../examples/priority_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkRBACModelWithDomainPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedDomainMatchingFunc("g", "", util.KeyMatch4)
	_ = e.BuildRoleLinks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = e.Enforce("staffUser1001", "/orgs/1/sites/site001", "App001.Module001.Action1001")
	}
}
