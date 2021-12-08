// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt64pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"int64 pointer": plog.Int64p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"any int64 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"reflect int64 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int64 pointer":42
		}`,
	},
}

func TestMarshalInt64p(t *testing.T) {
	testMarshal(t, MarshalInt64pTests)
}
