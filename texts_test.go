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

var MarshalTextsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"texts": log0.Texts(log0.String("Hello, Wörld!"), log0.String("Hello, World!"))},
		expected:     `Hello, Wörld! Hello, World!`,
		expectedText: `Hello, Wörld! Hello, World!`,
		expectedJSON: `{
			"texts":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of text jsons": log0.Texts(log0.String(`{"foo":"bar"}`), log0.String("[42]"))},
		expected:     `{\"foo\":\"bar\"} [42]`,
		expectedText: `{\"foo\":\"bar\"} [42]`,
		expectedJSON: `{
			"slice of text jsons":["{\"foo\":\"bar\"}","[42]"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of texts with unescaped null byte": log0.Texts(log0.String("Hello, Wörld!\x00"), log0.String("Hello, World!"))},
		expected:     "Hello, Wörld!\\u0000 Hello, World!",
		expectedText: "Hello, Wörld!\\u0000 Hello, World!",
		expectedJSON: `{
			"slice of texts with unescaped null byte":["Hello, Wörld!\u0000","Hello, World!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of empty texts": log0.Texts(log0.String(""), log0.String(""))},
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"slice of empty texts":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of text nils": log0.Texts(nil, nil)},
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"slice of text nils":[null,null]
		}`,
	},
}

func TestMarshalTexts(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalTextsTestCases)
}
