// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"fmt"
)

// Complex128 returns stringer/JSON marshaler for the complex128 type.
func Complex128(v complex128) complex128V { return complex128V{V: v} }

type complex128V struct{ V complex128 }

func (v complex128V) String() string {
	s := fmt.Sprintf("%g", v.V)
	return s[1 : len(s)-1]
}

func (v complex128V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v complex128V) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
