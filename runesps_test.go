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

var MarshalRunespsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []rune("Hello, Wörld!"), []rune("Hello, World!")
			return map[string]json.Marshaler{"slice of rune slice pointers": log0.Runesps(&p, &p2)}
		}(),
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"slice of rune slice pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []rune{}, []rune{}
			return map[string]json.Marshaler{"slice of empty rune slice pointers": log0.Runesps(&p, &p2)}
		}(),
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"slice of empty rune slice pointers":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil rune slice pointers": log0.Runesps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil rune slice pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty slice of rune slice pointers": log0.Runesps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"empty slice of rune slice pointers":null
		}`,
	},
}

func TestMarshalRunesps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalRunespsTestCases)
}
