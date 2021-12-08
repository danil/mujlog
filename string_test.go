// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalStringTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"string": plog.String("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty string": plog.String("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"string with zero byte": plog.String(string(byte(0)))},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any string": plog.Any("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any empty string": plog.Any("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any string with zero byte": plog.Any(string(byte(0)))},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect string": plog.Reflect("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect empty string": plog.Reflect("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"reflect empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect string with zero byte": plog.Reflect(string(byte(0)))},
		want:     "\u0000",
		wantText: "\u0000",
		wantJSON: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
}

func TestMarshalString(t *testing.T) {
	testMarshal(t, MarshalStringTests)
}
