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

var MarshalIntsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int slice": log0.Ints(123, 321)},
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"int slice":[123,321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without int": log0.Ints()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without int":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int = 123, 321
			return map[string]json.Marshaler{"slice of any int": log0.Anys(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of any int":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int = 123, 321
			return map[string]json.Marshaler{"slice of int reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "123 321",
		expectedText: "123 321",
		expectedJSON: `{
			"slice of int reflects":[123,321]
		}`,
	},
}

func TestMarshalInts(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalIntsTestCases)
}
