// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt16pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"int16 pointer": plog.Int16p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"any int16 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"reflect int16 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int16 pointer":42
		}`,
	},
}

func TestInt16pMarshal(t *testing.T) {
	testMarshal(t, MarshalInt16pTests)
}
