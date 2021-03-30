// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Uint32 returns stringer/JSON marshaler for the uint32 type.
func Uint32(v uint32) uint32V { return uint32V{V: v} }

type uint32V struct{ V uint32 }

func (v uint32V) String() string {
	return strconv.FormatUint(uint64(v.V), 10)
}

func (v uint32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v uint32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
