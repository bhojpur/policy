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

var logger Logger = &DefaultLogger{}

// SetLogger sets the current logger.
func SetLogger(l Logger) {
	logger = l
}

// GetLogger returns the current logger.
func GetLogger() Logger {
	return logger
}

// LogModel logs the model information.
func LogModel(model [][]string) {
	logger.LogModel(model)
}

// LogEnforce logs the enforcer information.
func LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	logger.LogEnforce(matcher, request, result, explains)
}

// LogRole log info related to role.
func LogRole(roles []string) {
	logger.LogRole(roles)
}

// LogPolicy logs the policy information.
func LogPolicy(policy map[string][][]string) {
	logger.LogPolicy(policy)
}
