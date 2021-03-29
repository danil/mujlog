// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Bytessp returns stringer/JSON marshaler interface implementation for the slice of byte pointer slices type.
func Bytessp(a ...*[]byte) bytessP { return bytessP{A: a} }

type bytessP struct{ A []*[]byte }

func (a bytessP) String() string {
	t, _ := a.MarshalText()
	return string(t)
}

func (a bytessP) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := bytesP{P: p}.MarshalText()
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

func (a bytessP) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := bytesP{P: p}.MarshalJSON()
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
