// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Runesp returns stringer/JSON marshaler for the rune pointer slice type.
func Runesp(p *[]rune) runeSP { return runeSP{P: p} }

type runeSP struct{ P *[]rune }

func (p runeSP) String() string {
	if p.P == nil {
		return "null"
	}
	return runeS{S: *p.P}.String()
}

func (p runeSP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return runeS{S: *p.P}.MarshalText()
}

func (p runeSP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return runeS{S: *p.P}.MarshalJSON()
}
