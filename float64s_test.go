// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat64sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float64 slice": plog.Float64s(0.123456789, 0.987654641)},
		want:     "0.123456789 0.987654641",
		wantText: "0.123456789 0.987654641",
		wantJSON: `{
			"float64 slice":[0.123456789,0.987654641]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without float64": plog.Float64s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without float64":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654641
			return map[string]json.Marshaler{"slice of any float64": plog.Anys(f, f2)}
		}(),
		want:     "0.123456789 0.987654641",
		wantText: "0.123456789 0.987654641",
		wantJSON: `{
			"slice of any float64":[0.123456789, 0.987654641]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654641
			return map[string]json.Marshaler{"slice of float64 reflects": plog.Reflects(f, f2)}
		}(),
		want:     "0.123456789 0.987654641",
		wantText: "0.123456789 0.987654641",
		wantJSON: `{
			"slice of float64 reflects":[0.123456789, 0.987654641]
		}`,
	},
}

func TestMarshalFloat64s(t *testing.T) {
	testMarshal(t, MarshalFloat64sTests)
}
