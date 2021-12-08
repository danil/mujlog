// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalAnyTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any struct": plog.Any(Struct{Name: "John Doe", Age: 42})},
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
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
			return map[string]json.Marshaler{"any struct pointer": plog.Any(&s)}
		}(),
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any byte array": plog.Any([3]byte{'f', 'o', 'o'})},
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"any byte array pointer": plog.Any(&a)}
		}(),
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"any byte array pointer to nil": plog.Any(a)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"any byte array pointer to nil":null
		}`,
	},
}

func TestMarshalAny(t *testing.T) {
	testMarshal(t, MarshalAnyTests)
}
