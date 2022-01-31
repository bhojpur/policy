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
	"strings"
	"sync"
	"sync/atomic"

	"github.com/bhojpur/policy/pkg/persist/cache"
)

// CachedEnforcer wraps Enforcer and provides decision cache
type CachedEnforcer struct {
	*Enforcer
	expireTime  uint
	cache       cache.Cache
	enableCache int32
	locker      *sync.RWMutex
}

type CacheableParam interface {
	GetCacheKey() string
}

// NewCachedEnforcer creates a cached enforcer via file or DB.
func NewCachedEnforcer(params ...interface{}) (*CachedEnforcer, error) {
	e := &CachedEnforcer{}
	var err error
	e.Enforcer, err = NewEnforcer(params...)
	if err != nil {
		return nil, err
	}

	e.enableCache = 1
	cache := cache.DefaultCache(make(map[string]bool))
	e.cache = &cache
	e.locker = new(sync.RWMutex)
	return e, nil
}

// EnableCache determines whether to enable cache on Enforce(). When enableCache is enabled, cached result (true | false) will be returned for previous decisions.
func (e *CachedEnforcer) EnableCache(enableCache bool) {
	var enabled int32
	if enableCache {
		enabled = 1
	}
	atomic.StoreInt32(&e.enableCache, enabled)
}

// Enforce decides whether a "subject" can access a "object" with the operation "action", input parameters are usually: (sub, obj, act).
// if rvals is not string , ingore the cache
func (e *CachedEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	if atomic.LoadInt32(&e.enableCache) == 0 {
		return e.Enforcer.Enforce(rvals...)
	}

	key, ok := e.getKey(rvals...)
	if !ok {
		return e.Enforcer.Enforce(rvals...)
	}

	if res, err := e.getCachedResult(key); err == nil {
		return res, nil
	} else if err != cache.ErrNoSuchKey {
		return res, err
	}

	res, err := e.Enforcer.Enforce(rvals...)
	if err != nil {
		return false, err
	}

	err = e.setCachedResult(key, res, e.expireTime)
	return res, err
}

func (e *CachedEnforcer) LoadPolicy() error {
	if atomic.LoadInt32(&e.enableCache) != 0 {
		if err := e.cache.Clear(); err != nil {
			return err
		}
	}
	return e.Enforcer.LoadPolicy()
}

func (e *CachedEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	if atomic.LoadInt32(&e.enableCache) != 0 {
		key, ok := e.getKey(params...)
		if ok {
			if err := e.cache.Delete(key); err != nil && err != cache.ErrNoSuchKey {
				return false, err
			}
		}
	}
	return e.Enforcer.RemovePolicy(params...)
}

func (e *CachedEnforcer) RemovePolicies(rules [][]string) (bool, error) {
	if len(rules) != 0 {
		if atomic.LoadInt32(&e.enableCache) != 0 {
			irule := make([]interface{}, len(rules[0]))
			for _, rule := range rules {
				for i, param := range rule {
					irule[i] = param
				}
				key, _ := e.getKey(irule...)
				if err := e.cache.Delete(key); err != nil && err != cache.ErrNoSuchKey {
					return false, err
				}
			}
		}
	}
	return e.Enforcer.RemovePolicies(rules)
}

func (e *CachedEnforcer) getCachedResult(key string) (res bool, err error) {
	e.locker.RLock()
	defer e.locker.RUnlock()
	return e.cache.Get(key)
}

func (e *CachedEnforcer) SetExpireTime(expireTime uint) {
	e.expireTime = expireTime
}

func (e *CachedEnforcer) SetCache(c cache.Cache) {
	e.cache = c
}

func (e *CachedEnforcer) setCachedResult(key string, res bool, extra ...interface{}) error {
	e.locker.Lock()
	defer e.locker.Unlock()
	return e.cache.Set(key, res, extra...)
}

func (e *CachedEnforcer) getKey(params ...interface{}) (string, bool) {
	key := strings.Builder{}
	for _, param := range params {
		switch typedParam := param.(type) {
		case string:
			key.WriteString(typedParam)
		case CacheableParam:
			key.WriteString(typedParam.GetCacheKey())
		default:
			return "", false
		}
		key.WriteString("$$")
	}
	return key.String(), true
}

// InvalidateCache deletes all the existing cached decisions.
func (e *CachedEnforcer) InvalidateCache() error {
	e.locker.Lock()
	defer e.locker.Unlock()
	return e.cache.Clear()
}
