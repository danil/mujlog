// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalBoolpsTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b, b2 := true, false
			return map[string]json.Marshaler{"bool pointers to true and false": plog.Boolps(&b, &b2)}
		}(),
		want:     "true false",
		wantText: "true false",
		wantJSON: `{
			"bool pointers to true and false":[true,false]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bool pointers to nil": plog.Boolps(nil, nil)},
		want:     "null null",
		wantText: "null null",
		wantJSON: `{
			"bool pointers to nil":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without bool pointers": plog.Boolps()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"without bool pointers":null
		}`,
	},
}

func TestMarshalBoolps(t *testing.T) {
	testMarshal(t, MarshalBoolpsTests)
}
