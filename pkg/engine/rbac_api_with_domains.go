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

// GetUsersForRoleInDomain gets the users that has a role inside a domain. Add by Gordon
func (e *Enforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	res, _ := e.model["g"]["g"].RM.GetUsers(name, domain)
	return res
}

// GetRolesForUserInDomain gets the roles that a user has inside a domain.
func (e *Enforcer) GetRolesForUserInDomain(name string, domain string) []string {
	res, _ := e.model["g"]["g"].RM.GetRoles(name, domain)
	return res
}

// GetPermissionsForUserInDomain gets permissions for a user or role inside a domain.
func (e *Enforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	var res [][]string
	users, _ := e.GetRolesForUser(user, domain)
	users = append(users, user)
	for _, singleUser := range users {
		policy := e.GetFilteredPolicy(0, singleUser, domain)
		res = append(res, policy...)
	}
	return res
}

// AddRoleForUserInDomain adds a role for a user inside a domain.
// Returns false if the user already has the role (aka not affected).
func (e *Enforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return e.AddGroupingPolicy(user, role, domain)
}

// DeleteRoleForUserInDomain deletes a role for a user inside a domain.
// Returns false if the user does not have the role (aka not affected).
func (e *Enforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return e.RemoveGroupingPolicy(user, role, domain)
}

// DeleteRolesForUserInDomain deletes all roles for a user inside a domain.
// Returns false if the user does not have any roles (aka not affected).
func (e *Enforcer) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	roles, err := e.model["g"]["g"].RM.GetRoles(user, domain)
	if err != nil {
		return false, err
	}

	var rules [][]string
	for _, role := range roles {
		rules = append(rules, []string{user, role, domain})
	}

	return e.RemoveGroupingPolicies(rules)
}

// GetAllUsersByDomain would get all users associated with the domain.
func (e *Enforcer) GetAllUsersByDomain(domain string) []string {
	m := make(map[string]struct{})
	g := e.model["g"]["g"]
	p := e.model["p"]["p"]
	users := make([]string, 0)
	index := e.getDomainIndex("p")

	getUser := func(index int, policies [][]string, domain string, m map[string]struct{}) []string {
		if len(policies) == 0 || len(policies[0]) <= index {
			return []string{}
		}
		res := make([]string, 0)
		for _, policy := range policies {
			if _, ok := m[policy[0]]; policy[index] == domain && !ok {
				res = append(res, policy[0])
				m[policy[0]] = struct{}{}
			}
		}
		return res
	}

	users = append(users, getUser(2, g.Policy, domain, m)...)
	users = append(users, getUser(index, p.Policy, domain, m)...)
	return users
}

// DeleteAllUsersByDomain would delete all users associated with the domain.
func (e *Enforcer) DeleteAllUsersByDomain(domain string) (bool, error) {
	g := e.model["g"]["g"]
	p := e.model["p"]["p"]
	index := e.getDomainIndex("p")

	getUser := func(index int, policies [][]string, domain string) [][]string {
		if len(policies) == 0 || len(policies[0]) <= index {
			return [][]string{}
		}
		res := make([][]string, 0)
		for _, policy := range policies {
			if policy[index] == domain {
				res = append(res, policy)
			}
		}
		return res
	}

	users := getUser(2, g.Policy, domain)
	if _, err := e.RemoveGroupingPolicies(users); err != nil {
		return false, err
	}
	users = getUser(index, p.Policy, domain)
	if _, err := e.RemovePolicies(users); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteDomains would delete all associated users and roles.
// It would delete all domains if parameter is not provided.
func (e *Enforcer) DeleteDomains(domains ...string) (bool, error) {
	if len(domains) == 0 {
		e.ClearPolicy()
		return true, nil
	}
	for _, domain := range domains {
		if _, err := e.DeleteAllUsersByDomain(domain); err != nil {
			return false, err
		}
	}
	return true, nil
}
