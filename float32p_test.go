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

var MarshalFloat32pTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"float32 pointer": log0.Float32p(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 0.123456789
			return map[string]json.Marshaler{"high precision float32 pointer": log0.Float32p(&f)}
		}(),
		expected:     "0.12345679",
		expectedText: "0.12345679",
		expectedJSON: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"float32 nil pointer": log0.Float32p(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"any float32 pointer": log0.Any(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"reflect float32 pointer": log0.Reflect(&f)}
		}(),
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float32
			return map[string]json.Marshaler{"reflect float32 pointer to nil": log0.Reflect(f)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect float32 pointer to nil":null
		}`,
	},
}

func TestMarshalFloat32p(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat32pTestCases)
}
