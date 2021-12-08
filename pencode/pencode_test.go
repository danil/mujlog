// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pencode

import (
	"bytes"
	"testing"
)

func TestBytes(t *testing.T) {
	for tt, want := range codec {
		tt := tt
		want := want
		t.Run(string(tt), func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			err := Bytes(&buf, []byte(string(tt)))
			if err != nil {
				t.Fatalf("encode bytes write error: %s", err)
			}

			if !bytes.Equal(buf.Bytes(), want) {
				t.Errorf("want: %s, get: %s", want, buf.String())
			}
		})
	}
}

func TestRunes(t *testing.T) {
	for tt, want := range codec {
		tt := tt
		want := want
		t.Run(string(tt), func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			err := Runes(&buf, []rune{tt})
			if err != nil {
				t.Fatalf("encode runes write error: %s", err)
			}

			if !bytes.Equal(buf.Bytes(), want) {
				t.Errorf("want: %s, get: %s", want, buf.String())
			}
		})
	}
}

func TestString(t *testing.T) {
	for tt, want := range codec {
		tt := tt
		want := want
		t.Run(string(tt), func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer

			err := String(&buf, string(tt))
			if err != nil {
				t.Fatalf("encode string write error: %s", err)
			}

			if !bytes.Equal(buf.Bytes(), want) {
				t.Errorf("want: %s, get: %s", want, buf.String())
			}
		})
	}
}
