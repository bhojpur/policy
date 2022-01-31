package cache

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

var ErrNoSuchKey = errors.New("there's no such key existing in cache")

type Cache interface {
	// Set puts key and value into cache.
	// First parameter for extra should be uint denoting expected survival time.
	// If survival time equals 0 or less, the key will always be survival.
	Set(key string, value bool, extra ...interface{}) error

	// Get returns result for key,
	// If there's no such key existing in cache,
	// ErrNoSuchKey will be returned.
	Get(key string) (bool, error)

	// Delete will remove the specific key in cache.
	// If there's no such key existing in cache,
	// ErrNoSuchKey will be returned.
	Delete(key string) error

	// Clear deletes all the items stored in cache.
	Clear() error
}
