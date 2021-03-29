// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"

	"github.com/danil/log0/encode0"
)

// Runes returns stringer/JSON marshaler interface implementation for the rune slice type.
func Runes(v []rune) runesV { return runesV{V: v} }

type runesV struct{ V []rune }

func (v runesV) String() string {
	if v.V == nil {
		return "null"
	}
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := encode0.Runes(buf, v.V)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (v runesV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	err := encode0.Runes(&buf, v.V)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v runesV) MarshalJSON() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}

	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
