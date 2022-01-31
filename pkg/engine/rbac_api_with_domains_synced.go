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
func (e *SyncedEnforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetUsersForRoleInDomain(name, domain)
}

// GetRolesForUserInDomain gets the roles that a user has inside a domain.
func (e *SyncedEnforcer) GetRolesForUserInDomain(name string, domain string) []string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetRolesForUserInDomain(name, domain)
}

// GetPermissionsForUserInDomain gets permissions for a user or role inside a domain.
func (e *SyncedEnforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetPermissionsForUserInDomain(user, domain)
}

// AddRoleForUserInDomain adds a role for a user inside a domain.
// Returns false if the user already has the role (aka not affected).
func (e *SyncedEnforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.AddRoleForUserInDomain(user, role, domain)
}

// DeleteRoleForUserInDomain deletes a role for a user inside a domain.
// Returns false if the user does not have the role (aka not affected).
func (e *SyncedEnforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteRoleForUserInDomain(user, role, domain)
}

// DeleteRolesForUserInDomain deletes all roles for a user inside a domain.
// Returns false if the user does not have any roles (aka not affected).
func (e *SyncedEnforcer) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteRolesForUserInDomain(user, domain)
}
