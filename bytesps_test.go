// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalBytespsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte("Hello, Wörld!"), []byte("Hello, World!")
			return map[string]json.Marshaler{"slice of byte slice pointers": log0.Bytesps(&p, &p2)}
		}(),
		expected:     "Hello, Wörld! Hello, World!",
		expectedText: "Hello, Wörld! Hello, World!",
		expectedJSON: `{
			"slice of byte slice pointers":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p, p2 := []byte{}, []byte{}
			return map[string]json.Marshaler{"slice of empty byte slice pointers": log0.Bytesps(&p, &p2)}
		}(),
		expected:     " ",
		expectedText: " ",
		expectedJSON: `{
			"slice of empty byte slice pointers":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of nil byte slice pointers": log0.Bytesps(nil, nil)},
		expected:     "null null",
		expectedText: "null null",
		expectedJSON: `{
			"slice of nil byte slice pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty slice of byte slice pointers": log0.Bytesps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"empty slice of byte slice pointers":null
		}`,
	},
}

func TestMarshalBytesps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBytespsTestCases)
}
