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

var MarshalStringpTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"string pointer": log0.Stringp(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"empty string pointer": log0.Stringp(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"empty string pointer":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil string pointer": log0.Stringp(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"any string pointer": log0.Any(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"any empty string pointer": log0.Any(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"reflect string pointer": log0.Reflect(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"reflect empty string pointer": log0.Reflect(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"reflect empty string pointer":""
		}`,
	},
}

func TestMarshalStingp(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalStringpTestCases)
}
