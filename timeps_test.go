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

var MarshalTimepsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"time pointer slice": plog.Timeps(&t, &t2)}
		}(),
		want:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantJSON: `{
			"time pointer slice":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil time pointers": plog.Timeps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil time pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without time pointers": plog.Timeps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without time pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"slice of any time pointers": plog.Anys(&t, &t2)}
		}(),
		want:     "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantText: "1970-01-01T02:03:04.000000042Z 1970-12-05T04:03:02.000000001Z",
		wantJSON: `{
			"slice of any time pointers":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t, t2 time.Time = time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC)
			return map[string]json.Marshaler{"slice of reflects of time pointers": plog.Reflects(&t, &t2)}
		}(),
		want:     "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		wantText: "1970-01-01 02:03:04.000000042 +0000 UTC 1970-12-05 04:03:02.000000001 +0000 UTC",
		wantJSON: `{
			"slice of reflects of time pointers":["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]
		}`,
	},
}

func TestMarshalTimeps(t *testing.T) {
	testMarshal(t, MarshalTimepsTests)
}
