// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
)

// Bytesps returns stringer/JSON/text marshaler for slice of pointers to byte slice type.
func Bytesps(a ...*[]byte) byteSPS { return byteSPS{A: a} }

type byteSPS struct{ A []*[]byte }

func (a byteSPS) String() string {
	t, _ := a.MarshalText()
	return string(t)
}

func (a byteSPS) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		b, err := byteSP{P: p}.MarshalText()
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

func (a byteSPS) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		b, err := byteSP{P: p}.MarshalJSON()
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
