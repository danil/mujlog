// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint8sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint8 slice": plog.Uint8s(42, 77)},
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uint8 slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uint8": plog.Uint8s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without uint8":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of any uint8": plog.Anys(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uint8":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint8 = 42, 77
			return map[string]json.Marshaler{"slice of uint8 reflects": plog.Reflects(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of uint8 reflects":[42,77]
		}`,
	},
}

func TestMarshalUint8s(t *testing.T) {
	testMarshal(t, MarshalUint8sTests)
}
