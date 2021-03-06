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

// GetRolesForUser gets the roles that a user has.
func (e *SyncedEnforcer) GetRolesForUser(name string, domain ...string) ([]string, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetRolesForUser(name, domain...)
}

// GetUsersForRole gets the users that has a role.
func (e *SyncedEnforcer) GetUsersForRole(name string, domain ...string) ([]string, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetUsersForRole(name, domain...)
}

// HasRoleForUser determines whether a user has a role.
func (e *SyncedEnforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.HasRoleForUser(name, role, domain...)
}

// AddRoleForUser adds a role for a user.
// Returns false if the user already has the role (aka not affected).
func (e *SyncedEnforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.AddRoleForUser(user, role, domain...)
}

// AddRolesForUser adds roles for a user.
// Returns false if the user already has the roles (aka not affected).
func (e *SyncedEnforcer) AddRolesForUser(user string, roles []string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.AddRolesForUser(user, roles, domain...)
}

// DeleteRoleForUser deletes a role for a user.
// Returns false if the user does not have the role (aka not affected).
func (e *SyncedEnforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteRoleForUser(user, role, domain...)
}

// DeleteRolesForUser deletes all roles for a user.
// Returns false if the user does not have any roles (aka not affected).
func (e *SyncedEnforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteRolesForUser(user, domain...)
}

// DeleteUser deletes a user.
// Returns false if the user does not exist (aka not affected).
func (e *SyncedEnforcer) DeleteUser(user string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteUser(user)
}

// DeleteRole deletes a role.
// Returns false if the role does not exist (aka not affected).
func (e *SyncedEnforcer) DeleteRole(role string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeleteRole(role)
}

// DeletePermission deletes a permission.
// Returns false if the permission does not exist (aka not affected).
func (e *SyncedEnforcer) DeletePermission(permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeletePermission(permission...)
}

// AddPermissionForUser adds a permission for a user or role.
// Returns false if the user or role already has the permission (aka not affected).
func (e *SyncedEnforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.AddPermissionForUser(user, permission...)
}

// DeletePermissionForUser deletes a permission for a user or role.
// Returns false if the user or role does not have the permission (aka not affected).
func (e *SyncedEnforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeletePermissionForUser(user, permission...)
}

// DeletePermissionsForUser deletes permissions for a user or role.
// Returns false if the user or role does not have any permissions (aka not affected).
func (e *SyncedEnforcer) DeletePermissionsForUser(user string) (bool, error) {
	e.m.Lock()
	defer e.m.Unlock()
	return e.Enforcer.DeletePermissionsForUser(user)
}

// GetPermissionsForUser gets permissions for a user or role.
func (e *SyncedEnforcer) GetPermissionsForUser(user string, domain ...string) [][]string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetPermissionsForUser(user, domain...)
}

// HasPermissionForUser determines whether a user has a permission.
func (e *SyncedEnforcer) HasPermissionForUser(user string, permission ...string) bool {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.HasPermissionForUser(user, permission...)
}

// GetImplicitRolesForUser gets implicit roles that a user has.
// Compared to GetRolesForUser(), this function retrieves indirect roles besides direct roles.
// For example:
// g, alice, role:admin
// g, role:admin, role:user
//
// GetRolesForUser("alice") can only get: ["role:admin"].
// But GetImplicitRolesForUser("alice") will get: ["role:admin", "role:user"].
func (e *SyncedEnforcer) GetImplicitRolesForUser(name string, domain ...string) ([]string, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetImplicitRolesForUser(name, domain...)
}

// GetImplicitPermissionsForUser gets implicit permissions for a user or role.
// Compared to GetPermissionsForUser(), this function retrieves permissions for inherited roles.
// For example:
// p, admin, data1, read
// p, alice, data2, read
// g, alice, admin
//
// GetPermissionsForUser("alice") can only get: [["alice", "data2", "read"]].
// But GetImplicitPermissionsForUser("alice") will get: [["admin", "data1", "read"], ["alice", "data2", "read"]].
func (e *SyncedEnforcer) GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetImplicitPermissionsForUser(user, domain...)
}

// GetImplicitUsersForPermission gets implicit users for a permission.
// For example:
// p, admin, data1, read
// p, bob, data1, read
// g, alice, admin
//
// GetImplicitUsersForPermission("data1", "read") will get: ["alice", "bob"].
// Note: only users will be returned, roles (2nd arg in "g") will be excluded.
func (e *SyncedEnforcer) GetImplicitUsersForPermission(permission ...string) ([]string, error) {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.Enforcer.GetImplicitUsersForPermission(permission...)
}
