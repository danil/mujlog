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

var MarshalComplex128TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"complex128": log0.Complex128(complex(1, 23))},
		expected:     "1+23i",
		expectedText: "1+23i",
		expectedJSON: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any complex128": log0.Any(complex(1, 23))},
		expected:     "1+23i",
		expectedText: "1+23i",
		expectedJSON: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"reflect complex128": log0.Reflect(complex(1, 23))},
		expected:      "(1+23i)",
		expectedText:  "(1+23i)",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex128TestCases)
}
