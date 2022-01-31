package acl

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

// It contains functions to enforce access control lists.
// It allows you to register multiple security policies for enforcing
// ACLs for users or HTTP requests. The specific policy to use must be
// specified from a command line argument and cannot be changed on-the-fly.
//
// For actual authentication and authorization, you would need to implement your
// own policy as a package that calls RegisterPolicy(), and compile it into all
// Bhojpur.NET Platform application binaries that you use.
//
// By default (when no security_policy is specified), everyone is allowed to do
// anything.
//
// For convenience, there are two other built-in policies that also do NOT do
// any authentication, but allow you to globally disable some roles entirely:
//   * `deny-all` disallows all roles for everyone. Note that access is still
//     allowed to endpoints that are considered "public" (no ACL check at all).
//   * `read-only` allows anyone to act as DEBUGGING or MONITORING, but no one
//     is allowed to act as ADMIN. It also disallows any other custom roles that
//     are requested.

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	log "github.com/bhojpur/policy/pkg/engine"
)

// This is a list of predefined roles. Applications are free
// to invent more roles, as long as the acl policies they use can
// understand what they mean.
const (
	ADMIN      = "admin"
	DEBUGGING  = "debugging"
	MONITORING = "monitoring"
)

var (
	securityPolicy = flag.String("security_policy", "", "the name of a registered security policy to use for controlling access to URLs - empty means allow all for anyone (built-in policies: deny-all, read-only)")
	policies       = make(map[string]Policy)
	once           sync.Once
	currentPolicy  Policy
)

// Policy defines the interface that needs to be satisfied by
// ACL policy implementors.
type Policy interface {
	// CheckAccessActor can be called to verify if an actor
	// has access to the role.
	CheckAccessActor(actor, role string) error
	// CheckAccessHTTP can be called to verify if an actor in
	// the http request has access to the role.
	CheckAccessHTTP(req *http.Request, role string) error
}

// RegisterPolicy registers a security policy. This function must be called
// before the first call to CheckAccess happens, preferably through an init.
// This will ensure that the requested policy can be found by other acl
// functions when needed.
func RegisterPolicy(name string, policy Policy) {
	if _, ok := policies[name]; ok {
		log.Fatalf("policy %s is already registered", name)
	}
	policies[name] = policy
}

func savePolicy() {
	if *securityPolicy == "" {
		// Setting the policy to nil means Allow All from Anyone.
		currentPolicy = nil
		return
	}
	if policy, ok := policies[*securityPolicy]; ok {
		currentPolicy = policy
		return
	}
	log.Warningf("security_policy %q not found; using fallback policy (deny-all)", *securityPolicy)
	currentPolicy = denyAllPolicy{}
}

// CheckAccessActor uses the current security policy to
// verify if an actor has access to the role.
func CheckAccessActor(actor, role string) error {
	once.Do(savePolicy)
	if currentPolicy != nil {
		return currentPolicy.CheckAccessActor(actor, role)
	}
	return nil
}

// CheckAccessHTTP uses the current security policy to
// verify if an actor in an http request has access to
// the role.
func CheckAccessHTTP(req *http.Request, role string) error {
	once.Do(savePolicy)
	if currentPolicy != nil {
		return currentPolicy.CheckAccessHTTP(req, role)
	}
	return nil
}

// SendError is a convenience function that sends an ACL
// error as an HTTP response.
func SendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "Access denied: %v\n", err)
}
