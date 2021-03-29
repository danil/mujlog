// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"

	"github.com/danil/log0"
)

var MarshalTimepTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"time pointer": &t}
		}(),
		expected:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		expectedText: "1970-01-01T00:00:00.000000042Z",
		expectedJSON: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t time.Time
			return map[string]json.Marshaler{"nil time pointer": t}
		}(),
		expected:     "0001-01-01 00:00:00 +0000 UTC",
		expectedText: "0001-01-01T00:00:00Z",
		expectedJSON: `{
			"nil time pointer":"0001-01-01T00:00:00Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"any time pointer": log0.Any(&t)}
		}(),
		expected:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
		expectedText: `1970-01-01T00:00:00.000000042Z`,
		expectedJSON: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"reflect time pointer": log0.Reflect(&t)}
		}(),
		expected:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		expectedText: "1970-01-01 00:00:00.000000042 +0000 UTC",
		expectedJSON: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
}

func TestMarshalTimep(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTimepTestCases)
}
