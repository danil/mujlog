// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pprint/plog"
)

var MarshalDurationpTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"duration pointer": plog.Durationp(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil duration pointer": plog.Durationp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"any duration pointer": plog.Any(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"reflect duration pointer": plog.Reflect(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"reflect duration pointer":42
		}`,
	},
}

func TestDurationpMarshal(t *testing.T) {
	testMarshal(t, MarshalDurationpTests)
}
