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

var MarshalBytessTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of byte slices": log0.Bytess([]byte("Hello, Wörld!"), []byte("Hello, World!"))},
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"slice of byte slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of byte slices with quote": log0.Bytess([]byte(`Hello, "Wörld"!`), []byte(`Hello, "World"!`))},
		expected:     `Hello, \"Wörld\"! Hello, \"World\"!`,
		expectedText: `Hello, \"Wörld\"! Hello, \"World\"!`,
		expectedJSON: `{
			"slice of byte slices with quote":["Hello, \"Wörld\"!","Hello, \"World\"!"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"quoted slice of byte slices": log0.Bytess([]byte(`"Hello, Wörld!"`), []byte(`"Hello, World!"`))},
		expected:     `\"Hello, Wörld!\" \"Hello, World!\"`,
		expectedText: `\"Hello, Wörld!\" \"Hello, World!\"`,
		expectedJSON: `{
			"quoted slice of byte slices":["\"Hello, Wörld!\"","\"Hello, World!\""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of byte slices with nested quote": log0.Bytess([]byte(`"Hello, "Wörld"!"`), []byte(`"Hello, "World"!"`))},
		expected:     `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		expectedText: `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		expectedJSON: `{
			"slice of byte slices with nested quote":["\"Hello, \"Wörld\"!\"","\"Hello, \"World\"!\""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of byte slices with json": log0.Bytess([]byte(`{"foo":"bar"}`), []byte(`{"baz":"xyz"}`))},
		expected:     `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		expectedText: `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		expectedJSON: `{
			"slice of byte slices with json":["{\"foo\":\"bar\"}","{\"baz\":\"xyz\"}"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of byte slices with quoted json": log0.Bytess([]byte(`"{"foo":"bar"}"`), []byte(`"{"baz":"xyz"}"`))},
		expected:     `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		expectedText: `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		expectedJSON: `{
			"slice of byte slices with quoted json":["\"{\"foo\":\"bar\"}\"","\"{\"baz\":\"xyz\"}\""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of empty byte slice": log0.Bytess([]byte{}, []byte{})},
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"slice of empty byte slice":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil byte slice": log0.Bytess(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil byte slice":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty slice of byte slices": log0.Bytess()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"empty slice of byte slices":null
		}`,
	},
}

func TestMarshalBytess(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBytessTestCases)
}
