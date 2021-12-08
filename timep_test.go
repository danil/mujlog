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

var MarshalTimepTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"time pointer": &t}
		}(),
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01T00:00:00.000000042Z",
		wantJSON: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t time.Time
			return map[string]json.Marshaler{"nil time pointer": t}
		}(),
		want:     "0001-01-01 00:00:00 +0000 UTC",
		wantText: "0001-01-01T00:00:00Z",
		wantJSON: `{
			"nil time pointer":"0001-01-01T00:00:00Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"any time pointer": plog.Any(&t)}
		}(),
		want:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
		wantText: `1970-01-01T00:00:00.000000042Z`,
		wantJSON: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"reflect time pointer": plog.Reflect(&t)}
		}(),
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantJSON: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
}

func TestMarshalTimep(t *testing.T) {
	testMarshal(t, MarshalTimepTests)
}
