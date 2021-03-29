// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Uint8 returns stringer/JSON marshaler interface implementation for the uint8 type.
func Uint8(v uint8) uint8V { return uint8V{V: v} }

type uint8V struct{ V uint8 }

func (v uint8V) String() string {
	return strconv.FormatUint(uint64(v.V), 10)
}

func (v uint8V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v uint8V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
