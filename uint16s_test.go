// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint16sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint16 slice": plog.Uint16s(42, 77)},
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uint16 slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uint16": plog.Uint16s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without uint16":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint16 = 42, 77
			return map[string]json.Marshaler{"slice of any uint16": plog.Anys(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uint16":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint16 = 42, 77
			return map[string]json.Marshaler{"slice of uint16 reflects": plog.Reflects(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of uint16 reflects":[42,77]
		}`,
	},
}

func TestMarshalUint16s(t *testing.T) {
	testMarshal(t, MarshalUint16sTests)
}
