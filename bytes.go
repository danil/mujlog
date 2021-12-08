// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"bytes"
	"sync"

	"github.com/pprint/plog/pencode"
)

var bufPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Bytes returns stringer/JSON/text marshaler for slice of bytes type.
func Bytes(s ...byte) byteS { return byteS{S: s} }

type byteS struct{ S []byte }

func (s byteS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s byteS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	err := pencode.Bytes(&buf, s.S)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s byteS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	b, err := s.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
