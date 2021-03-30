// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Intp returns stringer/JSON marshaler for the int pointer type.
func Intp(p *int) intP { return intP{P: p} }

type intP struct{ P *int }

func (p intP) String() string {
	if p.P == nil {
		return "null"
	}
	return intV{V: *p.P}.String()
}

func (p intP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p intP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
