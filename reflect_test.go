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

var MarshalReflectTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"struct reflect": log0.Reflect(Struct{Name: "John Doe", Age: 42})},
		expected:     "{John Doe 42}",
		expectedText: "{John Doe 42}",
		expectedJSON: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			s := Struct{Name: "John Doe", Age: 42}
			return map[string]json.Marshaler{"struct reflect pointer": log0.Reflect(&s)}
		}(),
		expected:     "{John Doe 42}",
		expectedText: "{John Doe 42}",
		expectedJSON: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect byte array": log0.Reflect([3]byte{'f', 'o', 'o'})},
		expected:     "[102 111 111]",
		expectedText: "[102 111 111]",
		expectedJSON: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"reflect byte array pointer": log0.Reflect(&a)}
		}(),
		expected:     "[102 111 111]",
		expectedText: "[102 111 111]",
		expectedJSON: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"reflect byte array pointer to nil": log0.Reflect(a)}
		}(),
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect untyped nil": log0.Reflect(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"reflect untyped nil":null
		}`,
	},
}

func TestMarshalReflect(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalReflectTestCases)
}
