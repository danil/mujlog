// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import "bytes"

// Complex64s returns stringer/JSON marshaler for the complex64 type.
func Complex64s(s ...complex64) complex64sV { return complex64sV{S: s} }

type complex64sV struct{ S []complex64 }

func (s complex64sV) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s complex64sV) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := complex64V{V: v}.MarshalText()
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

func (s complex64sV) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := complex64V{V: v}.MarshalJSON()
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
