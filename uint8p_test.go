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

var MarshalUint8pTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"uint8 pointer": log0.Uint8p(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil uint8 pointer": log0.Uint8p(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"any uint8 pointer": log0.Any(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"reflect uint8 pointer": log0.Reflect(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uint8 pointer":42
		}`,
	},
}

func TestMarshalUint8p(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint8pTestCases)
}
