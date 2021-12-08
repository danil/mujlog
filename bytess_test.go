// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

// TODO: Add any/relfect test cases.
var MarshalBytessTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices": plog.Bytess([]byte("Hello, Wörld!"), []byte("Hello, World!"))},
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"slice of byte slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with quote": plog.Bytess([]byte(`Hello, "Wörld"!`), []byte(`Hello, "World"!`))},
		want:     `Hello, \"Wörld\"! Hello, \"World\"!`,
		wantText: `Hello, \"Wörld\"! Hello, \"World\"!`,
		wantJSON: `{
			"slice of byte slices with quote":["Hello, \"Wörld\"!","Hello, \"World\"!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"quoted slice of byte slices": plog.Bytess([]byte(`"Hello, Wörld!"`), []byte(`"Hello, World!"`))},
		want:     `\"Hello, Wörld!\" \"Hello, World!\"`,
		wantText: `\"Hello, Wörld!\" \"Hello, World!\"`,
		wantJSON: `{
			"quoted slice of byte slices":["\"Hello, Wörld!\"","\"Hello, World!\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with nested quote": plog.Bytess([]byte(`"Hello, "Wörld"!"`), []byte(`"Hello, "World"!"`))},
		want:     `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		wantText: `\"Hello, \"Wörld\"!\" \"Hello, \"World\"!\"`,
		wantJSON: `{
			"slice of byte slices with nested quote":["\"Hello, \"Wörld\"!\"","\"Hello, \"World\"!\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with json": plog.Bytess([]byte(`{"foo":"bar"}`), []byte(`{"baz":"xyz"}`))},
		want:     `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		wantText: `{\"foo\":\"bar\"} {\"baz\":\"xyz\"}`,
		wantJSON: `{
			"slice of byte slices with json":["{\"foo\":\"bar\"}","{\"baz\":\"xyz\"}"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of byte slices with quoted json": plog.Bytess([]byte(`"{"foo":"bar"}"`), []byte(`"{"baz":"xyz"}"`))},
		want:     `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		wantText: `\"{\"foo\":\"bar\"}\" \"{\"baz\":\"xyz\"}\"`,
		wantJSON: `{
			"slice of byte slices with quoted json":["\"{\"foo\":\"bar\"}\"","\"{\"baz\":\"xyz\"}\""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty byte slice": plog.Bytess([]byte{}, []byte{})},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of empty byte slice":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil byte slice": plog.Bytess(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil byte slice":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty slice of byte slices": plog.Bytess()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"empty slice of byte slices":null
		}`,
	},
}

func TestMarshalBytess(t *testing.T) {
	testMarshal(t, MarshalBytessTests)
}
