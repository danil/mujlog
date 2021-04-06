// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Float32p returns stringer/JSON/text marshaler for the float32 pointer type.
func Float32p(p *float32) float32P { return float32P{P: p} }

type float32P struct{ P *float32 }

func (p float32P) String() string {
	if p.P == nil {
		return "null"
	}
	return float32V{V: *p.P}.String()
}

func (p float32P) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return float32V{V: *p.P}.MarshalText()
}

func (p float32P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
