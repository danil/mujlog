// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Complex128s returns stringer/JSON marshaler for the complex128 slice type.
func Complex128s(s ...complex128) complex128S { return complex128S{S: s} }

type complex128S struct{ S []complex128 }

func (s complex128S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s complex128S) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := complex128V{V: v}.MarshalText()
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

func (s complex128S) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := complex128V{V: v}.MarshalJSON()
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
