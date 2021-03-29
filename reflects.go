// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// Reflects returns stringer/JSON marshaler interface implementation uses reflection for the slice of some typea.

func Reflects(a ...interface{}) reflectsV { return reflectsV{A: a} }

type reflectsV struct{ A []interface{} }

func (a reflectsV) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a reflectsV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := reflectV{V: v}.MarshalText()
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

func (a reflectsV) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := reflectV{V: v}.MarshalJSON()
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
