// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Uint32p returns stringer/JSON/text marshaler for the uint32 pointer type.
func Uint32p(p *uint32) uint32P { return uint32P{P: p} }

type uint32P struct{ P *uint32 }

func (p uint32P) String() string {
	if p.P == nil {
		return "null"
	}
	return uint32V{V: *p.P}.String()
}

func (p uint32P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uint32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
