// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalUintpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint = 42, 77
			return map[string]json.Marshaler{"uint pointer slice": log0.Uintps(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uint pointer slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil uint pointers": log0.Uintps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil uint pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uint pointers": log0.Uintps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without uint pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint = 42, 77
			return map[string]json.Marshaler{"slice of any uint pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uint pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint = 42, 77
			return map[string]json.Marshaler{"slice of reflects of uint pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of reflects of uint pointers":[42,77]
		}`,
	},
}

func TestMarshalUintps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintpsTestCases)
}
