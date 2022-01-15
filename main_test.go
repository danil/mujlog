// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kinbiko/jsonassert"
)

type Struct struct {
	Name string
	Age  int
}

type testprinter struct {
	t    *testing.T
	link string
}

func (p testprinter) Errorf(msg string, args ...interface{}) {
	p.t.Errorf(p.link+"\n"+msg, args...)
}

type marshalTest struct {
	line      string
	input     map[string]json.Marshaler
	want      string
	wantText  string
	wantJSON  string
	wantError error
}

func testMarshal(t *testing.T, tests []marshalTest) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.line+"/"+fmt.Sprint(tt.input), func(t *testing.T) {
			t.Parallel()

			for k, v := range tt.input {
				str, ok := v.(fmt.Stringer)
				if !ok {
					t.Errorf("%q does not implement the stringer interface", k)

				} else {
					s := str.String()
					if s != tt.want {
						t.Errorf("%q unwant string, want: %q, recieved: %q %s", k, tt.want, s, tt.line)
					}
				}

				txt, ok := v.(encoding.TextMarshaler)
				if !ok {
					t.Errorf("%q does not implement the text marshaler interface", k)

				} else {
					p, err := txt.MarshalText()
					if err != nil {
						t.Fatalf("%q encoding marshal text error: %s %s", k, err, tt.line)
					}

					if string(p) != tt.wantText {
						t.Errorf("%q unwant text, want: %q, recieved: %q %s", k, tt.wantText, string(p), tt.line)
					}
				}
			}

			p, err := json.Marshal(tt.input)

			if fmt.Sprint(err) != fmt.Sprint(tt.wantError) {
				t.Fatalf("marshal error want: %s, recieved: %s %s", tt.wantError, err, tt.line)
			}

			if err == nil {
				ja := jsonassert.New(testprinter{t: t, link: tt.line})
				ja.Assertf(string(p), tt.wantJSON)
			}
		})
	}
}

// line reports file and line number information about function invocations.
func line() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return "It was not possible to recover file and line number information about function invocations!"
}
