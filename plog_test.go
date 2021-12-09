// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"bytes"
	"encoding"
	"fmt"
	"io"
	"log"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/pprint/pfmt"
	"github.com/pprint/plog"
)

var WriteTests = []struct {
	name      string
	line      string
	log       plog.Logger
	input     []byte
	kv        []plog.KV
	want      string
	benchmark bool
}{
	{
		name:  "nil",
		line:  line(),
		log:   dummy(),
		input: nil,
		want:  `{}`,
	},
	{
		name: `nil message with "foo" key with "bar" value`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("foo", "bar")},
		},
		input: nil,
		want: `{
	    "foo":"bar"
		}`,
	},
	{
		name:  "empty",
		line:  line(),
		log:   dummy(),
		input: []byte{},
		want: `{
	    "message":"",
			"excerpt":"_EMPTY_"
		}`,
	},
	{
		name: "blank",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Original,
			Trunc:  120,
			Marks:  [3][]byte{[]byte("…"), []byte("_EMPTY_"), []byte("_BLANK_")},
		},
		input: []byte(" "),
		want: `{
	    "message":" ",
			"excerpt":"_BLANK_"
		}`,
	},
	{
		name: `"string" key with "foo" value and "string" key with "bar" value`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("string", "foo")},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
			Trunc:  120,
		},
		input: []byte("Hello, World!"),
		kv:    []plog.KV{plog.StringString("string", "bar")},
		want: `{
			"message":"Hello, World!",
		  "string": "bar"
		}`,
		benchmark: true,
	},
	{
		name:  "kv is nil",
		line:  line(),
		log:   dummy(),
		input: []byte("Hello, World!"),
		kv:    nil,
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name: `bytes appends to the "message" key with "string value"`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			KV:      []plog.KV{plog.StringString("message", "string value")},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("Hello,\nWorld!"),
		want: `{
			"message":"string value",
			"excerpt":"Hello, World!",
			"trail":"Hello,\nWorld!"
		}`,
	},
	{
		name:  `bytes appends to the "message" key with "string value"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello,\nWorld!"),
		kv:    []plog.KV{plog.StringString("message", "string value")},
		want: `{
			"message":"string value",
			"excerpt":"Hello, World!",
			"trail":"Hello,\nWorld!"
		}`,
	},
	{
		name: `bytes is nil and "message" key with "string value"`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("message", "string value")},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
			Trunc:  120,
		},
		want: `{
			"message":"string value"
		}`,
	},
	{
		name: `input is nil and "message" key with "string value"`,
		line: line(),
		log:  dummy(),
		kv:   []plog.KV{plog.StringString("message", "string value")},
		want: `{
			"message":"string value"
		}`,
	},
	{
		name:  `bytes appends to the integer key "message"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello, World!\n"),
		kv:    []plog.KV{plog.StringInt("message", 1)},
		want: `{
			"message":1,
			"excerpt":"Hello, World!",
			"trail":"Hello, World!\n"
		}`,
	},
	{
		name:  `bytes appends to the float 32 bit key "message"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello,\nWorld!"),
		kv:    []plog.KV{plog.StringFloat32("message", 4.2)},
		want: `{
			"message":4.2,
			"excerpt":"Hello, World!",
			"trail":"Hello,\nWorld!"
		}`,
	},
	{
		name:  `bytes appends to the float 64 bit key "message"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello,\nWorld!"),
		kv:    []plog.KV{plog.StringFloat64("message", 4.2)},
		want: `{
			"message":4.2,
			"excerpt":"Hello, World!",
			"trail":"Hello,\nWorld!"
		}`,
	},
	{
		name:  `bytes appends to the boolean key "message"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello,\nWorld!"),
		kv:    []plog.KV{plog.StringBool("message", true)},
		want: `{
			"message":true,
			"excerpt":"Hello, World!",
			"trail":"Hello,\nWorld!"
		}`,
	},
	{
		name:  `bytes will appends to the nil key "message"`,
		line:  line(),
		log:   dummy(),
		input: []byte("Hello, World!"),
		kv:    []plog.KV{plog.StringReflect("message", nil)},
		want: `{
			"message":null,
			"trail":"Hello, World!"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "message" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
			Key:    plog.Original,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("message", "foo")},
		want: `{
			"message":"foo"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "message" key is present and with replace`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("message", "foo\n")},
		want: `{
			"message":"foo\n"
		}`,
	},
	{
		name: `default key is original and bytes is present and "message" key is present`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("message", "bar")},
		want: `{
			"message":"bar",
			"trail":"foo"
		}`,
	},
	{
		name: `default key is original and bytes is present and "message" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Trunc:  120,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:    plog.Original,
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar")},
		want: `{
			"message":"bar",
			"excerpt":"foo",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is original and bytes is present and "message" key is present and with replace input bytes and key`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar\n")},
		want: `{
			"message":"bar\n",
			"excerpt":"foo",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "excerpt" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Original,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("excerpt", "foo")},
		want: `{
			"excerpt":"foo"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "excerpt" key is present and with replace`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("excerpt", "foo\n")},
		want: `{
			"excerpt":"foo\n"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Original,
			Trunc:  120,
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" key is present and with replace input bytes and rey`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar\n")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar\n"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "excerpt" and "message" keys is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Original,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("message", "foo"), plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is original and bytes is nil and "excerpt" and "message" keys is present and replace keys`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("message", "foo\n"), plog.StringString("excerpt", "bar\n")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar\n"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" and "message" keys is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:    plog.Original,
			Trunc:  120,
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("message", "bar"), plog.StringString("excerpt", "xyz")},
		want: `{
			"message":"bar",
			"excerpt":"xyz",
			"trail":"foo"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" and "message" keys is present and replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv: []plog.KV{
			plog.StringString("message", "bar"),
			plog.StringString("excerpt", "xyz"),
		},
		want: `{
			"message":"bar",
			"excerpt":"xyz",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is original and bytes is present and "excerpt" and "message" keys is present and replace input bytes and keys`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Original,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar\n"), plog.StringString("excerpt", "xyz\n")},
		want: `{
			"message":"bar\n",
			"excerpt":"xyz\n",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "message" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("message", "foo")},
		want: `{
			"message":"foo"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "message" key is present and with replace`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("message", "foo\n")},
		want: `{
			"message":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "message" key is present`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("message", "bar")},
		want: `{
			"message":"bar",
			"excerpt":"foo"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "message" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar")},
		want: `{
			"message":"bar",
			"excerpt":"foo",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "message" key is present and with replace input bytes and key`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar\n")},
		want: `{
			"message":"bar\n",
			"excerpt":"foo",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "excerpt" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("excerpt", "foo")},
		want: `{
			"excerpt":"foo"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "excerpt" key is present and with replace`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("excerpt", "foo\n")},
		want: `{
			"excerpt":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" key is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" key is present and with replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" key is present and with replace input bytes and rey`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("excerpt", "bar\n")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "excerpt" and "message" keys is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		kv: []plog.KV{plog.StringString("message", "foo"), plog.StringString("excerpt", "bar")},
		want: `{
			"message":"foo",
			"excerpt":"bar"
		}`,
	},
	{
		name: `default key is excerpt and bytes is nil and "excerpt" and "message" keys is present and replace keys`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		kv: []plog.KV{plog.StringString("message", "foo\n"), plog.StringString("excerpt", "bar\n")},
		want: `{
			"message":"foo\n",
			"excerpt":"bar\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" and "message" keys is present`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:    plog.Excerpt,
			Trunc:  120,
		},
		input: []byte("foo"),
		kv:    []plog.KV{plog.StringString("message", "bar"), plog.StringString("excerpt", "xyz")},
		want: `{
			"message":"bar",
			"excerpt":"xyz",
			"trail":"foo"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" and "message" keys is present and replace input bytes`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar"), plog.StringString("excerpt", "xyz")},
		want: `{
			"message":"bar",
			"excerpt":"xyz",
			"trail":"foo\n"
		}`,
	},
	{
		name: `default key is excerpt and bytes is present and "excerpt" and "message" keys is present and replace input bytes and keys`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail")},
			Key:     plog.Excerpt,
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: []byte("foo\n"),
		kv:    []plog.KV{plog.StringString("message", "bar\n"), plog.StringString("excerpt", "xyz\n")},
		want: `{
			"message":"bar\n",
			"excerpt":"xyz\n",
			"trail":"foo\n"
		}`,
	},
	{
		name: `bytes is nil and bytes "message" key with json`,
		line: line(),
		log:  dummy(),
		kv:   []plog.KV{plog.StringBytes("message", []byte(`{"foo":"bar"}`)...)},
		want: `{
			"message":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		name: `bytes is nil and raw "message" key with json`,
		line: line(),
		log:  dummy(),
		kv:   []plog.KV{plog.StringRaw("message", []byte(`{"foo":"bar"}`))},
		want: `{
			"message":{"foo":"bar"}
		}`,
	},
	{
		name: "bytes is nil and flag is long file",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		kv: []plog.KV{plog.StringString("foo", "bar")},
		want: `{
			"foo":"bar"
		}`,
	},
	{
		name: "bytes is one char and flag is long file",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: []byte("a"),
		want: `{
			"message":"a"
		}`,
	},
	{
		name: "bytes is two chars and flag is long file",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
		},
		input: []byte("ab"),
		want: `{
			"message":"ab",
			"file":"a"
		}`,
	},
	{
		name: "bytes is three chars and flag is long file",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
		},
		input: []byte("abc"),
		want: `{
			"message":"abc",
			"file":"ab"
		}`,
	},
	{
		name: "permanent kv overwritten by the additional kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("foo", "bar")},
		},
		kv: []plog.KV{plog.StringString("foo", "baz")},
		want: `{
			"foo":"baz"
		}`,
	},
	{
		name: "permanent kv and first additional kv overwritten by the second additional kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("foo", "bar")},
		},
		kv: []plog.KV{
			plog.StringString("foo", "baz"),
			plog.StringString("foo", "xyz"),
		},
		want: `{
			"foo":"xyz"
		}`,
	},
}

func TestWrite(t *testing.T) {
	for _, tt := range WriteTests {
		tt := tt
		t.Run(tt.line+"/io writer %"+tt.name, func(t *testing.T) {
			t.Parallel()

			l0, ok := tt.log.(*plog.Log)
			if !ok {
				t.Fatal("unwant logger type")
			}

			buf, ok := l0.Output.(*bytes.Buffer)
			if !ok {
				t.Fatal("unwant output type")
			}

			*buf = bytes.Buffer{}

			l := tt.log.Tee(tt.kv...)
			defer l.Close()

			_, err := l.Write(tt.input)
			if err != nil {
				t.Fatalf("unwant write error: %s", err)
			}

			ja := jsonassert.New(testprinter{t: t, link: tt.line})
			ja.Assertf(buf.String(), tt.want)
		})
	}
}

var FprintWriteTests = []struct {
	name      string
	line      string
	log       plog.Logger
	input     interface{}
	want      string
	benchmark bool
}{
	{
		name: "readme example 1",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Marks:   [3][]byte{[]byte("…")},
			Trunc:   12,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		line:  line(),
		input: "Hello,\nWorld!",
		want: `{
			"message":"Hello,\nWorld!",
			"excerpt":"Hello, World…"
		}`,
	},
	{
		name: "readme example 2",
		line: line(),
		log: func() plog.Logger {
			l1 := plog.GELF()
			l1.Output = &bytes.Buffer{}
			l := l1.Tee(
				plog.StringString("version", "1.1"),
				plog.StringFunc("timestamp", func() plog.KV {
					t := time.Date(2020, time.October, 15, 18, 9, 0, 0, time.UTC)
					return pfmt.Int64(t.Unix())
				}),
			)
			return l
		}(),
		input: "Hello,\nGELF!",
		want: `{
			"version":"1.1",
			"short_message":"Hello, GELF!",
			"full_message":"Hello,\nGELF!",
			"timestamp":1602785340
		}`,
	},
	{
		name: "readme example 3.1",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		line:  line(),
		input: 3.21,
		want: `{
			"message":"3.21"
		}`,
	},
	{
		name: "readme example 3.2",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		line:  line(),
		input: 123,
		want: `{
			"message":"123"
		}`,
	},
	{
		name:  "string",
		line:  line(),
		log:   dummy(),
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name:  "integer type appears in the messages excerpt as a string",
		line:  line(),
		log:   dummy(),
		input: 123,
		want: `{
			"message":"123"
		}`,
	},
	{
		name:  "float type appears in the messages excerpt as a string",
		line:  line(),
		log:   dummy(),
		input: 3.21,
		want: `{
			"message":"3.21"
		}`,
	},
	{
		name:  "nil message",
		line:  line(),
		log:   dummy(),
		input: nil,
		want: `{
			"message":"<nil>"
		}`,
	},
	// FIXME: not working(
	// {
	// 	name: "empty message",
	// 	line: line(),
	// 	log: &plog.Log{
	// 		Output: &bytes.Buffer{},
	// 		Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
	// 		Key:    plog.Original,
	// 		Trunc:  120,
	// 		Marks:  [3][]byte{[]byte("…"), []byte("_EMPTY_"), []byte("_BLANK_")},
	// 		Replace:  [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
	// 	},
	// 	input: "",
	// 	want: `{
	// 		"message":"",
	// 		"excerpt":"_EMPTY_"
	// 	}`,
	// },
	{
		name:  "blank message",
		line:  line(),
		log:   dummy(),
		input: " ",
		want: `{
	    "message":" ",
			"excerpt":"_BLANK_"
		}`,
	},
	{
		name:  "single quotes",
		line:  line(),
		log:   dummy(),
		input: "foo 'bar'",
		want: `{
			"message":"foo 'bar'"
		}`,
	},
	{
		name:  "double quotes",
		line:  line(),
		log:   dummy(),
		input: `foo "bar"`,
		want: `{
			"message":"foo \"bar\""
		}`,
	},
	{
		name:  `leading/trailing "spaces"`,
		line:  line(),
		log:   dummy(),
		input: " \n\tHello, World! \t\n",
		want: `{
			"message":" \n\tHello, World! \t\n",
			"excerpt":"Hello, World!"
		}`,
	},
	{
		name:  "JSON string",
		line:  line(),
		log:   dummy(),
		input: `{"foo":"bar"}`,
		want: `{
			"message":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		name: `"string" key with "foo" value`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("string", "foo")},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!",
		  "string": "foo"
		}`,
	},
	{
		name: `"integer" key with 123 value`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringInt("integer", 123)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!",
		  "integer": 123
		}`,
	},
	{
		name: `"float" key with 3.21 value`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringFloat32("float", 3.21)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!",
		  "float": 3.21
		}`,
	},
	{
		name:  "fmt.Fprint prints nil as <nil>",
		line:  line(),
		log:   dummy(),
		input: nil,
		want: `{
			"message":"<nil>"
		}`,
	},
	{
		name:  "multiline string",
		line:  line(),
		log:   dummy(),
		input: "Hello,\nWorld\n!",
		want: `{
			"message":"Hello,\nWorld\n!",
			"excerpt":"Hello, World !"
		}`,
	},
	{
		name:  "long string",
		line:  line(),
		log:   dummy(),
		input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		want: `{
			"message":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"excerpt":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliq…"
		}`,
	},
	{
		name:  "multiline long string with leading spaces",
		line:  line(),
		log:   dummy(),
		input: " \n \tLorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		want: `{
			"message":" \n \tLorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"excerpt":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliq…"
		}`,
	},
	{
		name:  "multiline long string with leading spaces and multibyte character",
		line:  line(),
		log:   dummy(),
		input: " \n \tLorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt ut labore et dolore magna Ää.",
		want: `{
			"message":" \n \tLorem ipsum dolor sit amet,\nconsectetur adipiscing elit,\nsed do eiusmod tempor incididunt ut labore et dolore magna Ää.",
			"excerpt":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna Ää…"
		}`,
		benchmark: true,
	},
	{
		name: "zero maximum length",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
			Trunc:  0,
		},
		line:  line(),
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name: "without message key names",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{},
		},
		line:  line(),
		input: "Hello, World!",
		want: `{
			"":"Hello, World!"
		}`,
	},
	{
		name: "only original message key name",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		line:  line(),
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name: "explicit byte slice as message excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringBytes("excerpt", []byte("Explicit byte slice")...)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":"Explicit byte slice"
		}`,
	},
	{
		name: "explicit string as message excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringString("excerpt", "Explicit string")},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":"Explicit string"
		}`,
	},
	{
		name: "explicit integer as message excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringInt("excerpt", 42)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":42
		}`,
	},
	{
		name: "explicit float as message excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringFloat32("excerpt", 4.2)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":4.2
		}`,
	},
	{
		name: "explicit boolean as message excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringBool("excerpt", true)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":true
		}`,
	},
	{
		name: "explicit rune slice as messages excerpt key",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV:     []plog.KV{plog.StringRunes("excerpt", []rune("Explicit rune slice")...)},
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:  120,
		},
		input: "Hello, World!",
		want: `{
		  "message": "Hello, World!",
			"excerpt":"Explicit rune slice"
		}`,
	},
	{
		name: `dynamic "time" key`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV: []plog.KV{
				plog.StringFunc("time", func() plog.KV {
					t := time.Date(2020, time.October, 15, 18, 9, 0, 0, time.UTC)
					return pfmt.String(t.String())
				}),
			},
			Keys: [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!",
			"time":"2020-10-15 18:09:00 +0000 UTC"
		}`,
	},
	{
		name: `"standard flag" do not respects file path`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.LstdFlags,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message")},
		},
		input: "path/to/file1:23: Hello, World!",
		want: `{
			"message":"path/to/file1:23: Hello, World!"
		}`,
	},
	{
		name: `"long file" flag respects file path`,
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
			Trunc:  120,
		},
		input: "path/to/file1:23: Hello, World!",
		want: `{
			"message":"path/to/file1:23: Hello, World!",
			"excerpt":"Hello, World!",
			"file":"path/to/file1:23"
		}`,
	},
	{
		name: "replace newline character by whitespace character",
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		input: "Hello,\nWorld!",
		want: `{
			"message":"Hello,\nWorld!",
			"excerpt":"Hello, World!"
		}`,
	},
	{
		name: "remove exclamation marks",
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("!")}},
		},
		input: "Hello, World!!!",
		want: `{
			"message":"Hello, World!!!",
			"excerpt":"Hello, World"
		}`,
	},
	{
		name: `replace word "World" by world "Work"`,
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("World"), []byte("Work")}},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!",
			"excerpt":"Hello, Work!"
		}`,
	},
	{
		name: "ignore pointless replace",
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{[]byte("!"), []byte("!")}},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name: "ignore empty replace",
		line: line(),
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Keys:    [4]encoding.TextMarshaler{pfmt.String("message")},
			Trunc:   120,
			Replace: [][2][]byte{[2][]byte{}},
		},
		input: "Hello, World!",
		want: `{
			"message":"Hello, World!"
		}`,
	},
	{
		name: "file path with empty message",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
			Trunc:  120,
			Marks:  [3][]byte{[]byte("…"), []byte("_EMPTY_")},
		},
		input: "path/to/file1:23:",
		want: `{
			"message":"path/to/file1:23:",
			"excerpt":"_EMPTY_",
			"file":"path/to/file1:23"
		}`,
	},
	{
		name: "file path with blank message",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Flag:   log.Llongfile,
			Keys:   [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
			Trunc:  120,
			Marks:  [3][]byte{[]byte("…"), []byte("_EMPTY_"), []byte("_BLANK_")},
		},
		input: "path/to/file4:56:  ",
		want: `{
			"message":"path/to/file4:56:  ",
			"excerpt":"_BLANK_",
			"file":"path/to/file4:56"
		}`,
	},
	{
		name: "GELF",
		line: line(),
		log: func() plog.Logger {
			l1 := plog.GELF()
			l1.Output = &bytes.Buffer{}
			l := l1.Tee(
				plog.StringString("version", "1.1"),
				plog.StringString("host", "example.tld"),
				plog.StringFunc("timestamp", func() plog.KV {
					t := time.Date(2020, time.October, 15, 18, 9, 0, 0, time.UTC)
					return pfmt.Int64(t.Unix())
				}),
			)
			return l
		}(),
		input: "Hello, GELF!",
		want: `{
			"version":"1.1",
			"short_message":"Hello, GELF!",
			"host":"example.tld",
			"timestamp":1602785340
		}`,
	},
	{
		name: "GELF with file path",
		line: line(),
		log: func() plog.Logger {
			l1 := plog.GELF()
			l1.Output = &bytes.Buffer{}
			l1.Flag = log.Llongfile
			l := l1.Tee(
				plog.StringString("version", "1.1"),
				plog.StringString("host", "example.tld"),
				plog.StringFunc("timestamp", func() plog.KV {
					t := time.Date(2020, time.October, 15, 18, 9, 0, 0, time.UTC)
					return pfmt.Int64(t.Unix())
				}),
			)
			return l
		}(),
		input: "path/to/file7:89: Hello, GELF!",
		want: `{
			"version":"1.1",
			"short_message":"Hello, GELF!",
			"full_message":"path/to/file7:89: Hello, GELF!",
			"host":"example.tld",
			"timestamp":1602785340,
			"_file":"path/to/file7:89"
		}`,
	},
}

func TestFprintWrite(t *testing.T) {
	for _, tt := range FprintWriteTests {
		tt := tt
		t.Run(tt.line+"/io writer via fmt fprint "+tt.name, func(t *testing.T) {
			t.Parallel()

			l, ok := tt.log.(*plog.Log)
			if !ok {
				t.Fatal("unwant logger type")
			}

			var buf bytes.Buffer
			l.Output = &buf

			_, err := fmt.Fprint(l, tt.input)
			if err != nil {
				t.Fatalf("unwant write error: %s", err)
			}

			ja := jsonassert.New(testprinter{t: t, link: tt.line})
			ja.Assertf(buf.String(), tt.want)
		})
	}
}

func BenchmarkPlog(b *testing.B) {
	for _, tt := range WriteTests {
		if !tt.benchmark {
			continue
		}
		b.Run(tt.line+"/io.Writer", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l := tt.log.Tee(tt.kv...)
				_, err := l.Write(tt.input)
				if err != nil {
					fmt.Println(err)
				}
				l.Close()
			}
		})
	}

	for _, tt := range FprintWriteTests {
		if !tt.benchmark {
			continue
		}
		b.Run(tt.line+"/fmt.Fprint io.Writer", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := fmt.Fprint(tt.log, tt.input)
				if err != nil {
					fmt.Println(err)
				}
			}
		})
	}
}

var dummy = func() plog.Logger {
	return &plog.Log{
		Output:  &bytes.Buffer{},
		Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt"), pfmt.String("trail"), pfmt.String("file")},
		Key:     plog.Original,
		Trunc:   120,
		Marks:   [3][]byte{[]byte("…"), []byte("_EMPTY_"), []byte("_BLANK_")},
		Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
	}
}

func TestLogWriteTrailingNewLine(t *testing.T) {
	var buf bytes.Buffer

	l := &plog.Log{Output: &buf}

	_, err := l.Write([]byte("Hello, Wrold!"))
	if err != nil {
		t.Fatalf("unwant write error: %s", err)
	}

	if buf.Bytes()[len(buf.Bytes())-1] != '\n' {
		t.Errorf("trailing new line want but not present: %q", buf.String())
	}
}

var TruncateTests = []struct {
	name      string
	line      string
	log       plog.Logger
	input     []byte
	want      []byte
	benchmark bool
}{
	{
		name: "do nothing",
		log: &plog.Log{
			Output: &bytes.Buffer{},
		},
		line:  line(),
		input: []byte("Hello,\nWorld!"),
		want:  []byte("Hello,\nWorld!"),
	},
	{
		name: "truncate last character",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Trunc:  12,
		},
		line:  line(),
		input: []byte("Hello, World!"),
		want:  []byte("Hello, World"),
	},
	{
		name: "truncate last character and places ellipsis instead",
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Trunc:  12,
			Marks:  [3][]byte{[]byte("…")},
		},
		line:  line(),
		input: []byte("Hello, World!"),
		want:  []byte("Hello, World…"),
	},
	{
		name: "replace new lines by spaces",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
		},
		line:  line(),
		input: []byte("Hello\n,\nWorld\n!"),
		want:  []byte("Hello , World !"),
	},
	{
		name: "replace new lines by empty string",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte("")}},
		},
		line:  line(),
		input: []byte("Hello\n,\nWorld\n!"),
		want:  []byte("Hello,World!"),
	},
	{
		name: "remove new lines",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("\n")}},
		},
		line:  line(),
		input: []byte("Hello\n,\nWorld\n!"),
		want:  []byte("Hello,World!"),
	},
	{
		name: "replace three characters by one",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("foo"), []byte("f")}, [2][]byte{[]byte("bar"), []byte("b")}},
		},
		line:  line(),
		input: []byte("foobar"),
		want:  []byte("fb"),
	},
	{
		name: "replace one characters by three",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("f"), []byte("foo")}, [2][]byte{[]byte("b"), []byte("bar")}},
		},
		line:  line(),
		input: []byte("fb"),
		want:  []byte("foobar"),
	},
	{
		name: "remove three characters",
		log: &plog.Log{
			Output:  &bytes.Buffer{},
			Replace: [][2][]byte{[2][]byte{[]byte("foo")}, [2][]byte{[]byte("bar")}},
		},
		line:  line(),
		input: []byte("foobar foobar"),
		want:  []byte(" "),
	},
}

func TestTruncate(t *testing.T) {
	for _, tt := range TruncateTests {
		tt := tt
		t.Run(tt.line+"/truncate "+tt.name, func(t *testing.T) {
			t.Parallel()

			l0, ok := tt.log.(*plog.Log)
			if !ok {
				t.Fatal("unwant logger type")
			}

			n := len(tt.input) + 10*10
			for _, m := range l0.Marks {
				if n < len(m) {
					n = len(m)
				}
			}

			excerpt := make([]byte, n)

			n, err := l0.Truncate(excerpt, tt.input)
			if err != nil {
				t.Fatalf("unwant write error: %s", err)
			}

			excerpt = excerpt[:n]

			if !bytes.Equal(excerpt, tt.want) {
				t.Errorf("unwant excerpt, want: %q, get: %q %s", tt.want, excerpt, tt.line)
			}
		})
	}
}

var NewTests = []struct {
	name      string
	line      string
	log       plog.Logger
	kv        []plog.KV
	want      string
	benchmark bool
}{
	{
		name: "one kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV: []plog.KV{
				plog.StringString("foo", "bar"),
			},
		},
		want: `{
			"foo":"bar"
		}`,
	},
	{
		name: "two kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV: []plog.KV{
				plog.StringString("foo", "bar"),
				plog.StringString("baz", "xyz"),
			},
		},
		want: `{
			"foo":"bar",
			"baz":"xyz"
		}`,
	},
	{
		name: "one additional kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
		},
		kv: []plog.KV{
			plog.StringString("baz", "xyz"),
		},
		want: `{
			"baz":"xyz"
		}`,
	},
	{
		name: "two additional kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
		},
		kv: []plog.KV{
			plog.StringString("foo", "bar"),
			plog.StringString("baz", "xyz"),
		},
		want: `{
			"foo":"bar",
			"baz":"xyz"
		}`,
	},
	{
		name: "one kv with additional one kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV: []plog.KV{
				plog.StringString("foo", "bar"),
			},
		},
		kv: []plog.KV{
			plog.StringString("baz", "xyz"),
		},
		want: `{
			"foo":"bar",
			"baz":"xyz"
		}`,
	},
	{
		name: "two kv with two additional kv",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			KV: []plog.KV{
				plog.StringString("foo", "bar"),
				plog.StringString("abc", "dfg"),
			},
		},
		kv: []plog.KV{
			plog.StringString("baz", "xyz"),
			plog.StringString("hjk", "lmn"),
		},
		want: `{
			"foo":"bar",
			"abc":"dfg",
			"baz":"xyz",
			"hjk":"lmn"
		}`,
	},
}

func TestNew(t *testing.T) {
	for _, tt := range NewTests {
		tt := tt
		t.Run(tt.line+"/with "+tt.name, func(t *testing.T) {
			t.Parallel()

			l0, ok := tt.log.(*plog.Log)
			if !ok {
				t.Fatal("unwant logger type")
			}

			buf, ok := l0.Output.(*bytes.Buffer)
			if !ok {
				t.Fatal("unwant output type")
			}

			*buf = bytes.Buffer{}

			l := tt.log.Tee(tt.kv...)
			defer l.Close()

			_, err := l.Write(nil)
			if err != nil {
				t.Fatalf("unwant write error: %s", err)
			}

			ja := jsonassert.New(testprinter{t: t, link: tt.line})
			ja.Assertf(buf.String(), tt.want)
		})
	}
}

var SeverityLevelTests = []struct {
	name      string
	line      string
	log       plog.Logger
	levels    []string
	kv        []plog.KV
	want      string
	benchmark bool
}{
	{
		name: "just level 7",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Level:  func(level string) io.Writer { return nil },
		},
		levels: []string{"7"},
		want: `{
			"level":"7"
		}`,
	},
	{
		name: "just level 7 next to the key-value pair",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Level:  func(level string) io.Writer { return nil },
			KV: []plog.KV{
				plog.StringString("foo", "bar"),
			},
		},
		levels: []string{"7"},
		want: `{
			"foo":"bar",
			"level":"7"
		}`,
	},
	{
		name: "two consecutive levels call 7 and 6",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Level:  func(level string) io.Writer { return nil },
		},
		levels: []string{"7", "6"},
		want: `{
			"level":"6"
		}`,
	},
	{
		name: "level 7 overwrites level 42 from key-value pair",
		line: line(),
		log: &plog.Log{
			Output: &bytes.Buffer{},
			Level:  func(level string) io.Writer { return nil },
			KV: []plog.KV{
				plog.StringString("level", "42"),
			},
		},
		levels: []string{"7"},
		want: `{
			"level":"7"
		}`,
	},
}

func TestSeverityLevel(t *testing.T) {
	for _, tt := range SeverityLevelTests {
		tt := tt
		t.Run(tt.line+"/with "+tt.name, func(t *testing.T) {
			t.Parallel()

			l0, ok := tt.log.(*plog.Log)
			if !ok {
				t.Fatal("unwant logger type")
			}

			buf, ok := l0.Output.(*bytes.Buffer)
			if !ok {
				t.Fatal("unwant output type")
			}

			*buf = bytes.Buffer{}

			l := tt.log.Tee(tt.kv...)
			defer l.Close()

			for _, sev := range tt.levels {
				l = l.Tee(plog.StringLevel("level", sev))
			}

			_, err := l.Write(nil)
			if err != nil {
				t.Fatalf("unwant write error: %s", err)
			}

			ja := jsonassert.New(testprinter{t: t, link: tt.line})
			ja.Assertf(buf.String(), tt.want)
		})
	}
}

func TestEncode(t *testing.T) {
	l0 := &plog.Log{
		Output:  &bytes.Buffer{},
		Keys:    [4]encoding.TextMarshaler{pfmt.String("message"), pfmt.String("excerpt")},
		Marks:   [3][]byte{[]byte("…")},
		Trunc:   12,
		Replace: [][2][]byte{[2][]byte{[]byte("\n"), []byte(" ")}},
	}

	ja := jsonassert.New(testprinter{t: t, link: line()})

	l := l0.Tee(plog.StringString("foo", "bar"))

	ja.Assertf(string(l.(plog.Encoder).Encode()), `{"foo":"bar"}`)

	ja.Assertf(
		string(l.(plog.Encoder).Encode(plog.StringString("greeting", "Hello,\nWorld!"))),
		`{"foo":"bar","greeting":"Hello,\nWorld!"}`,
	)
}
