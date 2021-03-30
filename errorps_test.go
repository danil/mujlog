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

var MarshalErrorpsTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"error pointer slice": log0.Errorps(&err, &err2)}
		}(),
		expected:     "something went wrong we have a problem",
		expectedText: "something went wrong we have a problem",
		expectedJSON: `{
			"error pointer slice":["something went wrong","we have a problem"]
		}`,
	},
	{
		line:  line(),
		input: map[string]json.Marshaler{"nil error pointers": log0.Errorps(nil, nil)},
		expectedJSON: `{
			"nil error pointers":[null,null]
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"without error pointers": log0.Errorps()},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"without error pointers":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of any error pointers": log0.Anys(&err, &err2)}
		}(),
		expected:     "something went wrong we have a problem",
		expectedText: "something went wrong we have a problem",
		expectedJSON: `{
			"slice of any error pointers":["something went wrong","we have a problem"]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err, err2 := errors.New("something went wrong"), errors.New("we have a problem")
			return map[string]json.Marshaler{"slice of reflect of error pointers": log0.Reflects(&err, &err2)}
		}(),
		expected:     "{something went wrong} {we have a problem}",
		expectedText: "{something went wrong} {we have a problem}",
		expectedJSON: `{
			"slice of reflect of error pointers":[{},{}]
		}`,
	},
}

func TestMarshalErrorps(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalErrorpsTestCases)
}
