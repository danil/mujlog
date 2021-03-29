// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import "bytes"

// Complex128sp returns stringer/JSON marshaler interface implementation for slice of pointers to the complex128 type.
func Complex128sp(a ...*complex128) complex128sp { return complex128sp{A: a} }

type complex128sp struct{ A []*complex128 }

func (a complex128sp) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a complex128sp) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := complex128P{P: p}.MarshalText()
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

func (a complex128sp) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := complex128P{P: p}.MarshalJSON()
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
