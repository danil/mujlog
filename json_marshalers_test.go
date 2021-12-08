// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pprint/plog"
)

var MarshalJSONMarshalersTests = []marshalTests{
	{
		line:     line(),
		input:    map[string]json.Marshaler{"json slice": plog.JSONMarshalers(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		want:     `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
		wantText: `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
		wantJSON: `{
			"json slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without jsons": plog.JSONMarshalers()},
		want:     `null`,
		wantText: `null`,
		wantJSON: `{
			"without jsons":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of empty jsons": plog.JSONMarshalers(plog.String(""), plog.String(""))},
		want:     `["",""]`,
		wantText: `["",""]`,
		wantJSON: `{
			"slice of empty jsons":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"slice of json nils": plog.JSONMarshalers(nil, nil)},
		want:     `[null,null]`,
		wantText: `[null,null]`,
		wantJSON: `{
			"slice of json nils":[null,null]
		}`,
	},
}

func TestMarshalJSONMarshalers(t *testing.T) {
	testMarshal(t, MarshalJSONMarshalersTests)
}
