// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt32psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"int32 pointer slice": plog.Int32ps(&f, &f2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int32 pointer slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil int32 pointers": plog.Int32ps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil int32 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int32 pointers": plog.Int32ps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without int32 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of any int32 pointers": plog.Anys(&f, &f2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int32 pointers":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 int32 = 123, 321
			return map[string]json.Marshaler{"slice of reflects of int32 pointers": plog.Reflects(&f, &f2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of reflects of int32 pointers":[123,321]
		}`,
	},
}

func TestMarshalInt32ps(t *testing.T) {
	testMarshal(t, MarshalInt32psTests)
}
