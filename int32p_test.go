// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt32pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"int32 pointer": plog.Int32p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"any int32 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"reflect int32 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int32 pointer":42
		}`,
	},
}

func TestMarshalInt32p(t *testing.T) {
	testMarshal(t, MarshalInt32pTests)
}
