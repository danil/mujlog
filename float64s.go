// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Float64s returns stringer/JSON/text marshaler for the float64 slice type.
func Float64s(s ...float64) float64S { return float64S{S: s} }

type float64S struct{ S []float64 }

func (s float64S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s float64S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.S {
		b, err := float64V{V: v}.MarshalText()
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

func (s float64S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := float64V{V: v}.MarshalJSON()
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
