// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"

	"github.com/pprint/plog/pencode"
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
	var tail bool

	for _, v := range s.S {
		if v == nil {
			continue
		}
		if tail {
			buf.WriteString(" ")
		}
		err := pencode.String(&buf, v.Error())
		if err != nil {
			return nil, err
		}
		tail = true
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
		b, err := errorV{V: v}.MarshalJSON()
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
