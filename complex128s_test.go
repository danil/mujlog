// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalComplex128sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"complex128 slice": log0.Complex128s(complex(1, 23), complex(3, 21))},
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"complex128 slice":["1+23i","3+21i"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without complex128": log0.Complex128s()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without complex128":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex128": log0.Anys(c, c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of any complex128":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflect of complex128 pointers": log0.Reflects(c, c2)}
		}(),
		expected:      "(1+23i) (3+21i)",
		expectedText:  "(1+23i) (3+21i)",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex128sTestCases)
}
