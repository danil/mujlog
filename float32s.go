// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Float32s returns stringer/JSON/text marshaler for the float32 slice type.
func Float32s(s ...float32) float32S { return float32S{S: s} }

type float32S struct{ S []float32 }

func (s float32S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s float32S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		b, err := float32V{V: v}.MarshalText()
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

func (s float32S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := float32V{V: v}.MarshalJSON()
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
