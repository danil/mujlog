// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt8sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int8 slice": plog.Int8s(42, 77)},
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"int8 slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int8": plog.Int8s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without int8":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of any int8": plog.Anys(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any int8":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int8 = 42, 77
			return map[string]json.Marshaler{"slice of int8 reflects": plog.Reflects(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of int8 reflects":[42,77]
		}`,
	},
}

func TestMarshalInt8s(t *testing.T) {
	testMarshal(t, MarshalInt8sTests)
}
