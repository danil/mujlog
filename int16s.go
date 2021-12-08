// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Int16s returns stringer/JSON/text marshaler for the int16 slice type.
func Int16s(s ...int16) int16S { return int16S{S: s} }

type int16S struct{ S []int16 }

func (s int16S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s int16S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		b, err := int16V{V: v}.MarshalText()
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

func (s int16S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := int16V{V: v}.MarshalJSON()
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
