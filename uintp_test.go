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

var MarshalUintpTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"uint pointer": log0.Uintp(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uint pointer":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil uint pointer": log0.Uintp(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"any uint pointer": log0.Any(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"reflect uint pointer": log0.Reflect(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uint pointer":42
		}`,
	},
}

func TestMarshalUintp(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintpTestCases)
}
