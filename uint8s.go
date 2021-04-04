// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Uint8s returns stringer/JSON marshaler for the uint8 slice type.
func Uint8s(s ...uint8) uint8S { return uint8S{S: s} }

type uint8S struct{ S []uint8 }

func (s uint8S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s uint8S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := uint8V{V: v}.MarshalText()
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

func (s uint8S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := uint8V{V: v}.MarshalJSON()
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
