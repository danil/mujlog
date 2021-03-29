// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"sync"

	"github.com/danil/log0/encode0"
)

var bufPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Bytes returns stringer/JSON marshaler interface implementation for the byte slice type.
func Bytes(v []byte) bytesV { return bytesV{V: v} }

type bytesV struct{ V []byte }

func (v bytesV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v bytesV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}

	var buf bytes.Buffer

	err := encode0.Bytes(&buf, v.V)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v bytesV) MarshalJSON() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}

	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}

	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
