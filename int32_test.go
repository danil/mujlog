// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalInt32Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int32": plog.Int32(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int32": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int32": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int32":42
		}`,
	},
}

func TestMarshalInt32(t *testing.T) {
	testMarshal(t, MarshalInt32Tests)
}
