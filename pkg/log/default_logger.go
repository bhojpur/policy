package log

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
	"log"
	"strings"
)

// DefaultLogger is the implementation for a Logger using Go logging interfaces.
type DefaultLogger struct {
	enabled bool
}

func (l *DefaultLogger) EnableLog(enable bool) {
	l.enabled = enable
}

func (l *DefaultLogger) IsEnabled() bool {
	return l.enabled
}

func (l *DefaultLogger) LogModel(model [][]string) {
	if !l.enabled {
		return
	}
	var str strings.Builder
	str.WriteString("Model: ")
	for _, v := range model {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	log.Println(str.String())
}

func (l *DefaultLogger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	if !l.enabled {
		return
	}

	var reqStr strings.Builder
	reqStr.WriteString("Request: ")
	for i, rval := range request {
		if i != len(request)-1 {
			reqStr.WriteString(fmt.Sprintf("%v, ", rval))
		} else {
			reqStr.WriteString(fmt.Sprintf("%v", rval))
		}
	}
	reqStr.WriteString(fmt.Sprintf(" ---> %t\n", result))

	reqStr.WriteString("Hit Policy: ")
	for i, pval := range explains {
		if i != len(explains)-1 {
			reqStr.WriteString(fmt.Sprintf("%v, ", pval))
		} else {
			reqStr.WriteString(fmt.Sprintf("%v \n", pval))
		}
	}

	log.Println(reqStr.String())
}

func (l *DefaultLogger) LogPolicy(policy map[string][][]string) {
	if !l.enabled {
		return
	}

	var str strings.Builder
	str.WriteString("Policy: ")
	for k, v := range policy {
		str.WriteString(fmt.Sprintf("%s : %v\n", k, v))
	}

	log.Println(str.String())
}

func (l *DefaultLogger) LogRole(roles []string) {
	if !l.enabled {
		return
	}

	log.Println("Roles: ", roles)
}
