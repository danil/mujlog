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

var MarshalDurationTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"duration": log0.Duration(42 * time.Nanosecond)},
		expected:     "42ns",
		expectedText: "42ns",
		expectedJSON: `{
			"duration":"42ns"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any duration": log0.Any(42 * time.Nanosecond)},
		expected:     "42ns",
		expectedText: "42ns",
		expectedJSON: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect duration": log0.Reflect(42 * time.Nanosecond)},
		expected:     "42ns",
		expectedText: "42ns",
		expectedJSON: `{
			"reflect duration":42
		}`,
	},
}

func TestMarshalDuration(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalDurationTestCases)
}
