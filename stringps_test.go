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

var MarshalStringpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"string pointer slice": log0.Stringps(&f, &f2)}
		}(),
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"string pointer slice":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil string pointers": log0.Stringps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil string pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without string pointers": log0.Stringps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without string pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"slice of any string pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"slice of any string pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 string = "Hello, Wörld!", "Hello, World!"
			return map[string]json.Marshaler{"slice of reflects of string pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"slice of reflects of string pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
}

func TestMarshalStringps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalStringpsTestCases)
}
