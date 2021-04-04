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

var MarshalUintptrsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uintptr slice": log0.Uintptrs(42, 77)},
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uintptr slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uintptr": log0.Uintptrs()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without uintptr":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of any uintptr": log0.Anys(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uintptr":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of uintptr reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of uintptr reflects":[42,77]
		}`,
	},
}

func TestMarshalUintptrs(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintptrsTestCases)
}
