// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
)

// JSON returns stringer/JSON/text marshaler for the KV slice type.
func JSON(s ...KV) jsonV { return jsonV{S: s} }

type jsonV struct{ S []KV }

func (s jsonV) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s jsonV) MarshalText() ([]byte, error) {
	if s.S == nil {
		return nil, nil
	}
	var buf bytes.Buffer
	var tail bool
	for _, s := range s.S {
		if s == nil {
			continue
		}
		k, err := s.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			continue
		}
		v, err := s.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if tail {
			buf.WriteString(" ")
		}
		_, err = buf.Write(append(k, append([]byte(" "), v...)...))
		if err != nil {
			return nil, err
		}
		tail = true
	}
	return buf.Bytes(), nil
}

func (s jsonV) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("{")
	var tail bool
	for _, s := range s.S {
		if s == nil {
			continue
		}
		k, err := s.MarshalText()
		if err != nil {
			return nil, err
		}
		if k == nil {
			continue
		}
		v, err := s.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if tail {
			buf.WriteString(",")
		}
		_, err = buf.Write(append([]byte(`"`), append(k, append([]byte(`":`), v...)...)...))
		if err != nil {
			return nil, err
		}
		tail = true
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}
