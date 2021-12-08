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

var MarshalComplex128pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"complex128 pointer": plog.Complex128p(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil complex128 pointer": plog.Complex128p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"any complex128 pointer": plog.Any(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"reflect complex128 pointer": plog.Reflect(&c)}
		}(),
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"complex64": plog.Complex64(complex(3, 21))},
		want:     "3+21i",
		wantText: "3+21i",
		wantJSON: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any complex64": plog.Any(complex(3, 21))},
		want:     "3+21i",
		wantText: "3+21i",
		wantJSON: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"reflect complex64": plog.Reflect(complex(3, 21))},
		want:      "(3+21i)",
		wantText:  "(3+21i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
}

func TestMarshalComplex128p(t *testing.T) {
	testMarshal(t, MarshalComplex128pTests)
}
