// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint64pTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"uint64 pointer": plog.Uint64p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint64 pointer": plog.Uint64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"any uint64 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"reflect uint64 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint64 pointer":42
		}`,
	},
}

func TestMarshalUint64p(t *testing.T) {
	testMarshal(t, MarshalUint64pTests)
}
