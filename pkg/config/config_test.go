package config

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
	"testing"
)

func TestGet(t *testing.T) {
	config, cerr := NewConfig("testdata/testini.ini")
	if cerr != nil {
		t.Errorf("Configuration file loading failed, err:%v", cerr.Error())
		t.Fatalf("err: %v", cerr)
	}

	// default::key test
	if v, err := config.Bool("debug"); err != nil || !v {
		t.Errorf("Get failure: expected different value for debug (expected: [%#v] got: [%#v])", true, v)
		t.Fatalf("err: %v", err)
	}
	if v := config.String("url"); v != "act.wiki" {
		t.Errorf("Get failure: expected different value for url (expected: [%#v] got: [%#v])", "act.wiki", v)
	}

	// redis::key test
	if v := config.Strings("redis::redis.key"); len(v) != 2 || v[0] != "push1" || v[1] != "push2" {
		t.Errorf("Get failure: expected different value for redis::redis.key (expected: [%#v] got: [%#v])", "[]string{push1,push2}", v)
	}
	if v := config.String("mysql::mysql.dev.host"); v != "127.0.0.1" {
		t.Errorf("Get failure: expected different value for mysql::mysql.dev.host (expected: [%#v] got: [%#v])", "127.0.0.1", v)
	}
	if v := config.String("mysql::mysql.master.host"); v != "10.0.0.1" {
		t.Errorf("Get failure: expected different value for mysql::mysql.master.host (expected: [%#v] got: [%#v])", "10.0.0.1", v)
	}
	if v := config.String("mysql::mysql.master.user"); v != "root" {
		t.Errorf("Get failure: expected different value for mysql::mysql.master.user (expected: [%#v] got: [%#v])", "root", v)
	}
	if v := config.String("mysql::mysql.master.pass"); v != "89dds)2$" {
		t.Errorf("Get failure: expected different value for mysql::mysql.master.pass (expected: [%#v] got: [%#v])", "89dds)2$", v)
	}
	// math::key test
	if v, err := config.Int64("math::math.i64"); err != nil || v != 64 {
		t.Errorf("Get failure: expected different value for math::math.i64 (expected: [%#v] got: [%#v])", 64, v)
		t.Fatalf("err: %v", err)
	}
	if v, err := config.Float64("math::math.f64"); err != nil || v != 64.1 {
		t.Errorf("Get failure: expected different value for math::math.f64 (expected: [%#v] got: [%#v])", 64.1, v)
		t.Fatalf("err: %v", err)
	}

	_ = config.Set("other::key1", "new test key")

	if v := config.String("other::key1"); v != "new test key" {
		t.Errorf("Get failure: expected different value for other::key1 (expected: [%#v] got: [%#v])", "new test key", v)
	}

	_ = config.Set("other::key1", "test key")

	if v := config.String("multi1::name"); v != "r.sub==p.sub && r.obj==p.obj" {
		t.Errorf("Get failure: expected different value for multi1::name (expected: [%#v] got: [%#v])", "r.sub==p.sub&&r.obj==p.obj", v)
	}

	if v := config.String("multi2::name"); v != "r.sub==p.sub && r.obj==p.obj" {
		t.Errorf("Get failure: expected different value for multi2::name (expected: [%#v] got: [%#v])", "r.sub==p.sub&&r.obj==p.obj", v)
	}

	if v := config.String("multi3::name"); v != "r.sub==p.sub && r.obj==p.obj" {
		t.Errorf("Get failure: expected different value for multi3::name (expected: [%#v] got: [%#v])", "r.sub==p.sub&&r.obj==p.obj", v)
	}

	if v := config.String("multi4::name"); v != "" {
		t.Errorf("Get failure: expected different value for multi4::name (expected: [%#v] got: [%#v])", "", v)
	}

	if v := config.String("multi5::name"); v != "r.sub==p.sub && r.obj==p.obj" {
		t.Errorf("Get failure: expected different value for multi5::name (expected: [%#v] got: [%#v])", "r.sub==p.sub&&r.obj==p.obj", v)
	}
}
