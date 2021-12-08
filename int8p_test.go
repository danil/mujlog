// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt8pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"int8 pointer": plog.Int8p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"any int8 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"reflect int8 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int8 pointer":42
		}`,
	},
}

func TestMarshalInt8p(t *testing.T) {
	testMarshal(t, MarshalInt8pTests)
}
