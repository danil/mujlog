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

var MarshalRunesTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"runes": log0.Runes([]rune("Hello, Wörld!")...)},
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty runes": log0.Runes([]rune{}...)},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var p []rune
			return map[string]json.Marshaler{"nil runes": log0.Runes(p...)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil runes":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"rune slice with zero rune": log0.Runes([]rune{rune(0)}...)},
		expected:     "\\u0000",
		expectedText: "\\u0000",
		expectedJSON: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any runes": log0.Any([]rune("Hello, Wörld!"))},
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any empty runes": log0.Any([]rune{})},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"any empty runes":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any rune slice with zero rune": log0.Any([]rune{rune(0)})},
		expected:     "\\u0000",
		expectedText: "\\u0000",
		expectedJSON: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect runes": log0.Reflect([]rune("Hello, Wörld!"))},
		expected:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		expectedText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		expectedJSON: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect empty runes": log0.Reflect([]rune{})},
		expected:     "[]",
		expectedText: "[]",
		expectedJSON: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect rune slice with zero rune": log0.Reflect([]rune{rune(0)})},
		expected:     "[0]",
		expectedText: "[0]",
		expectedJSON: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
}

func TestMarshalRunes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalRunesTestCases)
}
