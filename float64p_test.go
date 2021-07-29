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

var MarshalFloat64pTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"float64 pointer": log0.Float64p(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 0.123456789
			return map[string]json.Marshaler{"high precision float64 pointer": log0.Float64p(&f)}
		}(),
		expected:     "0.123456789",
		expectedText: "0.123456789",
		expectedJSON: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"float64 nil pointer": log0.Float64p(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"any float64 pointer": log0.Any(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"reflect float64 pointer": log0.Reflect(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float64
			return map[string]json.Marshaler{"reflect float64 pointer to nil": log0.Reflect(f)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect float64 pointer to nil":null
		}`,
	},
}

func TestMarshalFloat64p(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat64pTestCases)
}
