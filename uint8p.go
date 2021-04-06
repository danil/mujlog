// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Uint8p returns stringer/JSON/text marshaler for the uint8 pointer type.
func Uint8p(p *uint8) uint8P { return uint8P{P: p} }

type uint8P struct{ P *uint8 }

func (p uint8P) String() string {
	if p.P == nil {
		return "null"
	}
	return uint8V{V: *p.P}.String()
}

func (p uint8P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uint8P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
