// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Errorps returns stringer/JSON marshaler for the slice of error pointers type.
func Errorps(s ...*error) errorPS { return errorPS{S: s} }

type errorPS struct{ S []*error }

func (s errorPS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s errorPS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	for i, p := range s.S {
		if p == nil {
			continue
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := errorP{P: p}.MarshalText()
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

func (s errorPS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := errorP{P: p}.MarshalJSON()
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
