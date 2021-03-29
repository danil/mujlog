// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Bytess returns stringer/JSON marshaler interface implementation for the slice of byte slice type.
func Bytess(a ...[]byte) bytessV { return bytessV{A: a} }

type bytessV struct{ A [][]byte }

func (a bytessV) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a bytessV) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := bytesV{V: v}.MarshalText()
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

func (a bytessV) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := bytesV{V: v}.MarshalJSON()
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
