// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"fmt"
)

// Complex64 returns stringer/JSON/text marshaler for the complex64 type.
func Complex64(v complex64) complex64V { return complex64V{V: v} }

type complex64V struct{ V complex64 }

func (v complex64V) String() string {
	s := fmt.Sprintf("%g", v.V)
	return s[1 : len(s)-1]
}

func (v complex64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v complex64V) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
