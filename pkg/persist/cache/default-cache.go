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

type DefaultCache map[string]bool

func (c *DefaultCache) Set(key string, value bool, extra ...interface{}) error {
	(*c)[key] = value
	return nil
}

func (c *DefaultCache) Get(key string) (bool, error) {
	if res, ok := (*c)[key]; !ok {
		return false, ErrNoSuchKey
	} else {
		return res, nil
	}
}

func (c *DefaultCache) Delete(key string) error {
	if _, ok := (*c)[key]; !ok {
		return ErrNoSuchKey
	} else {
		delete(*c, key)
		return nil
	}
}

func (c *DefaultCache) Clear() error {
	*c = make(DefaultCache)
	return nil
}
