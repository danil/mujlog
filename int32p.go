// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Int32p returns stringer/JSON marshaler interface implementation for the pointer to the int32 type.
func Int32p(p *int32) int32P { return int32P{P: p} }

type int32P struct{ P *int32 }

func (p int32P) String() string {
	if p.P == nil {
		return "null"
	}
	return int32V{V: *p.P}.String()
}

func (p int32P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p int32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
