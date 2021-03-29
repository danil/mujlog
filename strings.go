// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"
	"strings"

	"github.com/danil/log0/encode0"
)

// Strings returns stringer/JSON marshaler interface implementation for the string slice type.
func Strings(a ...string) stringsV { return stringsV{A: a} }

type stringsV struct{ A []string }

func (a stringsV) String() string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := encode0.String(buf, strings.Join(a.A, " "))
	if err != nil {
		return ""
	}
	return buf.String()
}

func (a stringsV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := encode0.String(&buf, strings.Join(a.A, " "))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (a stringsV) MarshalJSON() ([]byte, error) {
	if a.A == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range a.A {
		if i != 0 {
			buf.WriteString(",")
		}
		b, err := stringV{V: v}.MarshalJSON()
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
