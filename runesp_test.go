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

var MarshalRunespTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"runes pointer": log0.Runesp(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"empty runes pointer": log0.Runesp(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil runes pointer": log0.Runesp(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"any runes pointer": log0.Any(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"any empty runes pointer": log0.Any(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect runes pointer": log0.Reflect(&p)}
		}(),
		expected:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		expectedText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		expectedJSON: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"reflect empty runes pointer": log0.Reflect(&p)}
		}(),
		expected:     "[]",
		expectedText: "[]",
		expectedJSON: `{
			"reflect empty runes pointer":[]
		}`,
	},
}

func TestMarshalRunesp(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalRunespTestCases)
}
