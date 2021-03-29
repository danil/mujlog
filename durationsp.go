// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"time"
)

// Durationsp returns stringer/JSON marshaler interface implementation for slice of pointers to time duration type.
func Durationsp(a ...*time.Duration) durationsp { return durationsp{A: a} }

type durationsp struct{ A []*time.Duration }

func (a durationsp) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a durationsp) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(" ")
		}
		b, err := durationP{P: p}.MarshalText()
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

func (a durationsp) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := durationP{P: p}.MarshalJSON()
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
