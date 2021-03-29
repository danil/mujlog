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

var MarshalBoolsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bools true false": log0.Bools(true, false)},
		expected:     "true false",
		expectedText: "true false",
		expectedJSON: `{
			"bools true false":[true,false]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without bools": log0.Bools()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"without bools":[]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any bools": log0.Anys(true, false)},
		expected:     "true false",
		expectedText: "true false",
		expectedJSON: `{
			"any bools":[true, false]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflects bools": log0.Reflects(true, false)},
		expected:     "true false",
		expectedText: "true false",
		expectedJSON: `{
			"reflects bools":[true, false]
		}`,
	},
}

func TestMarshalBools(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBoolsTestCases)
}
