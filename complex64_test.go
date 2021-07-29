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

var MarshalComplex64TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"complex64": log0.Complex64(complex(1, 23))},
		expected:     "1+23i",
		expectedText: "1+23i",
		expectedJSON: `{
			"complex64":"1+23i"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"complex64": log0.Complex64(complex(3, 21))},
		expected:     "3+21i",
		expectedText: "3+21i",
		expectedJSON: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any complex64": log0.Any(complex(1, 23))},
		expected:     "1+23i",
		expectedText: "1+23i",
		expectedJSON: `{
			"any complex64":"1+23i"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any complex64": log0.Any(complex(3, 21))},
		expected:     "3+21i",
		expectedText: "3+21i",
		expectedJSON: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"reflect complex64": log0.Reflect(complex(1, 23))},
		expected:      "(1+23i)",
		expectedText:  "(1+23i)",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"reflect complex64": log0.Reflect(complex(3, 21))},
		expected:      "(3+21i)",
		expectedText:  "(3+21i)",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex64(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex64TestCases)
}
