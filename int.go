// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Int returns stringer/JSON marshaler interface implementation for the int type.
func Int(v int) intV { return intV{V: v} }

type intV struct{ V int }

func (v intV) String() string {
	return strconv.Itoa(v.V)
}

func (v intV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v intV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
