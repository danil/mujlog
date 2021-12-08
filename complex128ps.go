// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import "bytes"

// Complex128ps returns stringer/JSON/text marshaler for the complex128 pointer slice type.
func Complex128ps(a ...*complex128) complex128PS { return complex128PS{A: a} }

type complex128PS struct{ A []*complex128 }

func (a complex128PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a complex128PS) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		b, err := complex128P{P: p}.MarshalText()
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

func (a complex128PS) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		b, err := complex128P{P: p}.MarshalJSON()
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
