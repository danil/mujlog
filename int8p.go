// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

// Int8p returns stringer/JSON/text marshaler for the int8 pointer type.
func Int8p(p *int8) int8P { return int8P{P: p} }

type int8P struct{ P *int8 }

func (p int8P) String() string {
	if p.P == nil {
		return "null"
	}
	return int8V{V: *p.P}.String()
}

func (p int8P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p int8P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
