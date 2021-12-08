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

var MarshalRawpsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`), []byte("[42]")
			return map[string]json.Marshaler{"slice of raw jsons": plog.Rawps(&p, &p2)}
		}(),
		want:     `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		wantText: `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		wantJSON: `{
			"slice of raw jsons":[{"foo":{"bar":{"xyz":"Hello, Wörld!"}}},[42]]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`Hello, "Wörld"!`), []byte("[42]")
			return map[string]json.Marshaler{"rawps with quote": plog.Rawps(&p, &p2)}
		}(),
		want:      `Hello, "Wörld"! [42]`,
		wantText:  `Hello, "Wörld"! [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'H' looking for beginning of value"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`"Hello, Wörld!"`), []byte("[42]")
			return map[string]json.Marshaler{"quoted rawps": plog.Rawps(&p, &p2)}
		}(),
		want:     `"Hello, Wörld!" [42]`,
		wantText: `"Hello, Wörld!" [42]`,
		wantJSON: `{
			"quoted rawps":["Hello, Wörld!",[42]]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`"Hello, "Wörld"!"`), []byte("[42]")
			return map[string]json.Marshaler{"rawps with nested quote": plog.Rawps(&p, &p2)}
		}(),
		want:      `"Hello, "Wörld"!" [42]`,
		wantText:  `"Hello, "Wörld"!" [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'W' after array element"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`"{"foo":"bar"}"`), []byte("[42]")
			return map[string]json.Marshaler{"raw quoted jsons": plog.Rawps(&p, &p2)}
		}(),
		want:      `"{"foo":"bar"}" [42]`,
		wantText:  `"{"foo":"bar"}" [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'f' after array element"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`xyz{"foo":"bar"}`), []byte("[42]")
			return map[string]json.Marshaler{"slice of raw malformed json objects": plog.Rawps(&p, &p2)}
		}(),
		want:      `xyz{"foo":"bar"} [42]`,
		wantText:  `xyz{"foo":"bar"} [42]`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte(`{"foo":"bar""}`), []byte("[42]")
			return map[string]json.Marshaler{"slice of raw malformed json key/value": plog.Rawps(&p, &p2)}
		}(),
		want:      `{"foo":"bar""} [42]`,
		wantText:  `{"foo":"bar""} [42]`,
		wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...), []byte("[42]")
			return map[string]json.Marshaler{"slice of raw json with unescaped null byte": plog.Rawps(&p, &p2)}
		}(),
		want:      "{\"foo\":\"\u0000xyz\"} [42]",
		wantText:  "{\"foo\":\"\u0000xyz\"} [42]",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte{}, []byte{}
			return map[string]json.Marshaler{"slice of empty rawps": plog.Rawps(&p, &p2)}
		}(),
		want:      " ",
		wantText:  " ",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character ',' looking for beginning of value"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var p, p2 []byte
			return map[string]json.Marshaler{"slice of raw nils": plog.Rawps(&p, &p2)}
		}(),
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of raw nils":[null,null]
		}`,
	},
}

func TestMarshalRawps(t *testing.T) {
	testMarshal(t, MarshalRawpsTests)
}
