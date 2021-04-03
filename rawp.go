// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Rawp returns stringer/JSON marshaler for the raw byte slice pointer.
func Rawp(p *[]byte) rawP { return rawP{P: p} }

type rawP struct{ P *[]byte }

func (p rawP) String() string {
	t, _ := p.MarshalText()
	return string(t)
}

func (p rawP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return rawV{V: *p.P}.MarshalText()
}

func (p rawP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return rawV{V: *p.P}.MarshalJSON()
}
