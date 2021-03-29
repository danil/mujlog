// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Errorp returns stringer/JSON marshaler interface implementation for a pointer to a error type.
func Errorp(p *error) errorp { return errorp{P: p} }

type errorp struct{ P *error }

func (p errorp) String() string {
	if p.P == nil {
		return "null"
	}
	return errorV{V: *p.P}.String()
}

func (p errorp) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return errorV{V: *p.P}.MarshalText()
}

func (p errorp) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return errorV{V: *p.P}.MarshalJSON()
}
