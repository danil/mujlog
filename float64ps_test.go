// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat64psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"float64 pointer slice": plog.Float64ps(&f, &f2)}
		}(),
		want:     "0.123456789 0.987654321",
		wantText: "0.123456789 0.987654321",
		wantJSON: `{
			"float64 pointer slice":[0.123456789,0.987654321]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil float64 pointers": plog.Float64ps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil float64 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without float64 pointers": plog.Float64ps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without float64 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of any float64 pointers": plog.Anys(&f, &f2)}
		}(),
		want:     "0.123456789 0.987654321",
		wantText: "0.123456789 0.987654321",
		wantJSON: `{
			"slice of any float64 pointers":[0.123456789,0.987654321]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 float64 = 0.123456789, 0.987654321
			return map[string]json.Marshaler{"slice of reflects of float64 pointers": plog.Reflects(&f, &f2)}
		}(),
		want:     "0.123456789 0.987654321",
		wantText: "0.123456789 0.987654321",
		wantJSON: `{
			"slice of reflects of float64 pointers":[0.123456789,0.987654321]
		}`,
	},
}

func TestMarshalFloat64ps(t *testing.T) {
	testMarshal(t, MarshalFloat64psTests)
}
