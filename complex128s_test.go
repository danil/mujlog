// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/danil/log0"
)

var MarshalComplex128sTestCases = []marshalTestCase{
	{
		line:         line(),
		input:        map[string]json.Marshaler{"complex128s": log0.Complex128s(complex(1, 23), complex(3, 21))},
		expected:     "1+23i 3+21i",
		expectedText: "1+23i 3+21i",
		expectedJSON: `{
			"complex128s":["1+23i","3+21i"]
		}`,
	},
}

func TestMarshalComplex128s(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalComplex128sTestCases)
}
