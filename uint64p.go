// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

// Uint64p returns stringer/JSON/text marshaler for the uint64 pointer type.
func Uint64p(p *uint64) uint64P { return uint64P{P: p} }

type uint64P struct{ P *uint64 }

func (p uint64P) String() string {
	if p.P == nil {
		return "null"
	}
	return uint64V{V: *p.P}.String()
}

func (p uint64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uint64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
