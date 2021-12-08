// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalUint32sTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint32 slice": plog.Uint32s(42, 77)},
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"uint32 slice":[42,77]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice without uint32": plog.Uint32s()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice without uint32":[]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of any uint32": plog.Anys(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of any uint32":[42,77]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i, i2 uint32 = 42, 77
			return map[string]json.Marshaler{"slice of uint32 reflects": plog.Reflects(i, i2)}
		}(),
		want:     "42 77",
		wantText: "42 77",
		wantJSON: `{
			"slice of uint32 reflects":[42,77]
		}`,
	},
}

func TestMarshalUint32s(t *testing.T) {
	testMarshal(t, MarshalUint32sTests)
}
