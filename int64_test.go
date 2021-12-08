// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt64Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int64": plog.Int64(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int64": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int64": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int64":42
		}`,
	},
}

func TestMarshalInt64(t *testing.T) {
	testMarshal(t, MarshalInt64Tests)
}
