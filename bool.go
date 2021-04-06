// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Bool returns stringer/JSON/text marshaler for the bool type.
func Bool(v bool) boolV { return boolV{V: v} }

type boolV struct{ V bool }

func (v boolV) String() string {
	b, _ := v.MarshalText()
	return string(b)
}

func (v boolV) MarshalText() ([]byte, error) {
	if v.V {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (v boolV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
