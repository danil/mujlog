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

var MarshalUint8TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uint8": log0.Uint8(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uint8":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any uint8": log0.Any(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uint8":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect uint8": log0.Reflect(42)},
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uint8":42
		}`,
	},
}

func TestMarshalUint8(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint8TestCases)
}
