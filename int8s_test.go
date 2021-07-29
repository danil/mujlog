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

var MarshalInt8sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int8 slice": log0.Int8s(42, 77)},
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"int8 slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int8": log0.Int8s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without int8":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of any int8": log0.Anys(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any int8":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of int8 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of int8 reflects":[42,77]
		}`,
	},
}

func TestMarshalInt8s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt8sTestCases)
}
