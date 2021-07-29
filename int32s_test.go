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

var MarshalInt32sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int32 slice": log0.Int32s(123, 321)},
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"int32 slice":[123,321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int32": log0.Int32s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without int32":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of any int32": log0.Anys(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of any int32":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of int32 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of int32 reflects":[123,321]
		}`,
	},
}

func TestMarshalInt32s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt32sTestCases)
}
