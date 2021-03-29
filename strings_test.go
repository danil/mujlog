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

var MarshalStringsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"strings": log0.Strings("Hello, Wörld!", "Hello, World!")},
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty strings": log0.Strings("", "")},
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"empty strings":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"strings with zero byte": log0.Strings(string(byte(0)), string(byte(0)))},
		expected:     "\\u0000 \\u0000",
		expectedText: "\\u0000 \\u0000",
		expectedJSON: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without strings": log0.Strings()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"without strings":null
		}`,
	},
}

func TestMarshalStrings(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalStringsTestCases)
}
