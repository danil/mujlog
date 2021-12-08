// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalRunespTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"runes pointer": plog.Runesp(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"empty runes pointer": plog.Runesp(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil runes pointer": plog.Runesp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"any runes pointer": plog.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"any empty runes pointer": plog.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect runes pointer": plog.Reflect(&p)}
		}(),
		want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantJSON: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"reflect empty runes pointer": plog.Reflect(&p)}
		}(),
		want:     "[]",
		wantText: "[]",
		wantJSON: `{
			"reflect empty runes pointer":[]
		}`,
	},
}

func TestMarshalRunesp(t *testing.T) {
	testMarshal(t, MarshalRunespTests)
}
