// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint16pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"uint16 pointer": plog.Uint16p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint16 pointer": plog.Uint16p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"any uint16 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"reflect uint16 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i *uint16
			return map[string]json.Marshaler{"reflect uint16 pointer to nil": plog.Reflect(i)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
}

func TestMarshalUint16p(t *testing.T) {
	testMarshal(t, MarshalUint16pTests)
}
