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

var MarshalFloat64sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"float64 slice": log0.Float64s(0.123456789, 0.987654641)},
		expected:     "0.123456789 0.987654641",
		expectedText: "0.123456789 0.987654641",
		expectedJSON: `{
			"float64 slice":[0.123456789,0.987654641]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without float64": log0.Float64s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without float64":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654641
			return map[string]json.Marshaler{"slice of any float64": log0.Anys(f, f2)}
		}(),
		expected:     "0.123456789 0.987654641",
		expectedText: "0.123456789 0.987654641",
		expectedJSON: `{
			"slice of any float64":[0.123456789, 0.987654641]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654641
			return map[string]json.Marshaler{"slice of float64 reflects": log0.Reflects(f, f2)}
		}(),
		expected:     "0.123456789 0.987654641",
		expectedText: "0.123456789 0.987654641",
		expectedJSON: `{
			"slice of float64 reflects":[0.123456789, 0.987654641]
		}`,
	},
}

func TestMarshalFloat64s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat64sTestCases)
}
