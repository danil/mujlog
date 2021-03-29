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

var MarshalRawTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"raw malformed json object": log0.Raw([]byte(`xyz{"foo":"bar"}`))},
		expected:     `xyz{"foo":"bar"}`,
		expectedText: `xyz{"foo":"bar"}`,
		error:        errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"raw malformed json key/value": log0.Raw([]byte(`{"foo":"bar""}`))},
		expected:     `{"foo":"bar""}`,
		expectedText: `{"foo":"bar""}`,
		error:        errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"raw json with unescaped null byte": log0.Raw(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...))},
		expected:     "{\"foo\":\"\u0000xyz\"}",
		expectedText: "{\"foo\":\"\u0000xyz\"}",
		error:        errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"raw nil": log0.Raw(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"raw nil":null
		}`,
	},
}

func TestMarshalRaw(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalRawTestCases)
}
