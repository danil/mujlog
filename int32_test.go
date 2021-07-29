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

var MarshalInt32TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int32": log0.Int32(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"int32":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any int32": log0.Any(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any int32":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect int32": log0.Reflect(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect int32":42
		}`,
	},
}

func TestMarshalInt32(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt32TestCases)
}
