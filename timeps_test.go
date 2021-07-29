// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"

	"github.com/kvlog/log0"
)

var MarshalTimepsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"time pointer slice": log0.Timeps(&t, &t2)}
		}(),
		expected:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedJSON: `{
			"time pointer slice":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil time pointers": log0.Timeps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil time pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without time pointers": log0.Timeps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without time pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"slice of any time pointers": log0.Anys(&t, &t2)}
		}(),
		expected:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedJSON: `{
			"slice of any time pointers":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"slice of reflects of time pointers": log0.Reflects(&t, &t2)}
		}(),
		expected:     "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		expectedText: "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		expectedJSON: `{
			"slice of reflects of time pointers":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
}

func TestMarshalTimeps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTimepsTestCases)
}
