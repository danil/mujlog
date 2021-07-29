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

var MarshalUintsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uint slice": log0.Uints(42, 77)},
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uint slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uint": log0.Uints()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without uint":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint = 42, 77
			return map[string]json.Marshaler{"slice of any uint": log0.Anys(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uint":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint = 42, 77
			return map[string]json.Marshaler{"slice of uint reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of uint reflects":[42,77]
		}`,
	},
}

func TestMarshalUints(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintsTestCases)
}
