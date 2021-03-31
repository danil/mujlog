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

var MarshalInt8psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int8 = 42, 77
			return map[string]json.Marshaler{"int8 pointer slice": log0.Int8ps(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"int8 pointer slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil int8 pointers": log0.Int8ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil int8 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int8 pointers": log0.Int8ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without int8 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of any int8 pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any int8 pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of reflects of int8 pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of reflects of int8 pointers":[42,77]
		}`,
	},
}

func TestMarshalInt8ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt8psTestCases)
}
