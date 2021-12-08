// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalBoolTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bool true": plog.Bool(true)},
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"bool true":true
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bool false": plog.Bool(false)},
		want:     "false",
		wantText: "false",
		wantJSON: `{
			"bool false":false
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any bool false": plog.Any(false)},
		want:     "false",
		wantText: "false",
		wantJSON: `{
			"any bool false":false
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect bool false": plog.Reflect(false)},
		want:     "false",
		wantText: "false",
		wantJSON: `{
			"reflect bool false":false
		}`,
	},
}

func TestMarshalBool(t *testing.T) {
	testMarshal(t, MarshalBoolTests)
}
