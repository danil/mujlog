// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"encoding/json"
)

// JSONs returns stringer/JSON/text marshaler for the JSON marshaler slice type.
func JSONs(s ...json.Marshaler) jsonS { return jsonS{S: s} }

type jsonS struct{ S []json.Marshaler }

func (s jsonS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s jsonS) MarshalText() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	return s.MarshalJSON()
}

func (s jsonS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		if i != 0 {
			buf.WriteString(",")
		}
		if v == nil {
			buf.WriteString("null")
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
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
