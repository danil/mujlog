// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalRawsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of raw jsons": log0.Raws([]byte(`{"foo":{"bar":{"xyz":"Hello, Wörld!"}}}`), []byte("[42]"))},
		expected:     `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		expectedText: `{"foo":{"bar":{"xyz":"Hello, Wörld!"}}} [42]`,
		expectedJSON: `{
			"slice of raw jsons":[{"foo":{"bar":{"xyz":"Hello, Wörld!"}}},[42]]
		}`,
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"raws with quote": log0.Raws([]byte(`Hello, "Wörld"!`), []byte("[42]"))},
		expected:      `Hello, "Wörld"! [42]`,
		expectedText:  `Hello, "Wörld"! [42]`,
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'H' looking for beginning of value"),
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"quoted raws": log0.Raws([]byte(`"Hello, Wörld!"`), []byte("[42]"))},
		expected:     `"Hello, Wörld!" [42]`,
		expectedText: `"Hello, Wörld!" [42]`,
		expectedJSON: `{
			"quoted raws":["Hello, Wörld!",[42]]
		}`,
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"raws with nested quote": log0.Raws([]byte(`"Hello, "Wörld"!"`), []byte("[42]"))},
		expected:      `"Hello, "Wörld"!" [42]`,
		expectedText:  `"Hello, "Wörld"!" [42]`,
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'W' after array element"),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"raw quoted jsons": log0.Raws([]byte(`"{"foo":"bar"}"`), []byte("[42]"))},
		expected:      `"{"foo":"bar"}" [42]`,
		expectedText:  `"{"foo":"bar"}" [42]`,
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'f' after array element"),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"slice of raw malformed json objects": log0.Raws([]byte(`xyz{"foo":"bar"}`), []byte("[42]"))},
		expected:      `xyz{"foo":"bar"} [42]`,
		expectedText:  `xyz{"foo":"bar"} [42]`,
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"slice of raw malformed json key/value": log0.Raws([]byte(`{"foo":"bar""}`), []byte("[42]"))},
		expected:      `{"foo":"bar""} [42]`,
		expectedText:  `{"foo":"bar""} [42]`,
		expectedError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"slice of raw json with unescaped null byte": log0.Raws(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...), []byte("[42]"))},
		expected:      "{\"foo\":\"\u0000xyz\"} [42]",
		expectedText:  "{\"foo\":\"\u0000xyz\"} [42]",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:          line(),
		input:         map[string]json.Marshaler{"slice of empty raws": log0.Raws([]byte{}, []byte{})},
		expected:      " ",
		expectedText:  " ",
		expectedError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character ',' looking for beginning of value"),
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of raw nils": log0.Raws(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of raw nils":[null,null]
		}`,
	},
}

func TestMarshalRaws(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalRawsTestCases)
}
