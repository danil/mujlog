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

var MarshalFloat32sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"float slice": log0.Float32s(0.123456789, 0.987654321)},
		expected:     "0.12345679 0.9876543",
		expectedText: "0.12345679 0.9876543",
		expectedJSON: `{
			"float slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without float32": log0.Float32s()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice without float32":[]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of any float32": log0.Anys(0.123456789, 0.987654321)},
		expected:     "0.123456789 0.987654321",
		expectedText: "0.123456789 0.987654321",
		expectedJSON: `{
			"slice of any float32":[0.123456789, 0.987654321]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflects float32s": log0.Reflects(0.123456789, 0.987654321)},
		expected:     "0.123456789 0.987654321",
		expectedText: "0.123456789 0.987654321",
		expectedJSON: `{
			"reflects float32s":[0.123456789, 0.987654321]
		}`,
	},
}

func TestMarshalFloat32s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat32sTestCases)
}
