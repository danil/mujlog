// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Uintp returns stringer/JSON/text marshaler for the uint pointer type.
func Uintp(p *uint) uintP { return uintP{P: p} }

type uintP struct{ P *uint }

func (p uintP) String() string {
	if p.P == nil {
		return "null"
	}
	return uintV{V: *p.P}.String()
}

func (p uintP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uintP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
