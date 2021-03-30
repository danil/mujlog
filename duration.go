// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"time"
)

// Duration returns stringer/JSON marshaler for the time duration type.
func Duration(v time.Duration) durationV { return durationV{V: v} }

type durationV struct{ V time.Duration }

func (v durationV) String() string {
	return v.V.String()
}

func (v durationV) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v durationV) MarshalJSON() ([]byte, error) {
	return append([]byte(`"`), append([]byte(v.String()), []byte(`"`)...)...), nil
}
