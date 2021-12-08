// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/plog"
)

var MarshalRawTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of raw jsons": plog.Raw([]byte(`{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`))},
		want:     `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`,
		wantText: `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`,
		wantJSON: `{
			"slice of raw jsons":{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw with quote": plog.Raw([]byte(`Hello, "Wörld"!`))},
		want:      `Hello, "Wörld"!`,
		wantText:  `Hello, "Wörld"!`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'H' looking for beginning of value"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"quoted raw": plog.Raw([]byte(`"Hello, Wörld!"`))},
		want:     `"Hello, Wörld!"`,
		wantText: `"Hello, Wörld!"`,
		wantJSON: `{
			"quoted raw":"Hello, Wörld!"
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw with nested quote": plog.Raw([]byte(`"Hello, "Wörld"!"`))},
		want:      `"Hello, "Wörld"!"`,
		wantText:  `"Hello, "Wörld"!"`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'W' after top-level value"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw quoted json": plog.Raw([]byte(`"{"foo":"bar"}"`))},
		want:      `"{"foo":"bar"}"`,
		wantText:  `"{"foo":"bar"}"`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'f' after top-level value"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw malformed json object": plog.Raw([]byte(`xyz{"foo":"bar"}`))},
		want:      `xyz{"foo":"bar"}`,
		wantText:  `xyz{"foo":"bar"}`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw malformed json key/value": plog.Raw([]byte(`{"foo":"bar""}`))},
		want:      `{"foo":"bar""}`,
		wantText:  `{"foo":"bar""}`,
		wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw json with unescaped null byte": plog.Raw(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...))},
		want:      "{\"foo\":\"\u0000xyz\"}",
		wantText:  "{\"foo\":\"\u0000xyz\"}",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"raw nil": plog.Raw(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"raw nil":null
		}`,
	},
}

func TestMarshalRaw(t *testing.T) {
	testMarshal(t, MarshalRawTests)
}
