// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat64pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"float64 pointer": plog.Float64p(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 0.123456789
			return map[string]json.Marshaler{"high precision float64 pointer": plog.Float64p(&f)}
		}(),
		want:     "0.123456789",
		wantText: "0.123456789",
		wantJSON: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float64 nil pointer": plog.Float64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"any float64 pointer": plog.Any(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"reflect float64 pointer": plog.Reflect(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float64
			return map[string]json.Marshaler{"reflect float64 pointer to nil": plog.Reflect(f)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect float64 pointer to nil":null
		}`,
	},
}

func TestMarshalFloat64p(t *testing.T) {
	testMarshal(t, MarshalFloat64pTests)
}
