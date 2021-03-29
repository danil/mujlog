// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"runtime"
	"testing"
)

// FIXME: test missing!!!
var MarshalReflectsTestCases = []marshalTestCase{}

func TestMarshalReflects(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	testMarshal(t, testFile, MarshalReflectsTestCases)
}
