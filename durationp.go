// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"time"
)

// Durationp returns stringer/JSON/text marshaler for the  time duration pointer type.
func Durationp(p *time.Duration) durationP { return durationP{P: p} }

type durationP struct{ P *time.Duration }

func (p durationP) String() string {
	if p.P == nil {
		return "null"
	}
	return durationV{V: *p.P}.String()
}

func (p durationP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return durationV{V: *p.P}.MarshalText()
}

func (p durationP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return durationV{V: *p.P}.MarshalJSON()
}
