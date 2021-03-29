// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Runesp returns stringer/JSON marshaler interface implementation for the pointer to the rune slice type.
func Runesp(p *[]rune) runesP { return runesP{P: p} }

type runesP struct{ P *[]rune }

func (p runesP) String() string {
	if p.P == nil {
		return "null"
	}
	return runesV{V: *p.P}.String()
}

func (p runesP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return runesV{V: *p.P}.MarshalText()
}

func (p runesP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return runesV{V: *p.P}.MarshalJSON()
}
