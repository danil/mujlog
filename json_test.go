// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalJSONTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"kv slice": log0.JSON(log0.StringString("foo", "bar"), log0.StringInt("xyz", 42))},
		expected:     `foo "bar" xyz 42`,
		expectedText: `foo "bar" xyz 42`,
		expectedJSON: `{
			"kv slice":{"foo":"bar","xyz":42}
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without jsons": log0.JSON()},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"without jsons":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of empty jsons": log0.JSON(log0.String(""), log0.String(""))},
		expected:     ``,
		expectedText: ``,
		expectedJSON: `{
			"slice of empty jsons":{}
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"slice of json nils": log0.JSON(nil, nil)},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"slice of json nils":{}
		}`,
	},
}

func TestMarshalJSON(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalJSONTestCases)
}
