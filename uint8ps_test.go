// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalUint8psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint8 = 42, 77
			return map[string]json.Marshaler{"uint8 pointer slice": log0.Uint8ps(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uint8 pointer slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil uint8 pointers": log0.Uint8ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil uint8 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uint8 pointers": log0.Uint8ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without uint8 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of any uint8 pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uint8 pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of reflects of uint8 pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of reflects of uint8 pointers":[42,77]
		}`,
	},
}

func TestMarshalUint8ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint8psTestCases)
}
