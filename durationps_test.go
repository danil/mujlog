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

var MarshalDurationpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of durations pointers": log0.Durationps(&d, &d2)}
		}(),
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of durations pointers":["42ns","42s"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil durations pointers": log0.Durationps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil durations pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without durations pointers": log0.Durationps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without durations pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of any duration pointers": log0.Anys(&d, &d2)}
		}(),
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of any duration pointers":["42ns","42s"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of reflect of duration pointers": log0.Reflects(&d, &d2)}
		}(),
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of reflect of duration pointers":[42,42000000000]
		}`,
	},
}

func TestMarshalDurationps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalDurationpsTestCases)
}
