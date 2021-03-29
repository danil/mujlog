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

var MarshalFloat64TestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"float64": log0.Float64(4.2)},
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"float64":4.2
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"high precision float64": log0.Float64(0.123456789)},
		expected:     "0.123456789",
		expectedText: "0.123456789",
		expectedJSON: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"zero float64": log0.Float64(0)},
		expected:     "0",
		expectedText: "0",
		expectedJSON: `{
			"zero float64":0
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any float64": log0.Any(4.2)},
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"any float64":4.2
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any zero float64": log0.Any(0)},
		expected:     "0",
		expectedText: "0",
		expectedJSON: `{
			"any zero float64":0
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect float64": log0.Reflect(4.2)},
		expected:     "4.2",
		expectedText: "4.2",
		expectedJSON: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect zero float64": log0.Reflect(0)},
		expected:     "0",
		expectedText: "0",
		expectedJSON: `{
			"reflect zero float64":0
		}`,
	},
}

func TestMarshalFloat64(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalFloat64TestCases)
}
