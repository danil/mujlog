// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import "bytes"

// Complex64sp returns stringer/JSON marshaler interface implementation for slice of a pointer to the complex64 type.
func Complex64sp(a ...*complex64) complex64sp { return complex64sp{A: a} }

type complex64sp struct{ A []*complex64 }

func (a complex64sp) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a complex64sp) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := complex64P{P: p}.MarshalText()
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

func (a complex64sp) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := complex64P{P: p}.MarshalJSON()
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
