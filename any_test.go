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

var MarshalAnyTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any struct": log0.Any(Struct{Name: "John Doe", Age: 42})},
		expected:     "{John Doe 42}",
		expectedText: "{John Doe 42}",
		expectedJSON: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			s := Struct{Name: "John Doe", Age: 42}
			return map[string]json.Marshaler{"any struct pointer": log0.Any(&s)}
		}(),
		expected:     "{John Doe 42}",
		expectedText: "{John Doe 42}",
		expectedJSON: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any byte array": log0.Any([3]byte{'f', 'o', 'o'})},
		expected:     "[102 111 111]",
		expectedText: "[102 111 111]",
		expectedJSON: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"any byte array pointer": log0.Any(&a)}
		}(),
		expected:     "[102 111 111]",
		expectedText: "[102 111 111]",
		expectedJSON: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"any byte array pointer to nil": log0.Any(a)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"any byte array pointer to nil":null
		}`,
	},
}

func TestMarshalAny(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalAnyTestCases)
}
