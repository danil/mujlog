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

var MarshalRawsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of raw jsons": plog.Raws([]byte(`{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`), []byte("[42]"))},
		want:     `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		wantText: `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		wantJSON: `{
			"slice of raw jsons":[{"foo":{"bar":{"xyz":"Hello, Wörld!"}}},[42]]
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raws with quote": plog.Raws([]byte(`Hello, "Wörld"!`), []byte("[42]"))},
		want:      `Hello, "Wörld"! [42]`,
		wantText:  `Hello, "Wörld"! [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'H' looking for beginning of value"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"quoted raws": plog.Raws([]byte(`"Hello, Wörld!"`), []byte("[42]"))},
		want:     `"Hello, Wörld!" [42]`,
		wantText: `"Hello, Wörld!" [42]`,
		wantJSON: `{
			"quoted raws":["Hello, Wörld!",[42]]
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raws with nested quote": plog.Raws([]byte(`"Hello, "Wörld"!"`), []byte("[42]"))},
		want:      `"Hello, "Wörld"!" [42]`,
		wantText:  `"Hello, "Wörld"!" [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'W' after array element"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw quoted jsons": plog.Raws([]byte(`"{"foo":"bar"}"`), []byte("[42]"))},
		want:      `"{"foo":"bar"}" [42]`,
		wantText:  `"{"foo":"bar"}" [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'f' after array element"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"slice of raw malformed json objects": plog.Raws([]byte(`xyz{"foo":"bar"}`), []byte("[42]"))},
		want:      `xyz{"foo":"bar"} [42]`,
		wantText:  `xyz{"foo":"bar"} [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"slice of raw malformed json key/value": plog.Raws([]byte(`{"foo":"bar""}`), []byte("[42]"))},
		want:      `{"foo":"bar""} [42]`,
		wantText:  `{"foo":"bar""} [42]`,
		wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"slice of raw json with unescaped null byte": plog.Raws(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...), []byte("[42]"))},
		want:      "{\"foo\":\"\u0000xyz\"} [42]",
		wantText:  "{\"foo\":\"\u0000xyz\"} [42]",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"slice of empty raws": plog.Raws([]byte{}, []byte{})},
		want:      " ",
		wantText:  " ",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character ',' looking for beginning of value"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of raw nils": plog.Raws(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of raw nils":[null,null]
		}`,
	},
}

func TestMarshalRaws(t *testing.T) {
	testMarshal(t, MarshalRawsTests)
}
