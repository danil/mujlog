// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Anys returns stringer/JSON marshaler for the slice of any type.
func Anys(s ...interface{}) anyS { return anyS{S: s} }

type anyS struct{ S []interface{} }

func (s anyS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s anyS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := anyV{V: v}.MarshalText()
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

func (s anyS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := anyV{V: v}.MarshalJSON()
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
