package rbac

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

import "github.com/bhojpur/policy/pkg/log"

// RoleManager provides interface to define the operations for managing roles.
type RoleManager interface {
	// Clear clears all stored data and resets the role manager to the initial state.
	Clear() error
	// AddLink adds the inheritance link between two roles. role: name1 and role: name2.
	// domain is a prefix to the roles (can be used for other purposes).
	AddLink(name1 string, name2 string, domain ...string) error
	BuildRelationship(name1 string, name2 string, domain ...string) error
	// DeleteLink deletes the inheritance link between two roles. role: name1 and role: name2.
	// domain is a prefix to the roles (can be used for other purposes).
	DeleteLink(name1 string, name2 string, domain ...string) error
	// HasLink determines whether a link exists between two roles. role: name1 inherits role: name2.
	// domain is a prefix to the roles (can be used for other purposes).
	HasLink(name1 string, name2 string, domain ...string) (bool, error)
	// GetRoles gets the roles that a user inherits.
	// domain is a prefix to the roles (can be used for other purposes).
	GetRoles(name string, domain ...string) ([]string, error)
	// GetUsers gets the users that inherits a role.
	// domain is a prefix to the users (can be used for other purposes).
	GetUsers(name string, domain ...string) ([]string, error)
	// GetDomains gets domains that a user has
	GetDomains(name string) ([]string, error)
	// PrintRoles prints all the roles to log.
	PrintRoles() error
	// SetLogger sets role manager's logger.
	SetLogger(logger log.Logger)
}
