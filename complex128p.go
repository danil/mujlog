// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Complex128p returns stringer/JSON marshaler for the complex128 pointer type.
func Complex128p(p *complex128) complex128P { return complex128P{P: p} }

type complex128P struct{ P *complex128 }

func (p complex128P) String() string {
	if p.P == nil {
		return "null"
	}
	return complex128V{V: *p.P}.String()
}

func (p complex128P) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return complex128V{V: *p.P}.MarshalText()
}

func (p complex128P) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return complex128V{V: *p.P}.MarshalJSON()
}
