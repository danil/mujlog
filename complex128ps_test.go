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

var MarshalComplex128psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"complex128 pointers slice": plog.Complex128ps(&c, &c2)}
		}(),
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"complex128 pointers slice":["1+23i","3+21i"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil complex128 pointers": plog.Complex128ps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil complex128 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without complex128 pointers": plog.Complex128ps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without complex128 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex128 pointers": plog.Anys(&c, &c2)}
		}(),
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"slice of any complex128 pointers":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflects of complex128 pointers": plog.Reflects(&c, &c2)}
		}(),
		want:      "(1+23i) (3+21i)",
		wantText:  "(1+23i) (3+21i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128ps(t *testing.T) {
	testMarshal(t, MarshalComplex128psTests)
}
