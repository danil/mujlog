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

var MarshalUintptrpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uintptr = 42, 77
			return map[string]json.Marshaler{"uintptr pointer slice": log0.Uintptrps(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uintptr pointer slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil uintptr pointers": log0.Uintptrps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil uintptr pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uintptr pointers": log0.Uintptrps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without uintptr pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of any uintptr pointers": log0.Anys(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uintptr pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of reflects of uintptr pointers": log0.Reflects(&f, &f2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of reflects of uintptr pointers":[42,77]
		}`,
	},
}

func TestMarshalUintptrps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintptrpsTestCases)
}
