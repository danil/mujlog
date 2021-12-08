// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

// Errorp returns stringer/JSON/text marshaler for the error pointer type.
func Errorp(p *error) errorP { return errorP{P: p} }

type errorP struct{ P *error }

func (p errorP) String() string {
	if p.P == nil {
		return "null"
	}
	return errorV{V: *p.P}.String()
}

func (p errorP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return errorV{V: *p.P}.MarshalText()
}

func (p errorP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return errorV{V: *p.P}.MarshalJSON()
}
