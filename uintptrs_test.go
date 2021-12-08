// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUintptrsTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uintptr slice": plog.Uintptrs(42, 77)},
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uintptr slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uintptr": plog.Uintptrs()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without uintptr":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of any uintptr": plog.Anys(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uintptr":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uintptr = 42, 77
			return map[string]json.Marshaler{"slice of uintptr reflects": plog.Reflects(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of uintptr reflects":[42,77]
		}`,
	},
}

func TestMarshalUintptrs(t *testing.T) {
	testMarshal(t, MarshalUintptrsTests)
}
