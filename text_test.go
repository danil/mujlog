// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalTestTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"text": log0.Text(log0.String("Hello, Wörld!"))},
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"text":"Hello, Wörld!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"text json": log0.Text(log0.String(`{"foo":"bar"}`))},
		expected:     `{\"foo\":\"bar\"}`,
		expectedText: `{\"foo\":\"bar\"}`,
		expectedJSON: `{
			"text json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"text with zero byte": log0.Text(log0.String("Hello, Wörld!\x00"))},
		expected:     "Hello, Wörld!\\u0000",
		expectedText: "Hello, Wörld!\\u0000",
		expectedJSON: `{
			"text with zero byte":"Hello, Wörld!\u0000"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty text": log0.Text(log0.String(""))},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"empty text":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"text nil": log0.Text(nil)},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"text nil":null
		}`,
	},
}

func TestMarshalText(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTestTestCases)
}
