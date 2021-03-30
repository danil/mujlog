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

var MarshalComplex64psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of complex64 pointers": log0.Complex64ps(&c, &c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of complex64 pointers":["1+23i","3+21i"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil complex64 pointers": log0.Complex64ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil complex64 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without complex64 pointers": log0.Complex64ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without complex64 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex64 pointers": log0.Anys(&c, &c2)}
		}(),
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"slice of any complex64 pointers":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflects of complex64 pointers": log0.Reflects(&c, &c2)}
		}(),
		expected:     "(1+23i) (3+21i)",
		expectedText: "(1+23i) (3+21i)",
		error:        errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
}

func TestMarshalComplex64ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex64psTestCases)
}
