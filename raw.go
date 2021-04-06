// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

// Raw returns stringer/JSON/text marshaler for the raw byte slice.
func Raw(v []byte) rawV { return rawV{V: v} }

type rawV struct{ V []byte }

func (v rawV) String() string {
	if v.V == nil {
		return "null"
	}
	return string(v.V)
}

func (v rawV) MarshalText() ([]byte, error) {
	if v.V == nil {
		return []byte("null"), nil
	}
	return v.V, nil
}

func (v rawV) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
