// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalRunesTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"runes": plog.Runes([]rune("Hello, Wörld!")...)},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty runes": plog.Runes([]rune{}...)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var p []rune
			return map[string]json.Marshaler{"nil runes": plog.Runes(p...)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil runes":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"rune slice with zero rune": plog.Runes([]rune{rune(0)}...)},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any runes": plog.Any([]rune("Hello, Wörld!"))},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any empty runes": plog.Any([]rune{})},
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty runes":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any rune slice with zero rune": plog.Any([]rune{rune(0)})},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect runes": plog.Reflect([]rune("Hello, Wörld!"))},
		want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantJSON: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect empty runes": plog.Reflect([]rune{})},
		want:     "[]",
		wantText: "[]",
		wantJSON: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect rune slice with zero rune": plog.Reflect([]rune{rune(0)})},
		want:     "[0]",
		wantText: "[0]",
		wantJSON: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
}

func TestMarshalRunes(t *testing.T) {
	testMarshal(t, MarshalRunesTests)
}
