// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Bytesp returns stringer/JSON marshaler interface implementation for the pointer to the byte slice type.
func Bytesp(p *[]byte) bytesP { return bytesP{P: p} }

type bytesP struct{ P *[]byte }

func (p bytesP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p bytesP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return bytesV{V: *p.P}.MarshalText()
}

func (p bytesP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return bytesV{V: *p.P}.MarshalJSON()
}
