// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUintTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint pointer": plog.Uint(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"any uint": plog.Any(i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"reflect uint": plog.Reflect(i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint":42
		}`,
	},
}

func TestMarshalUint(t *testing.T) {
	testMarshal(t, MarshalUintTests)
}
