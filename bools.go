// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Bools returns stringer/JSON marshaler interface implementation for the bool slice type.
func Bools(a ...bool) boolsV { return boolsV{A: a} }

type boolsV struct{ A []bool }

func (a boolsV) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a boolsV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := boolV{V: v}.MarshalText()
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

func (a boolsV) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := boolV{V: v}.MarshalJSON()
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
