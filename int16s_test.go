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

var MarshalInt16sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int16 slice": log0.Int16s(123, 321)},
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"int16 slice":[123,321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int16": log0.Int16s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without int16":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of any int16": log0.Anys(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of any int16":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of int16 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of int16 reflects":[123,321]
		}`,
	},
}

func TestMarshalInt16s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt16sTestCases)
}
