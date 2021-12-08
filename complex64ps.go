// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import "bytes"

// Complex64ps returns stringer/JSON/text marshaler for the slice of complex64 pointers type.
func Complex64ps(s ...*complex64) complex64PS { return complex64PS{S: s} }

type complex64PS struct{ S []*complex64 }

func (s complex64PS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s complex64PS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range s.S {
		b, err := complex64P{P: p}.MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (s complex64PS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range s.S {
		b, err := complex64P{P: p}.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
