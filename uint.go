// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Uint returns stringer/JSON marshaler interface implementation for the uint type.
func Uint(v uint) uintV { return uintV{V: v} }

type uintV struct{ V uint }

func (v uintV) String() string {
	return strconv.FormatUint(uint64(v.V), 10)
}

func (v uintV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v uintV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
