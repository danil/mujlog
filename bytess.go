// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Bytess returns stringer/JSON/text marshaler for the slice of byte slice type.
func Bytess(s ...[]byte) byteSS { return byteSS{S: s} }

type byteSS struct{ S [][]byte }

func (s byteSS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s byteSS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, s := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := byteS{S: s}.MarshalText()
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

func (s byteSS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, s := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := byteS{S: s}.MarshalJSON()
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
