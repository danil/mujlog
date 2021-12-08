// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"time"
)

// Time returns stringer/JSON/text marshaler for the time type.
func Time(v time.Time) timeV { return timeV{V: v} }

type timeV struct{ V time.Time }

func (v timeV) String() string {
	return v.V.String()
}

func (v timeV) MarshalText() ([]byte, error) {
	return v.V.MarshalText()
}

func (v timeV) MarshalJSON() ([]byte, error) {
	return v.V.MarshalJSON()
}
