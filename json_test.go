// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalJSONTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"kv slice": plog.JSON(plog.StringString("foo", "bar"), plog.StringInt("xyz", 42))},
		want:     `foo "bar" xyz 42`,
		wantText: `foo "bar" xyz 42`,
		wantJSON: `{
			"kv slice":{"foo":"bar","xyz":42}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without jsons": plog.JSON()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"without jsons":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty jsons": plog.JSON(plog.String(""), plog.String(""))},
		want:     ``,
		wantText: ``,
		wantJSON: `{
			"slice of empty jsons":{}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of json nils": plog.JSON(nil, nil)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"slice of json nils":{}
		}`,
	},
}

func TestMarshalJSON(t *testing.T) {
	testMarshal(t, MarshalJSONTests)
}
