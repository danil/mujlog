// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint32psTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint32 = 42, 77
			return map[string]json.Marshaler{"uint32 pointer slice": plog.Uint32ps(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uint32 pointer slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of nil uint32 pointers": plog.Uint32ps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"slice of nil uint32 pointers":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uint32 pointers": plog.Uint32ps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"slice without uint32 pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of any uint32 pointers": plog.Anys(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uint32 pointers":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f, f2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of reflects of uint32 pointers": plog.Reflects(&f, &f2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of reflects of uint32 pointers":[42,77]
		}`,
	},
}

func TestMarshalUint32ps(t *testing.T) {
	testMarshal(t, MarshalUint32psTests)
}
