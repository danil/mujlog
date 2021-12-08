// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat32pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"float32 pointer": plog.Float32p(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 0.123456789
			return map[string]json.Marshaler{"high precision float32 pointer": plog.Float32p(&f)}
		}(),
		want:     "0.12345679",
		wantText: "0.12345679",
		wantJSON: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float32 nil pointer": plog.Float32p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"any float32 pointer": plog.Any(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"reflect float32 pointer": plog.Reflect(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float32
			return map[string]json.Marshaler{"reflect float32 pointer to nil": plog.Reflect(f)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect float32 pointer to nil":null
		}`,
	},
}

func TestMarshalFloat32p(t *testing.T) {
	testMarshal(t, MarshalFloat32pTests)
}
