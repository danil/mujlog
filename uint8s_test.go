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

var MarshalUint8sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uint8 slice": log0.Uint8s(42, 77)},
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uint8 slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uint8": log0.Uint8s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without uint8":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of any uint8": log0.Anys(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uint8":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of uint8 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of uint8 reflects":[42,77]
		}`,
	},
}

func TestMarshalUint8s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint8sTestCases)
}
