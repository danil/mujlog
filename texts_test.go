// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalTextsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"texts": plog.Texts(plog.String("Hello, Wörld!"), plog.String("Hello, World!"))},
		want:     `Hello, Wörld! Hello, World!`,
		wantText: `Hello, Wörld! Hello, World!`,
		wantJSON: `{
			"texts":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of text jsons": plog.Texts(plog.String(`{"foo":"bar"}`), plog.String("[42]"))},
		want:     `{\"foo\":\"bar\"} [42]`,
		wantText: `{\"foo\":\"bar\"} [42]`,
		wantJSON: `{
			"slice of text jsons":["{\"foo\":\"bar\"}","[42]"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of texts with unescaped null byte": plog.Texts(plog.String("Hello, Wörld!\x00"), plog.String("Hello, World!"))},
		want:     "Hello, Wörld!\\u0000 Hello, World!",
		wantText: "Hello, Wörld!\\u0000 Hello, World!",
		wantJSON: `{
			"slice of texts with unescaped null byte":["Hello, Wörld!\u0000","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty texts": plog.Texts(plog.String(""), plog.String(""))},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of empty texts":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of text nils": plog.Texts(nil, nil)},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"slice of text nils":[null,null]
		}`,
	},
}

func TestMarshalTexts(t *testing.T) {
	testMarshal(t, MarshalTextsTests)
}
