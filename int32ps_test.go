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

var MarshalInt32psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"int32 pointer slice": log0.Int32ps(&f, &f2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"int32 pointer slice":[123,321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil int32 pointers": log0.Int32ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil int32 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int32 pointers": log0.Int32ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without int32 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of any int32 pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of any int32 pointers":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of reflects of int32 pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of reflects of int32 pointers":[123,321]
		}`,
	},
}

func TestMarshalInt32ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt32psTestCases)
}
