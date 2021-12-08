// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalFloat32Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"high precision float32": plog.Float32(0.123456789)},
		want:     "0.12345679",
		wantText: "0.12345679",
		wantJSON: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"zero float32": plog.Float32(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"zero float32":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any float32": plog.Any(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float32":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any zero float32": plog.Any(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"any zero float32":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect float32": plog.Reflect(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect zero float32": plog.Reflect(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"reflect zero float32":0
		}`,
	},
}

func TestMarshalFloat32(t *testing.T) {
	testMarshal(t, MarshalFloat32Tests)
}
