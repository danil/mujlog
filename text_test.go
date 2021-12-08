// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalTestTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text": plog.Text(plog.String("Hello, Wörld!"))},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"text":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text json": plog.Text(plog.String(`{"foo":"bar"}`))},
		want:     `{\"foo\":\"bar\"}`,
		wantText: `{\"foo\":\"bar\"}`,
		wantJSON: `{
			"text json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text with zero byte": plog.Text(plog.String("Hello, Wörld!\x00"))},
		want:     "Hello, Wörld!\\u0000",
		wantText: "Hello, Wörld!\\u0000",
		wantJSON: `{
			"text with zero byte":"Hello, Wörld!\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty text": plog.Text(plog.String(""))},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty text":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text nil": plog.Text(nil)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"text nil":null
		}`,
	},
}

func TestMarshalText(t *testing.T) {
	testMarshal(t, MarshalTestTests)
}
