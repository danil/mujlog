// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Int64p returns stringer/JSON/text marshaler for the int64 pointer type.
func Int64p(p *int64) int64P { return int64P{P: p} }

type int64P struct{ P *int64 }

func (p int64P) String() string {
	if p.P == nil {
		return "null"
	}
	return int64V{V: *p.P}.String()
}

func (p int64P) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p int64P) MarshalJSON() ([]byte, error) {
	return p.MarshalText()
}
