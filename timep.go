// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"time"
)

// Timep returns stringer/JSON/text marshaler for the time pointer type.
func Timep(p *time.Time) timeP { return timeP{P: p} }

type timeP struct{ P *time.Time }

func (p timeP) String() string {
	if p.P == nil {
		return "null"
	}
	return timeV{V: *p.P}.String()
}

func (p timeP) MarshalText() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return timeV{V: *p.P}.MarshalText()
}

func (p timeP) MarshalJSON() ([]byte, error) {
	if p.P == nil {
		return []byte("null"), nil
	}
	return timeV{V: *p.P}.MarshalJSON()
}
