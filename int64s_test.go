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

var MarshalInt64sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int64 slice": log0.Int64s(123, 321)},
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"int64 slice":[123,321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int64": log0.Int64s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without int64":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of any int64": log0.Anys(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of any int64":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int64 = 123, 321
			return map[string]json.Marshaler{"slice of int64 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of int64 reflects":[123,321]
		}`,
	},
}

func TestMarshalInt64s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt64sTestCases)
}
