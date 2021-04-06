// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"strconv"
)

// Float64 returns stringer/JSON/text marshaler for the float64 type.
func Float64(v float64) float64V { return float64V{V: v} }

type float64V struct{ V float64 }

func (v float64V) String() string {
	return strconv.FormatFloat(float64(v.V), 'f', -1, 64)
}

func (v float64V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v float64V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
