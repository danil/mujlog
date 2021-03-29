// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalErrorsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"errors": log0.Errors(errors.New("something went wrong"), errors.New("wrong"))},
		expected:     "something went wrong wrong",
		expectedText: "something went wrong wrong",
		expectedJSON: `{
			"errors":["something went wrong","wrong"]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil errors": log0.Errors(nil, nil)},
		expected:     "",
		expectedText: "",
		expectedJSON: `{
			"nil errors":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without errors": log0.Errors()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"without errors":null
		}`,
	},
}

func TestMarshalErrors(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalErrorsTestCases)
}
