// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

// Uint16p returns stringer/JSON/text marshaler for the uint16 pointer type.
func Uint16p(p *uint16) uint16P { return uint16P{P: p} }

type uint16P struct{ P *uint16 }

func (p uint16P) String() string {
	if p.P == nil {
		return "null"
	}
	return uint16V{V: *p.P}.String()
}

func (p uint16P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uint16P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
