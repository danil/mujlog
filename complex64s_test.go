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

var MarshalComplex64sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of complex64s": log0.Complex64s(complex(1, 23), complex(3, 21))},
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of complex64s":["1+23i","3+21i"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without complex64s": log0.Complex64s()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without complex64s":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex64s": log0.Anys(c, c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of any complex64s":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflect complex64": log0.Anys(c, c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of reflect complex64":["1+23i","3+21i"]
		}`,
	},
}

func TestMarshalComplex64s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex64sTestCases)
}
