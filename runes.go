// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"

	"github.com/kvlog/log0/encode0"
)

// Runes returns stringer/JSON/text marshaler for the rune slice type.
func Runes(s ...rune) runeS { return runeS{S: s} }

type runeS struct{ S []rune }

func (s runeS) String() string {
	if s.S == nil {
		return "null"
	}
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := encode0.Runes(buf, s.S)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (s runeS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer

	err := encode0.Runes(&buf, s.S)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s runeS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}

	b, err := s.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
