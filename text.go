// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"encoding"

	"github.com/danil/log0/encode0"
)

// Text returns stringer/JSON marshaler for the encoding.TextMarshaler type.
func Text(v encoding.TextMarshaler) textV { return textV{V: v} }

type textV struct{ V encoding.TextMarshaler }

func (v textV) String() string {
	if v.V == nil {
		return ""
	}
	b, err := v.V.MarshalText()
	if err != nil {
		return ""
	}
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err = encode0.Bytes(buf, b)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (v textV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte{}, nil
	}
	b, err := v.V.MarshalText()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer

	err = encode0.Bytes(&buf, b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v textV) MarshalJSON() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
