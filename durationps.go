// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"time"
)

// Durationps returns stringer/JSON marshaler for the time duration pointer slice type.
func Durationps(s ...*time.Duration) durationsp { return durationsp{S: s} }

type durationsp struct{ S []*time.Duration }

func (s durationsp) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s durationsp) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := durationP{P: p}.MarshalText()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s durationsp) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := durationP{P: p}.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
