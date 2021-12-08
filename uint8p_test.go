// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint8pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"uint8 pointer": plog.Uint8p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint8 pointer": plog.Uint8p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"any uint8 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"reflect uint8 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint8 pointer":42
		}`,
	},
}

func TestMarshalUint8p(t *testing.T) {
	testMarshal(t, MarshalUint8pTests)
}
