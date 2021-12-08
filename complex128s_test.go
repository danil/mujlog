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

var MarshalComplex128sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"complex128 slice": plog.Complex128s(complex(1, 23), complex(3, 21))},
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"complex128 slice":["1+23i","3+21i"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without complex128": plog.Complex128s()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without complex128":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex128": plog.Anys(c, c2)}
		}(),
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"slice of any complex128":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex128 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflect of complex128 pointers": plog.Reflects(c, c2)}
		}(),
		want:      "(1+23i) (3+21i)",
		wantText:  "(1+23i) (3+21i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128s(t *testing.T) {
	testMarshal(t, MarshalComplex128sTests)
}
