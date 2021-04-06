// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"

	"github.com/danil/log0/encode0"
)

// Errors returns stringer/JSON/text marshaler for the error slice type.
func Errors(s ...error) errorS { return errorS{S: s} }

type errorS struct{ S []error }

func (s errorS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s errorS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	for i, v := range s.S {
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

func (s errorS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
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
