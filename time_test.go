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

var MarshalTimeTestTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"time": time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC)},
		expected:     "1970-01-01 02:03:04.000000042 +0000 UTC",
		expectedText: "1970-01-01T02:03:04.000000042Z",
		expectedJSON: `{
			"time":"1970-01-01T02:03:04.000000042Z"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any time": log0.Any(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC))},
		expected:     `1970-01-01 02:03:04.000000042 +0000 UTC`,
		expectedText: `1970-01-01T02:03:04.000000042Z`,
		expectedJSON: `{
			"any time":"1970-01-01T02:03:04.000000042Z"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect time": log0.Reflect(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC))},
		expected:     "1970-01-01 02:03:04.000000042 +0000 UTC",
		expectedText: "1970-01-01 02:03:04.000000042 +0000 UTC",
		expectedJSON: `{
			"reflect time":"1970-01-01T02:03:04.000000042Z"
		}`,
	},
}

func TestTimeMarshalTest(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTimeTestTestCases)
}
