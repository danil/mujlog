// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint32Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint32": plog.Uint32(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint32": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint32": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint32":42
		}`,
	},
}

func TestMarshalUint32(t *testing.T) {
	testMarshal(t, MarshalUint32Tests)
}
