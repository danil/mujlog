// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Ints returns stringer/JSON/text marshaler for the int slice type.
func Ints(s ...int) intS { return intS{S: s} }

type intS struct{ S []int }

func (s intS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s intS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		b, err := intV{V: v}.MarshalText()
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

func (s intS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := intV{V: v}.MarshalJSON()
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
