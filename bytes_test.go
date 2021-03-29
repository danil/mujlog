// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalBytesTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes": log0.Bytes([]byte("Hello, Wörld!"))},
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes with quote": log0.Bytes([]byte(`Hello, "World"!`))},
		expected:     `Hello, \"World\"!`,
		expectedText: `Hello, \"World\"!`,
		expectedJSON: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes quote": log0.Bytes([]byte(`"Hello, World!"`))},
		expected:     `\"Hello, World!\"`,
		expectedText: `\"Hello, World!\"`,
		expectedJSON: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes nested quote": log0.Bytes([]byte(`"Hello, "World"!"`))},
		expected:     `\"Hello, \"World\"!\"`,
		expectedText: `\"Hello, \"World\"!\"`,
		expectedJSON: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes json": log0.Bytes([]byte(`{"foo":"bar"}`))},
		expected:     `{\"foo\":\"bar\"}`,
		expectedText: `{\"foo\":\"bar\"}`,
		expectedJSON: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"bytes json quote": log0.Bytes([]byte(`"{"foo":"bar"}"`))},
		expected:     `\"{\"foo\":\"bar\"}\"`,
		expectedText: `\"{\"foo\":\"bar\"}\"`,
		expectedJSON: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"empty bytes": log0.Bytes([]byte{})},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"empty bytes":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil bytes": log0.Bytes(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil bytes":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any bytes": log0.Any([]byte("Hello, Wörld!"))},
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any empty bytes": log0.Any([]byte{})},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"any empty bytes":""
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect bytes": log0.Reflect([]byte("Hello, Wörld!"))},
		expected:     "SGVsbG8sIFfDtnJsZCE=",
		expectedText: "SGVsbG8sIFfDtnJsZCE=",
		expectedJSON: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect empty bytes": log0.Reflect([]byte{})},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"any bytes pointer": log0.Any(&p)}
		}(),
		expected:     "Hello, Wörld!",
		expectedText: "Hello, Wörld!",
		expectedJSON: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"any empty bytes pointer": log0.Any(&p)}
		}(),
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect bytes pointer": log0.Reflect(&p)}
		}(),
		expected:     "SGVsbG8sIFfDtnJsZCE=",
		expectedText: "SGVsbG8sIFfDtnJsZCE=",
		expectedJSON: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"reflect empty bytes pointer": log0.Reflect(&p)}
		}(),
		expected: "",
		expectedJSON: `{
			"reflect empty bytes pointer":""
		}`,
	},
}

func TestMarshalBytes(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalBytesTestCases)
}
