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

var MarshalUint32sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"uint32 slice": log0.Uint32s(42, 77)},
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"uint32 slice":[42,77]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without uint32": log0.Uint32s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without uint32":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of any uint32": log0.Anys(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of any uint32":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of uint32 reflects": log0.Reflects(i, i2)}
		}(),
		expected:     "42 77",
		expectedText: "42 77",
		expectedJSON: `{
			"slice of uint32 reflects":[42,77]
		}`,
	},
}

func TestMarshalUint32s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint32sTestCases)
}
