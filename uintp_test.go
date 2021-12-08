// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUintpTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"uint pointer": plog.Uintp(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint pointer": plog.Uintp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"any uint pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"reflect uint pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint pointer":42
		}`,
	},
}

func TestMarshalUintp(t *testing.T) {
	testMarshal(t, MarshalUintpTests)
}
