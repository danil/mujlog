// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Int64 returns stringer/JSON marshaler interface implementation for the int64 type.
func Int64(v int64) int64V { return int64V{V: v} }

type int64V struct{ V int64 }

func (v int64V) String() string {
	return strconv.FormatInt(int64(v.V), 10)
}

func (v int64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v int64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
