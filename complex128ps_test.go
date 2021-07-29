// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalComplex128psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"complex128 pointers slice": log0.Complex128ps(&c, &c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"complex128 pointers slice":["1+23i","3+21i"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil complex128 pointers": log0.Complex128ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil complex128 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without complex128 pointers": log0.Complex128ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without complex128 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex128 pointers": log0.Anys(&c, &c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of any complex128 pointers":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflects of complex128 pointers": log0.Reflects(&c, &c2)}
		}(),
		expected:      "(1+23i) (3+21i)",
		expectedText:  "(1+23i) (3+21i)",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex128psTestCases)
}
