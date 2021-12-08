// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"strconv"
)

// Int32 returns stringer/JSON/text marshaler for the int32 type.
func Int32(v int32) int32V { return int32V{V: v} }

type int32V struct{ V int32 }

func (v int32V) String() string {
	return strconv.Itoa(int(v.V))
}

func (v int32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v int32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
