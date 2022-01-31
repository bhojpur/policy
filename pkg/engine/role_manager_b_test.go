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

func BenchmarkRoleManagerSmall(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", false)
	// Do not rebuild the role inheritance relations for every AddGroupingPolicy() call.
	e.EnableAutoBuildRoleLinks(false)

	// 100 roles, 10 resources.
	pPolicies := make([][]string, 0)
	for i := 0; i < 100; i++ {
		pPolicies = append(pPolicies, []string{fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i/10), "read"})
	}

	_, err := e.AddPolicies(pPolicies)
	if err != nil {
		b.Fatal(err)
	}

	// 1000 users.
	gPolicies := make([][]string, 0)
	for i := 0; i < 1000; i++ {
		gPolicies = append(gPolicies, []string{fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i/10)})
	}

	_, err = e.AddGroupingPolicies(gPolicies)
	if err != nil {
		b.Fatal(err)
	}

	rm := e.GetRoleManager()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			_, _ = rm.HasLink("user501", fmt.Sprintf("group%d", j))
		}
	}
}

func BenchmarkRoleManagerMedium(b *testing.B) {
	e, _ := NewEnforcer("../../examples/rbac_model.conf", false)
	// Do not rebuild the role inheritance relations for every AddGroupingPolicy() call.
	e.EnableAutoBuildRoleLinks(false)

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

	err = e.BuildRoleLinks()
	if err != nil {
		b.Fatal(err)
	}

	rm := e.GetRoleManager()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			_, _ = rm.HasLink("user501", fmt.Sprintf("group%d", j))
		}
	}
}

func BenchmarkRoleManagerLarge(b *testing.B) {
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

	rm := e.GetRoleManager()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			_, _ = rm.HasLink("user501", fmt.Sprintf("group%d", j))
		}
	}
}

func BenchmarkBuildRoleLinksWithPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedMatchingFunc("g", "", util.KeyMatch4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.BuildRoleLinks()
	}
}

func BenchmarkBuildRoleLinksWithDomainPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedDomainMatchingFunc("g", "", util.KeyMatch4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.BuildRoleLinks()
	}
}

func BenchmarkBuildRoleLinksWithPatternAndDomainPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedMatchingFunc("g", "", util.KeyMatch4)
	e.AddNamedDomainMatchingFunc("g", "", util.KeyMatch4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.BuildRoleLinks()
	}
}

func BenchmarkHasLinkWithPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedMatchingFunc("g", "", util.KeyMatch4)
	_ = e.BuildRoleLinks()
	rm := e.rmMap["g"]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rm.HasLink("staffUser1001", "staff001", "/orgs/1/sites/site001")
	}
}

func BenchmarkHasLinkWithDomainPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedDomainMatchingFunc("g", "", util.KeyMatch4)
	_ = e.BuildRoleLinks()
	rm := e.rmMap["g"]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rm.HasLink("staffUser1001", "staff001", "/orgs/1/sites/site001")
	}

}

func BenchmarkHasLinkWithPatternAndDomainPatternLarge(b *testing.B) {
	e, _ := NewEnforcer("../../examples/performance/rbac_with_pattern_large_scale_model.conf", "../../examples/performance/rbac_with_pattern_large_scale_policy.csv")
	e.AddNamedMatchingFunc("g", "", util.KeyMatch4)
	e.AddNamedDomainMatchingFunc("g", "", util.KeyMatch4)
	_ = e.BuildRoleLinks()
	rm := e.rmMap["g"]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rm.HasLink("staffUser1001", "staff001", "/orgs/1/sites/site001")
	}
}
