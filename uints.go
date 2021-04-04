// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Uints returns stringer/JSON marshaler for the uint slice type.
func Uints(s ...uint) uintS { return uintS{S: s} }

type uintS struct{ S []uint }

func (s uintS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s uintS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := uintV{V: v}.MarshalText()
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

func (s uintS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := uintV{V: v}.MarshalJSON()
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
