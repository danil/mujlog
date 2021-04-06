// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Int16p returns stringer/JSON/text marshaler for the int16 pointer type.
func Int16p(p *int16) int16P { return int16P{P: p} }

type int16P struct{ P *int16 }

func (p int16P) String() string {
	if p.P == nil {
		return "null"
	}
	return int16V{V: *p.P}.String()
}

func (p int16P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p int16P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
