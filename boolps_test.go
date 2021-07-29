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

var MarshalBoolpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b, b2 := true, false
			return map[string]json.Marshaler{"bool pointers to true and false": log0.Boolps(&b, &b2)}
		}(),
		expected:     "true false",
		expectedText: "true false",
		expectedJSON: `{
			"bool pointers to true and false":[true,false]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bool pointers to nil": log0.Boolps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"bool pointers to nil":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without bool pointers": log0.Boolps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"without bool pointers":null
		}`,
	},
}

func TestMarshalBoolps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBoolpsTestCases)
}
