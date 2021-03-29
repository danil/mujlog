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

var MarshalBoolTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bool true": log0.Bool(true)},
		expected:     "true",
		expectedText: "true",
		expectedJSON: `{
			"bool true":true
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bool false": log0.Bool(false)},
		expected:     "false",
		expectedText: "false",
		expectedJSON: `{
			"bool false":false
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any bool false": log0.Any(false)},
		expected:     "false",
		expectedText: "false",
		expectedJSON: `{
			"any bool false":false
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect bool false": log0.Reflect(false)},
		expected:     "false",
		expectedText: "false",
		expectedJSON: `{
			"reflect bool false":false
		}`,
	},
}

func TestMarshalBool(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBoolTestCases)
}
