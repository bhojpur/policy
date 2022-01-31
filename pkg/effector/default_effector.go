package effector

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

import "errors"

// DefaultEffector is default effector for Bhojpur Policy.
type DefaultEffector struct {
}

// NewDefaultEffector is the constructor for DefaultEffector.
func NewDefaultEffector() *DefaultEffector {
	e := DefaultEffector{}
	return &e
}

// MergeEffects merges all matching results collected by the enforcer into a single decision.
func (e *DefaultEffector) MergeEffects(expr string, effects []Effect, matches []float64, policyIndex int, policyLength int) (Effect, int, error) {
	result := Indeterminate
	explainIndex := -1

	switch expr {
	case "some(where (p_eft == allow))":
		if matches[policyIndex] == 0 {
			break
		}
		// only check the current policyIndex
		if effects[policyIndex] == Allow {
			result = Allow
			explainIndex = policyIndex
			break
		}
	case "!some(where (p_eft == deny))":
		// only check the current policyIndex
		if matches[policyIndex] != 0 && effects[policyIndex] == Deny {
			result = Deny
			explainIndex = policyIndex
			break
		}
		// if no deny rules are matched  at last, then allow
		if policyIndex == policyLength-1 {
			result = Allow
		}
	case "some(where (p_eft == allow)) && !some(where (p_eft == deny))":
		// short-circuit if matched deny rule
		if matches[policyIndex] != 0 && effects[policyIndex] == Deny {
			result = Deny
			// set hit rule to the (first) matched deny rule
			explainIndex = policyIndex
			break
		}

		// short-circuit some effects in the middle
		if policyIndex < policyLength-1 {
			// choose not to short-circuit
			return result, explainIndex, nil
		}
		// merge all effects at last
		for i, eft := range effects {
			if matches[i] == 0 {
				continue
			}

			if eft == Allow {
				result = Allow
				// set hit rule to first matched allow rule
				explainIndex = i
				break
			}
		}
	case "priority(p_eft) || deny", "subjectPriority(p_eft) || deny":
		// reverse merge, short-circuit may be earlier
		for i := len(effects) - 1; i >= 0; i-- {
			if matches[i] == 0 {
				continue
			}

			if effects[i] != Indeterminate {
				if effects[i] == Allow {
					result = Allow
				} else {
					result = Deny
				}
				explainIndex = i
				break
			}
		}
	default:
		return Deny, -1, errors.New("unsupported effect")
	}

	return result, explainIndex, nil
}
