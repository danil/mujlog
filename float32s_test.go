// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat32sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float32 slice": plog.Float32s(0.123456789, 0.987654321)},
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"float32 slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without float32": plog.Float32s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without float32":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of any float32": plog.Anys(f, f2)}
		}(),
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"slice of any float32":[0.123456789, 0.987654321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float32 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of float32 reflects": plog.Reflects(f, f2)}
		}(),
		want:     "0.12345679 0.9876543",
		wantText: "0.12345679 0.9876543",
		wantJSON: `{
			"slice of float32 reflects":[0.123456789, 0.987654321]
		}`,
	},
}

func TestMarshalFloat32s(t *testing.T) {
	testMarshal(t, MarshalFloat32sTests)
}
