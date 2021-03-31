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

var MarshalFloat64psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"float64 pointer slice": log0.Float64ps(&f, &f2)}
		}(),
		expected:     "0.123456789 0.987654321",
		expectedText: "0.123456789 0.987654321",
		expectedJSON: `{
			"float64 pointer slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil float64 pointers": log0.Float64ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil float64 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without float64 pointers": log0.Float64ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without float64 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of any float64 pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "0.123456789 0.987654321",
		expectedText: "0.123456789 0.987654321",
		expectedJSON: `{
			"slice of any float64 pointers":[0.123456789,0.987654321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of reflects of float64 pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "0.123456789 0.987654321",
		expectedText: "0.123456789 0.987654321",
		expectedJSON: `{
			"slice of reflects of float64 pointers":[0.123456789,0.987654321]
		}`,
	},
}

func TestMarshalFloat64ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat64psTestCases)
}
