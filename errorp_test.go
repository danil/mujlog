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

var MarshalErrorpTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"error pointer": log0.Errorp(&err)}
		}(),
		expected:     "something went wrong",
		expectedText: "something went wrong",
		expectedJSON: `{
			"error pointer":"something went wrong"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil error pointer": log0.Errorp(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil error pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"any error pointer": log0.Any(&err)}
		}(),
		expected:     "something went wrong",
		expectedText: "something went wrong",
		expectedJSON: `{
			"any error pointer":"something went wrong"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			err2 := &err
			return map[string]json.Marshaler{"any twice/nested pointer to error": log0.Any(&err2)}
		}(),
		expected:     "{something went wrong}",
		expectedText: "{something went wrong}",
		expectedJSON: `{
			"any twice/nested pointer to error":{}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"reflect error pointer": log0.Reflect(&err)}
		}(),
		expected:     "{something went wrong}",
		expectedText: "{something went wrong}",
		expectedJSON: `{
			"reflect error pointer":{}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			err2 := &err
			return map[string]json.Marshaler{"reflect twice/nested pointer to error": log0.Reflect(&err2)}
		}(),
		expected:     "{something went wrong}",
		expectedText: "{something went wrong}",
		expectedJSON: `{
			"reflect twice/nested pointer to error":{}
		}`,
	},
}

func TestErrorpMarshal(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalErrorpTestCases)
}
