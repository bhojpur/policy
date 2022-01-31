package model

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
	"errors"
	"strings"

	"github.com/bhojpur/policy/pkg/log"
	"github.com/bhojpur/policy/pkg/rbac"
)

// Assertion represents an expression in a section of the model.
// For example: r = sub, obj, act
type Assertion struct {
	Key       string
	Value     string
	Tokens    []string
	Policy    [][]string
	PolicyMap map[string]int
	RM        rbac.RoleManager

	logger        log.Logger
	priorityIndex int
}

func (ast *Assertion) buildIncrementalRoleLinks(rm rbac.RoleManager, op PolicyOp, rules [][]string) error {
	ast.RM = rm
	count := strings.Count(ast.Value, "_")
	if count < 2 {
		return errors.New("the number of \"_\" in role definition should be at least 2")
	}

	for _, rule := range rules {
		if len(rule) < count {
			return errors.New("grouping policy elements do not meet role definition")
		}
		if len(rule) > count {
			rule = rule[:count]
		}
		switch op {
		case PolicyAdd:
			err := rm.AddLink(rule[0], rule[1], rule[2:]...)
			if err != nil {
				return err
			}
		case PolicyRemove:
			err := rm.DeleteLink(rule[0], rule[1], rule[2:]...)
			if err != nil {
				return err
			}
		}
	}

	if op == PolicyAdd {
		for _, rule := range rules {
			err := rm.BuildRelationship(rule[0], rule[1], rule[2:]...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ast *Assertion) buildRoleLinks(rm rbac.RoleManager) error {
	ast.RM = rm
	count := strings.Count(ast.Value, "_")
	if count < 2 {
		return errors.New("the number of \"_\" in role definition should be at least 2")
	}
	for _, rule := range ast.Policy {
		if len(rule) < count {
			return errors.New("grouping policy elements do not meet role definition")
		}
		if len(rule) > count {
			rule = rule[:count]
		}
		err := ast.RM.AddLink(rule[0], rule[1], rule[2:]...)
		if err != nil {
			return err
		}
	}

	for _, rule := range ast.Policy {
		err := ast.RM.BuildRelationship(rule[0], rule[1], rule[2:]...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ast *Assertion) setLogger(logger log.Logger) {
	ast.logger = logger
}

func (ast *Assertion) initPriorityIndex() {
	ast.priorityIndex = -1
}

func (ast *Assertion) copy() *Assertion {
	tokens := append([]string(nil), ast.Tokens...)
	policy := make([][]string, len(ast.Policy))

	for i, p := range ast.Policy {
		policy[i] = append(policy[i], p...)
	}
	policyMap := make(map[string]int)
	for k, v := range ast.PolicyMap {
		policyMap[k] = v
	}

	newAst := &Assertion{
		Key:           ast.Key,
		Value:         ast.Value,
		PolicyMap:     policyMap,
		Tokens:        tokens,
		Policy:        policy,
		priorityIndex: ast.priorityIndex,
	}

	return newAst
}
