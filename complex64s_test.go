// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalComplex64sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of complex64s": plog.Complex64s(complex(1, 23), complex(3, 21))},
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"slice of complex64s":["1+23i","3+21i"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without complex64s": plog.Complex64s()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without complex64s":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of any complex64s": plog.Anys(c, c2)}
		}(),
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"slice of any complex64s":["1+23i","3+21i"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c, c2 complex64 = complex(1, 23), complex(3, 21)
			return map[string]json.Marshaler{"slice of reflect complex64": plog.Anys(c, c2)}
		}(),
		want:     "1+23i 3+21i",
		wantText: "1+23i 3+21i",
		wantJSON: `{
			"slice of reflect complex64":["1+23i","3+21i"]
		}`,
	},
}

func TestMarshalComplex64s(t *testing.T) {
	testMarshal(t, MarshalComplex64sTests)
}
