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

var MarshalInt64TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"int64": log0.Int64(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"int64":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any int64": log0.Any(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any int64":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect int64": log0.Reflect(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect int64":42
		}`,
	},
}

func TestMarshalInt64(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalInt64TestCases)
}
