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

var MarshalUintptrTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uintptr": log0.Uintptr(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uintptr":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any uintp": log0.Any(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uintp":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect uintp": log0.Reflect(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uintp":42
		}`,
	},
}

func TestMarshalUintptr(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUintptrTestCases)
}
