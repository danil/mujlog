// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Float64p returns stringer/JSON marshaler for the  float64 pointer type.
func Float64p(p *float64) float64P { return float64P{P: p} }

type float64P struct{ P *float64 }

func (p float64P) String() string {
	if p.P == nil {
		return "null"
	}
	return float64V{V: *p.P}.String()
}

func (p float64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p float64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
