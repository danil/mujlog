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

var MarshalComplex64pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"complex64 pointer": plog.Complex64p(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil complex64 pointer": plog.Complex64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"any complex64 pointer": plog.Any(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"reflect complex64 pointer": plog.Reflect(&c)}
		}(),
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
}

func TestMarshalComplex64p(t *testing.T) {
	testMarshal(t, MarshalComplex64pTests)
}
