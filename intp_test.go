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

var MarshalIntpTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"int pointer": log0.Intp(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"any int pointer": log0.Any(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"reflect int pointer": log0.Reflect(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect int pointer":42
		}`,
	},
}

func TestMarshalIntp(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalIntpTestCases)
}
