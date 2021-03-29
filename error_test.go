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

var MarshalErrorTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"error": log0.Error(errors.New("something went wrong"))},
		expected:     "something went wrong",
		expectedText: "something went wrong",
		expectedJSON: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil error": log0.Error(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil error":null
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"any error": log0.Any(errors.New("something went wrong"))},
		expected:     "something went wrong",
		expectedText: "something went wrong",
		expectedJSON: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"reflect error": log0.Reflect(errors.New("something went wrong"))},
		expected:     "{something went wrong}",
		expectedText: "{something went wrong}",
		expectedJSON: `{
			"reflect error":{}
		}`,
	},
}

func TestMarshalError(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalErrorTestCases)
}
