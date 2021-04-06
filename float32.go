// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"fmt"
)

// Float32 returns stringer/JSON/text marshaler for the float32 type.
func Float32(v float32) float32V { return float32V{V: v} }

type float32V struct{ V float32 }

func (v float32V) String() string {
	return fmt.Sprint(v.V)
}

func (v float32V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v float32V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
