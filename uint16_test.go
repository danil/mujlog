// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint16Tests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint16": plog.Uint16(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint16": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint16": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint16":42
		}`,
	},
}

func TestMarshalUint16(t *testing.T) {
	testMarshal(t, MarshalUint16Tests)
}
