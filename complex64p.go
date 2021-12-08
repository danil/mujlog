// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

// Complex64p returns stringer/JSON/text marshaler for the complex64 pointer type.
func Complex64p(p *complex64) complex64P { return complex64P{P: p} }

type complex64P struct{ P *complex64 }

func (p complex64P) String() string {
	if p.P == nil {
		return "null"
	}
	return complex64V{V: *p.P}.String()
}

func (p complex64P) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return complex64V{V: *p.P}.MarshalText()
}

func (p complex64P) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return complex64V{V: *p.P}.MarshalJSON()
}
