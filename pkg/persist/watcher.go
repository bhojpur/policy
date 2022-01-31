package persist

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

// Watcher is the interface for Bhojpur Policy watchers.
type Watcher interface {
	// SetUpdateCallback sets the callback function that the watcher will call
	// when the policy in DB has been changed by other instances.
	// A classic callback is Enforcer.LoadPolicy().
	SetUpdateCallback(func(string)) error
	// Update calls the update callback of other instances to synchronize their policy.
	// It is usually called after changing the policy in DB, like Enforcer.SavePolicy(),
	// Enforcer.AddPolicy(), Enforcer.RemovePolicy(), etc.
	Update() error
	// Close stops and releases the watcher, the callback function will not be called any more.
	Close()
}
