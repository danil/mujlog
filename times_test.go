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

var MarshalTimesTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"time slice": log0.Times(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		expected:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedJSON: `{
			"time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any time slice": log0.Anys(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		expected:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		expectedJSON: `{
			"any time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect time slice": log0.Reflects(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		expected:     "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		expectedText: "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		expectedJSON: `{
			"reflect time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
}

func TestMarshalTimes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTimesTestCases)
}
