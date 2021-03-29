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

var MarshalBoolpTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"bool pointer to true": log0.Boolp(&b)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := false
			return map[string]json.Marshaler{"bool pointer to false": log0.Boolp(&b)}
		}(),
		expected:     "false",
		expectedText: "false",
		expectedJSON: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bool nil pointer": log0.Boolp(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"any bool pointer to true": log0.Any(&b)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"any twice/nested pointer to bool true": log0.Any(&b2)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"reflect bool pointer to true": log0.Reflect(&b)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"reflect bool twice/nested pointer to true": log0.Reflect(&b2)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var b *bool
			return map[string]json.Marshaler{"reflect bool pointer to nil": log0.Reflect(b)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"any bool pointer to true": log0.Any(&b)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"any twice/nested pointer to bool true": log0.Any(&b2)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"reflect bool pointer to true": log0.Reflect(&b)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"reflect bool twice/nested pointer to true": log0.Reflect(&b2)}
		}(),
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var b *bool
			return map[string]json.Marshaler{"reflect bool pointer to nil": log0.Reflect(b)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect bool pointer to nil":null
		}`,
	},
}

func TestMarshalBoolp(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBoolpTestCases)
}
