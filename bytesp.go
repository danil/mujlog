// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Bytesp returns stringer/JSON marshaler for the pointer to byte slice type.
func Bytesp(p *[]byte) byteSP { return byteSP{P: p} }

type byteSP struct{ P *[]byte }

func (p byteSP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p byteSP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return byteS{S: *p.P}.MarshalText()
}

func (p byteSP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return byteS{S: *p.P}.MarshalJSON()
}
