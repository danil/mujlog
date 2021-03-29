// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"

	"github.com/danil/log0/encode0"
)

// Errors returns stringer/JSON marshaler interface implementation for the error slice type.
func Errors(a ...error) errorsV { return errorsV{A: a} }

type errorsV struct{ A []error }

func (a errorsV) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a errorsV) MarshalText() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	for i, v := range a.A {
		if v == nil {
			continue
		}
		if i != 0 {
			buf.WriteString(" ")
		}
		err := encode0.String(&buf, v.Error())
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (a errorsV) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := errorV{V: v}.MarshalJSON()
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
