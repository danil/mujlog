// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"strconv"
)

// Int16 returns stringer/JSON/text marshaler for the int16 type.
func Int16(v int16) int16V { return int16V{V: v} }

type int16V struct{ V int16 }

func (v int16V) String() string {
	return strconv.Itoa(int(v.V))
}

func (v int16V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v int16V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
