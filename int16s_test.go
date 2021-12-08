// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt16sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int16 slice": plog.Int16s(123, 321)},
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"int16 slice":[123,321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without int16": plog.Int16s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without int16":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of any int16": plog.Anys(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of any int16":[123,321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 int16 = 123, 321
			return map[string]json.Marshaler{"slice of int16 reflects": plog.Reflects(i, i2)}
		}(),
		want:     "123 321",
		wantText: "123 321",
		wantJSON: `{
			"slice of int16 reflects":[123,321]
		}`,
	},
}

func TestMarshalInt16s(t *testing.T) {
	testMarshal(t, MarshalInt16sTests)
}
