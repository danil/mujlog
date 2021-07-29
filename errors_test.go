// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"runtime"
	"testing"

	"github.com/kvlog/log0"
)

var MarshalErrorsTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"error slice": log0.Errors(errors.New("something went wrong"), errors.New("we have a problem"))},
		expected:     "something went wrong we have a problem",
		expectedText: "something went wrong we have a problem",
		expectedJSON: `{
			"error slice":["something went wrong","we have a problem"]
		}`,
	},
	{
		line:  line(),
		input: map[string]json.Marshaler{"nil errors": log0.Errors(nil, nil)},
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
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of any errors": log0.Anys(err, err2)}
		}(),
		expected:     "something went wrong we have a problem",
		expectedText: "something went wrong we have a problem",
		expectedJSON: `{
			"slice of any errors":["something went wrong","we have a problem"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of reflect of errors": log0.Reflects(err, err2)}
		}(),
		expected:     "{something went wrong} {we have a problem}",
		expectedText: "{something went wrong} {we have a problem}",
		expectedJSON: `{
			"slice of reflect of errors":[{},{}]
		}`,
	},
}

func TestMarshalErrors(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalErrorsTestCases)
}
