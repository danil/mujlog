// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Uint64 returns stringer/JSON marshaler interface implementation for the uint64 type.
func Uint64(v uint64) uint64V { return uint64V{V: v} }

type uint64V struct{ V uint64 }

func (v uint64V) String() string {
	return strconv.FormatUint(v.V, 10)
}

func (v uint64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v uint64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
