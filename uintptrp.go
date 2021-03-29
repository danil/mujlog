// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Uintptrp returns stringer/JSON marshaler interface implementation for the pointer to the uintptr type.
func Uintptrp(p *uintptr) uintptrP { return uintptrP{P: p} }

type uintptrP struct{ P *uintptr }

func (p uintptrP) String() string {
	if p.P == nil {
		return "null"
	}
	return uintptrV{V: *p.P}.String()
}

func (p uintptrP) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p uintptrP) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
