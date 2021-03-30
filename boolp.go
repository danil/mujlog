// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Boolp returns stringer/JSON marshaler for the bool pointer type.
func Boolp(p *bool) boolP { return boolP{P: p} }

type boolP struct{ P *bool }

func (p boolP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p boolP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return boolV{V: *p.P}.MarshalText()
}

func (p boolP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
