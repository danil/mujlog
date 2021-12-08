// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalReflectTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"struct reflect": plog.Reflect(Struct{Name: "John Doe", Age: 42})},
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
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
			return map[string]json.Marshaler{"struct reflect pointer": plog.Reflect(&s)}
		}(),
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect byte array": plog.Reflect([3]byte{'f', 'o', 'o'})},
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"reflect byte array pointer": plog.Reflect(&a)}
		}(),
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"reflect byte array pointer to nil": plog.Reflect(a)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect untyped nil": plog.Reflect(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect untyped nil":null
		}`,
	},
}

func TestMarshalReflect(t *testing.T) {
	testMarshal(t, MarshalReflectTests)
}
