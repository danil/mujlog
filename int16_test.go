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

var MarshalInt16TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int16": log0.Int16(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"int16":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"any int16": log0.Any(i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any int16":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"reflect int16": log0.Reflect(i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect int16":42
		}`,
	},
}

func TestMarshalInt16(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt16TestCases)
}
