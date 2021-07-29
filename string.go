// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"bytes"

	"github.com/kvlog/log0/encode0"
)

// String returns stringer/JSON/text marshaler for the string type.
func String(v string) stringV { return stringV{V: v} }

type stringV struct{ V string }

func (v stringV) String() string {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	err := encode0.String(buf, v.V)
	if err != nil {
		return ""
	}
	return buf.String()
}

func (v stringV) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	err := encode0.String(&buf, v.V)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (v stringV) MarshalJSON() ([]byte, error) {
	b, err := v.MarshalText()
	if err != nil {
		return nil, err
	}
	return append([]byte(`"`), append(b, []byte(`"`)...)...), nil
}
