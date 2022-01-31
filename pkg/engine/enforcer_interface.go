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
	"github.com/Knetic/govaluate"
	"github.com/bhojpur/policy/pkg/effector"
	"github.com/bhojpur/policy/pkg/model"
	"github.com/bhojpur/policy/pkg/persist"
	"github.com/bhojpur/policy/pkg/rbac"
)

var _ IEnforcer = &Enforcer{}
var _ IEnforcer = &SyncedEnforcer{}
var _ IEnforcer = &CachedEnforcer{}

// IEnforcer is the API interface of Enforcer
type IEnforcer interface {
	/* Enforcer API */
	InitWithFile(modelPath string, policyPath string) error
	InitWithAdapter(modelPath string, adapter persist.Adapter) error
	InitWithModelAndAdapter(m model.Model, adapter persist.Adapter) error
	LoadModel() error
	GetModel() model.Model
	SetModel(m model.Model)
	GetAdapter() persist.Adapter
	SetAdapter(adapter persist.Adapter)
	SetWatcher(watcher persist.Watcher) error
	GetRoleManager() rbac.RoleManager
	SetRoleManager(rm rbac.RoleManager)
	SetEffector(eft effector.Effector)
	ClearPolicy()
	LoadPolicy() error
	LoadFilteredPolicy(filter interface{}) error
	LoadIncrementalFilteredPolicy(filter interface{}) error
	IsFiltered() bool
	SavePolicy() error
	EnableEnforce(enable bool)
	EnableLog(enable bool)
	EnableAutoNotifyWatcher(enable bool)
	EnableAutoSave(autoSave bool)
	EnableAutoBuildRoleLinks(autoBuildRoleLinks bool)
	BuildRoleLinks() error
	Enforce(rvals ...interface{}) (bool, error)
	EnforceWithMatcher(matcher string, rvals ...interface{}) (bool, error)
	EnforceEx(rvals ...interface{}) (bool, []string, error)
	EnforceExWithMatcher(matcher string, rvals ...interface{}) (bool, []string, error)
	BatchEnforce(requests [][]interface{}) ([]bool, error)
	BatchEnforceWithMatcher(matcher string, requests [][]interface{}) ([]bool, error)

	/* RBAC API */
	GetRolesForUser(name string, domain ...string) ([]string, error)
	GetUsersForRole(name string, domain ...string) ([]string, error)
	HasRoleForUser(name string, role string, domain ...string) (bool, error)
	AddRoleForUser(user string, role string, domain ...string) (bool, error)
	AddPermissionForUser(user string, permission ...string) (bool, error)
	AddPermissionsForUser(user string, permissions ...[]string) (bool, error)
	DeletePermissionForUser(user string, permission ...string) (bool, error)
	DeletePermissionsForUser(user string) (bool, error)
	GetPermissionsForUser(user string, domain ...string) [][]string
	HasPermissionForUser(user string, permission ...string) bool
	GetImplicitRolesForUser(name string, domain ...string) ([]string, error)
	GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error)
	GetImplicitUsersForPermission(permission ...string) ([]string, error)
	DeleteRoleForUser(user string, role string, domain ...string) (bool, error)
	DeleteRolesForUser(user string, domain ...string) (bool, error)
	DeleteUser(user string) (bool, error)
	DeleteRole(role string) (bool, error)
	DeletePermission(permission ...string) (bool, error)

	/* RBAC API with domains*/
	GetUsersForRoleInDomain(name string, domain string) []string
	GetRolesForUserInDomain(name string, domain string) []string
	GetPermissionsForUserInDomain(user string, domain string) [][]string
	AddRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error)

	/* Management API */
	GetAllSubjects() []string
	GetAllNamedSubjects(ptype string) []string
	GetAllObjects() []string
	GetAllNamedObjects(ptype string) []string
	GetAllActions() []string
	GetAllNamedActions(ptype string) []string
	GetAllRoles() []string
	GetAllNamedRoles(ptype string) []string
	GetPolicy() [][]string
	GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string
	GetNamedPolicy(ptype string) [][]string
	GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string
	GetGroupingPolicy() [][]string
	GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) [][]string
	GetNamedGroupingPolicy(ptype string) [][]string
	GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string
	HasPolicy(params ...interface{}) bool
	HasNamedPolicy(ptype string, params ...interface{}) bool
	AddPolicy(params ...interface{}) (bool, error)
	AddPolicies(rules [][]string) (bool, error)
	AddNamedPolicy(ptype string, params ...interface{}) (bool, error)
	AddNamedPolicies(ptype string, rules [][]string) (bool, error)
	RemovePolicy(params ...interface{}) (bool, error)
	RemovePolicies(rules [][]string) (bool, error)
	RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error)
	RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error)
	RemoveNamedPolicies(ptype string, rules [][]string) (bool, error)
	RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error)
	HasGroupingPolicy(params ...interface{}) bool
	HasNamedGroupingPolicy(ptype string, params ...interface{}) bool
	AddGroupingPolicy(params ...interface{}) (bool, error)
	AddGroupingPolicies(rules [][]string) (bool, error)
	AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error)
	AddNamedGroupingPolicies(ptype string, rules [][]string) (bool, error)
	RemoveGroupingPolicy(params ...interface{}) (bool, error)
	RemoveGroupingPolicies(rules [][]string) (bool, error)
	RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error)
	RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error)
	RemoveNamedGroupingPolicies(ptype string, rules [][]string) (bool, error)
	RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error)
	AddFunction(name string, function govaluate.ExpressionFunction)

	UpdatePolicy(oldPolicy []string, newPolicy []string) (bool, error)
	UpdatePolicies(oldPolicies [][]string, newPolicies [][]string) (bool, error)
	UpdateFilteredPolicies(newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error)
}

var _ IDistributedEnforcer = &DistributedEnforcer{}

// IDistributedEnforcer defines dispatcher enforcer.
type IDistributedEnforcer interface {
	IEnforcer
	SetDispatcher(dispatcher persist.Dispatcher)
	/* Management API for DistributedEnforcer*/
	AddPoliciesSelf(shouldPersist func() bool, sec string, ptype string, rules [][]string) (effected [][]string, err error)
	RemovePoliciesSelf(shouldPersist func() bool, sec string, ptype string, rules [][]string) (effected [][]string, err error)
	RemoveFilteredPolicySelf(shouldPersist func() bool, sec string, ptype string, fieldIndex int, fieldValues ...string) (effected [][]string, err error)
	ClearPolicySelf(shouldPersist func() bool) error
	UpdatePolicySelf(shouldPersist func() bool, sec string, ptype string, oldRule, newRule []string) (effected bool, err error)
	UpdatePoliciesSelf(shouldPersist func() bool, sec string, ptype string, oldRules, newRules [][]string) (effected bool, err error)
	UpdateFilteredPoliciesSelf(shouldPersist func() bool, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) (bool, error)
}
