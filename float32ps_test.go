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

var MarshalFloat32psTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"float32 pointer slice": log0.Float32ps(&f, &f2)}
		}(),
		expected:     "0.12345679 0.9876543",
		expectedText: "0.12345679 0.9876543",
		expectedJSON: `{
			"float32 pointer slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil float32 pointers": log0.Float32ps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil float32 pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without float32 pointers": log0.Float32ps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without float32 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of any float32 pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "0.12345679 0.9876543",
		expectedText: "0.12345679 0.9876543",
		expectedJSON: `{
			"slice of any float32 pointers":[0.123456789,0.987654321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of reflects of float32 pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "0.12345679 0.9876543",
		expectedText: "0.12345679 0.9876543",
		expectedJSON: `{
			"slice of reflects of float32 pointers":[0.123456789,0.987654321]
		}`,
	},
}

func TestMarshalFloat32ps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat32psTestCases)
}
