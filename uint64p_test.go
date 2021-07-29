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

var MarshalUint64pTestCases = []marshalTestCase{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"uint64 pointer": log0.Uint64p(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:         line(),
		input:        map[string]json.Marshaler{"nil uint64 pointer": log0.Uint64p(nil)},
		expected:     "null",
		expectedText: "null",
		expectedJSON: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"any uint64 pointer": log0.Any(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"reflect uint64 pointer": log0.Reflect(&i)}
		}(),
		expected:     "42",
		expectedText: "42",
		expectedJSON: `{
			"reflect uint64 pointer":42
		}`,
	},
}

func TestMarshalUint64p(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalUint64pTestCases)
}
