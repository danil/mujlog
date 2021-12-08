// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Anys returns stringer/JSON/text marshaler for the slice of any type.
func Anys(s ...interface{}) AnyS { return AnyS{S: s} }

type AnyS struct{ S []interface{} }

func (s AnyS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s AnyS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		b, err := AnyV{V: v}.MarshalText()
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

func (s AnyS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := AnyV{V: v}.MarshalJSON()
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
