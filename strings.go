// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"strings"

	"github.com/kvlog/log0/encode0"
)

// Strings returns stringer/JSON/text marshaler for the string slice type.
func Strings(s ...string) stringS { return stringS{S: s} }

type stringS struct{ S []string }

func (s stringS) String() string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := encode0.String(buf, strings.Join(s.S, " "))
	if err != nil {
		return ""
	}
	return buf.String()
}

func (s stringS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := encode0.String(&buf, strings.Join(s.S, " "))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s stringS) MarshalJSON() ([]byte, error) {
	if s.S == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.S {
		b, err := stringV{V: v}.MarshalJSON()
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
