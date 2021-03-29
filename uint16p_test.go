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

var MarshalUint16pTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"uint16 pointer": log0.Uint16p(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uint16 pointer": log0.Uint16p(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"any uint16 pointer": log0.Any(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"reflect uint16 pointer": log0.Reflect(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i *uint16
			return map[string]json.Marshaler{"reflect uint16 pointer to nil": log0.Reflect(i)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
}

func TestMarshalUint16p(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint16pTestCases)
}
