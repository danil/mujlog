// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"

	"github.com/danil/log0"
)

var MarshalJSONsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"time slice": log0.JSONs(time.Date(1970, time.January, 1, 2, 3, 4, 42, time.UTC), time.Date(1970, time.December, 5, 4, 3, 2, 1, time.UTC))},
		expected:     `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
		expectedText: `["1970-01-01T02:03:04.000000042Z","1970-12-05T04:03:02.000000001Z"]`,
		expectedJSON: `{
			"time slice":["1970-01-01T02:03:04.000000042Z", "1970-12-05T04:03:02.000000001Z"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without jsons": log0.JSONs()},
		expected:     `null`,
		expectedText: `null`,
		expectedJSON: `{
			"without jsons":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of empty jsons": log0.JSONs(log0.String(""), log0.String(""))},
		expected:     `["",""]`,
		expectedText: `["",""]`,
		expectedJSON: `{
			"slice of empty jsons":["",""]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of json nils": log0.JSONs(nil, nil)},
		expected:     `[null,null]`,
		expectedText: `[null,null]`,
		expectedJSON: `{
			"slice of json nils":[null,null]
		}`,
	},
}

func TestMarshalJSONs(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalJSONsTestCases)
}
