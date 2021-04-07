// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"encoding/json"
)

// JSONMarshalers returns stringer/JSON/text marshaler for the JSON marshaler slice type.
func JSONMarshalers(s ...json.Marshaler) jsonMarshalerS { return jsonMarshalerS{S: s} }

type jsonMarshalerS struct{ S []json.Marshaler }

func (s jsonMarshalerS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s jsonMarshalerS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	return s.MarshalJSON()
}

func (s jsonMarshalerS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	var tail bool
	for _, v := range s.S {
		if tail {
			buf.WriteString(",")
		}
		if v == nil {
			buf.WriteString("null")
			tail = true
			continue
		}
		b, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(b)
		if err != nil {
			return nil, err
		}
		tail = true
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
