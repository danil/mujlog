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

var MarshalDurationsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of durations": log0.Durations(42*time.Nanosecond, 42*time.Second)},
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of durations":["42ns","42s"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice without durations": log0.Durations()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"slice without durations":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of any durations": log0.Anys(d, d2)}
		}(),
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of any durations":["42ns","42s"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var d, d2 = 42 * time.Nanosecond, 42 * time.Second
			return map[string]json.Marshaler{"slice of reflect of durations": log0.Reflects(d, d2)}
		}(),
		expected:     "42ns 42s",
		expectedText: "42ns 42s",
		expectedJSON: `{
			"slice of reflect of durations":[42,42000000000]
		}`,
	},
}

func TestMarshalDurations(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalDurationsTestCases)
}
