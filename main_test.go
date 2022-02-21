// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

type Struct struct {
	Name string
	Age  int
}

type testprinter struct {
	t    *testing.T
	link string
}

func (p testprinter) Errorf(msg string, args ...interface{}) {
	p.t.Errorf(p.link+"\n"+msg, args...)
}

// line reports file and line number information about function invocations.
func line() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "It was not possible to recover file and line number information about function invocations!"
}
