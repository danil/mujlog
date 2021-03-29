// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Stringp returns stringer/JSON marshaler interface implementation for the pointer to the string type.
func Stringp(p *string) stringP { return stringP{P: p} }

type stringP struct{ P *string }

func (p stringP) String() string {
	if p.P == nil {
		return "null"
	}
	return stringV{V: *p.P}.String()
}

func (p stringP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return stringV{V: *p.P}.MarshalText()
}

func (p stringP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return stringV{V: *p.P}.MarshalJSON()
}
