// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/pprint/plog"
)

var MarshalBytespTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"bytes pointer": plog.Bytesp(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"empty bytes pointer": plog.Bytesp(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty bytes pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil bytes pointer": plog.Bytesp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil bytes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"any bytes pointer": plog.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"any empty bytes pointer": plog.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect bytes pointer": plog.Reflect(&p)}
		}(),
		want:     "SGVsbG8sIFfDtnJsZCE=",
		wantText: "SGVsbG8sIFfDtnJsZCE=",
		wantJSON: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []byte{}
			return map[string]json.Marshaler{"reflect empty bytes pointer": plog.Reflect(&p)}
		}(),
		want: "",
		wantJSON: `{
			"reflect empty bytes pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"complex128": plog.Complex128(complex(1, 23))},
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any complex128": plog.Any(complex(1, 23))},
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"reflect complex128": plog.Reflect(complex(1, 23))},
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"complex128 pointer": plog.Complex128p(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil complex128 pointer": plog.Complex128p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"any complex128 pointer": plog.Any(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex128 = complex(1, 23)
			return map[string]json.Marshaler{"reflect complex128 pointer": plog.Reflect(&c)}
		}(),
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"complex64": plog.Complex64(complex(3, 21))},
		want:     "3+21i",
		wantText: "3+21i",
		wantJSON: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any complex64": plog.Any(complex(3, 21))},
		want:     "3+21i",
		wantText: "3+21i",
		wantJSON: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"reflect complex64": plog.Reflect(complex(3, 21))},
		want:      "(3+21i)",
		wantText:  "(3+21i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"error": plog.Error(errors.New("something went wrong"))},
		want:     "something went wrong",
		wantText: "something went wrong",
		wantJSON: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil error": plog.Error(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil error":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"errors": plog.Errors(errors.New("something went wrong"), errors.New("wrong"))},
		want:     "something went wrong wrong",
		wantText: "something went wrong wrong",
		wantJSON: `{
			"errors":["something went wrong","wrong"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil errors": plog.Errors(nil, nil)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"nil errors":[null,null]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without errors": plog.Errors()},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"without errors":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any error": plog.Any(errors.New("something went wrong"))},
		want:     "something went wrong",
		wantText: "something went wrong",
		wantJSON: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect error": plog.Reflect(errors.New("something went wrong"))},
		want:     "{something went wrong}",
		wantText: "{something went wrong}",
		wantJSON: `{
			"reflect error":{}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"complex64 pointer": plog.Complex64p(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil complex64 pointer": plog.Complex64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"any complex64 pointer": plog.Any(&c)}
		}(),
		want:     "1+23i",
		wantText: "1+23i",
		wantJSON: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var c complex64 = complex(1, 23)
			return map[string]json.Marshaler{"reflect complex64 pointer": plog.Reflect(&c)}
		}(),
		want:      "(1+23i)",
		wantText:  "(1+23i)",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float32": plog.Float32(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float32":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"high precision float32": plog.Float32(0.123456789)},
		want:     "0.12345679",
		wantText: "0.12345679",
		wantJSON: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"zero float32": plog.Float32(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"zero float32":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any float32": plog.Any(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float32":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any zero float32": plog.Any(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"any zero float32":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect float32": plog.Reflect(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect zero float32": plog.Reflect(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"reflect zero float32":0
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"float32 pointer": plog.Float32p(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 0.123456789
			return map[string]json.Marshaler{"high precision float32 pointer": plog.Float32p(&f)}
		}(),
		want:     "0.12345679",
		wantText: "0.12345679",
		wantJSON: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float32 nil pointer": plog.Float32p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"any float32 pointer": plog.Any(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float32 = 4.2
			return map[string]json.Marshaler{"reflect float32 pointer": plog.Reflect(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float32
			return map[string]json.Marshaler{"reflect float32 pointer to nil": plog.Reflect(f)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect float32 pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float64": plog.Float64(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float64":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"high precision float64": plog.Float64(0.123456789)},
		want:     "0.123456789",
		wantText: "0.123456789",
		wantJSON: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"zero float64": plog.Float64(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"zero float64":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any float64": plog.Any(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float64":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any zero float64": plog.Any(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"any zero float64":0
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect float64": plog.Reflect(4.2)},
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect zero float64": plog.Reflect(0)},
		want:     "0",
		wantText: "0",
		wantJSON: `{
			"reflect zero float64":0
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"float64 pointer": plog.Float64p(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 0.123456789
			return map[string]json.Marshaler{"high precision float64 pointer": plog.Float64p(&f)}
		}(),
		want:     "0.123456789",
		wantText: "0.123456789",
		wantJSON: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"float64 nil pointer": plog.Float64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"any float64 pointer": plog.Any(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f float64 = 4.2
			return map[string]json.Marshaler{"reflect float64 pointer": plog.Reflect(&f)}
		}(),
		want:     "4.2",
		wantText: "4.2",
		wantJSON: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var f *float64
			return map[string]json.Marshaler{"reflect float64 pointer to nil": plog.Reflect(f)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect float64 pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int": plog.Int(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"int pointer": plog.Intp(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"any int pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int = 42
			return map[string]json.Marshaler{"reflect int pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int16": plog.Int16(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int16": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int16": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int16":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"int16 pointer": plog.Int16p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"any int16 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int16 = 42
			return map[string]json.Marshaler{"reflect int16 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int16 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int32": plog.Int32(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int32": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int32": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int32":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"int32 pointer": plog.Int32p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"any int32 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int32 = 42
			return map[string]json.Marshaler{"reflect int32 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int32 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int64": plog.Int64(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int64": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int64": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int64":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"int64 pointer": plog.Int64p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"any int64 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int64 = 42
			return map[string]json.Marshaler{"reflect int64 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int64 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"int8": plog.Int8(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any int8": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect int8": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int8":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"int8 pointer": plog.Int8p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"any int8 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i int8 = 42
			return map[string]json.Marshaler{"reflect int8 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect int8 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"runes": plog.Runes([]rune("Hello, Wörld!")...)},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty runes": plog.Runes([]rune{}...)},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var p []rune
			return map[string]json.Marshaler{"nil runes": plog.Runes(p...)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil runes":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"rune slice with zero rune": plog.Runes([]rune{rune(0)}...)},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any runes": plog.Any([]rune("Hello, Wörld!"))},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any empty runes": plog.Any([]rune{})},
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty runes":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any rune slice with zero rune": plog.Any([]rune{rune(0)})},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect runes": plog.Reflect([]rune("Hello, Wörld!"))},
		want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantJSON: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect empty runes": plog.Reflect([]rune{})},
		want:     "[]",
		wantText: "[]",
		wantJSON: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect rune slice with zero rune": plog.Reflect([]rune{rune(0)})},
		want:     "[0]",
		wantText: "[0]",
		wantJSON: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"runes pointer": plog.Runesp(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"empty runes pointer": plog.Runesp(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil runes pointer": plog.Runesp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"any runes pointer": plog.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"any empty runes pointer": plog.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune("Hello, Wörld!")
			return map[string]json.Marshaler{"reflect runes pointer": plog.Reflect(&p)}
		}(),
		want:     "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantText: "[72 101 108 108 111 44 32 87 246 114 108 100 33]",
		wantJSON: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := []rune{}
			return map[string]json.Marshaler{"reflect empty runes pointer": plog.Reflect(&p)}
		}(),
		want:     "[]",
		wantText: "[]",
		wantJSON: `{
			"reflect empty runes pointer":[]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"string": plog.String("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty string": plog.String("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"string with zero byte": plog.String(string(byte(0)))},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"strings": plog.Strings("Hello, Wörld!", "Hello, World!")},
		want:     "Hello, Wörld! Hello, World!",
		wantText: "Hello, Wörld! Hello, World!",
		wantJSON: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty strings": plog.Strings("", "")},
		want:     " ",
		wantText: " ",
		wantJSON: `{
			"empty strings":["",""]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"strings with zero byte": plog.Strings(string(byte(0)), string(byte(0)))},
		want:     "\\u0000 \\u0000",
		wantText: "\\u0000 \\u0000",
		wantJSON: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"without strings": plog.Strings()},
		want:     "",
		wantText: "",
		wantJSON: `{
			"without strings":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any string": plog.Any("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any empty string": plog.Any("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any string with zero byte": plog.Any(string(byte(0)))},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect string": plog.Reflect("Hello, Wörld!")},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect empty string": plog.Reflect("")},
		want:     "",
		wantText: "",
		wantJSON: `{
			"reflect empty string":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect string with zero byte": plog.Reflect(string(byte(0)))},
		want:     "\u0000",
		wantText: "\u0000",
		wantJSON: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"string pointer": plog.Stringp(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"empty string pointer": plog.Stringp(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty string pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil string pointer": plog.Stringp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"any string pointer": plog.Any(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"any empty string pointer": plog.Any(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := "Hello, Wörld!"
			return map[string]json.Marshaler{"reflect string pointer": plog.Reflect(&p)}
		}(),
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			p := ""
			return map[string]json.Marshaler{"reflect empty string pointer": plog.Reflect(&p)}
		}(),
		want:     "",
		wantText: "",
		wantJSON: `{
			"reflect empty string pointer":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text": plog.Text(plog.String("Hello, Wörld!"))},
		want:     "Hello, Wörld!",
		wantText: "Hello, Wörld!",
		wantJSON: `{
			"text":"Hello, Wörld!"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"empty text": plog.Text(plog.String(""))},
		want:     "",
		wantText: "",
		wantJSON: `{
			"empty text":""
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"text with zero byte": plog.Text(plog.String(string(byte(0))))},
		want:     "\\u0000",
		wantText: "\\u0000",
		wantJSON: `{
			"text with zero byte":"\u0000"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint": plog.Uint(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"uint pointer": plog.Uintp(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint pointer": plog.Uintp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"any uint pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint = 42
			return map[string]json.Marshaler{"reflect uint pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint16": plog.Uint16(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint16": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint16":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint16": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint16":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"uint16 pointer": plog.Uint16p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint16 pointer": plog.Uint16p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"any uint16 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint16 = 42
			return map[string]json.Marshaler{"reflect uint16 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i *uint16
			return map[string]json.Marshaler{"reflect uint16 pointer to nil": plog.Reflect(i)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint32": plog.Uint32(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint32": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint32":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint32": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint32":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint32 = 42
			return map[string]json.Marshaler{"uint32 pointer": plog.Uint32p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint32 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint32 pointer": plog.Uint32p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint32 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint32 = 42
			return map[string]json.Marshaler{"any uint32 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint32 = 42
			return map[string]json.Marshaler{"reflect uint32 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint32 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint64": plog.Uint64(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint64": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint64":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint64": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint64":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"uint64 pointer": plog.Uint64p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint64 pointer": plog.Uint64p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"any uint64 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint64 = 42
			return map[string]json.Marshaler{"reflect uint64 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint64 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uint8": plog.Uint8(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uint8": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint8":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uint8": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint8":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"uint8 pointer": plog.Uint8p(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uint8 pointer": plog.Uint8p(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"any uint8 pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uint8 = 42
			return map[string]json.Marshaler{"reflect uint8 pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uint8 pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"uintptr": plog.Uintptr(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any uintptr": plog.Any(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect uintptr": plog.Reflect(42)},
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uintptr":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uintptr = 42
			return map[string]json.Marshaler{"uintptr pointer": plog.Uintptrp(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"uintptr pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil uintptr pointer": plog.Uintptrp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil uintptr pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uintptr = 42
			return map[string]json.Marshaler{"any uintptr pointer": plog.Any(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"any uintptr pointer":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var i uintptr = 42
			return map[string]json.Marshaler{"reflect uintptr pointer": plog.Reflect(&i)}
		}(),
		want:     "42",
		wantText: "42",
		wantJSON: `{
			"reflect uintptr pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"time": time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)},
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01T00:00:00.000000042Z",
		wantJSON: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any time": plog.Any(time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC))},
		want:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
		wantText: `1970-01-01T00:00:00.000000042Z`,
		wantJSON: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect time": plog.Reflect(time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC))},
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantJSON: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"time pointer": &t}
		}(),
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01T00:00:00.000000042Z",
		wantJSON: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var t time.Time
			return map[string]json.Marshaler{"nil time pointer": t}
		}(),
		want:     "0001-01-01 00:00:00 +0000 UTC",
		wantText: "0001-01-01T00:00:00Z",
		wantJSON: `{
			"nil time pointer":"0001-01-01T00:00:00Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"any time pointer": plog.Any(&t)}
		}(),
		want:     `1970-01-01 00:00:00.000000042 +0000 UTC`,
		wantText: `1970-01-01T00:00:00.000000042Z`,
		wantJSON: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return map[string]json.Marshaler{"reflect time pointer": plog.Reflect(&t)}
		}(),
		want:     "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantText: "1970-01-01 00:00:00.000000042 +0000 UTC",
		wantJSON: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"duration": plog.Duration(42 * time.Nanosecond)},
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"duration":"42ns"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any duration": plog.Any(42 * time.Nanosecond)},
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect duration": plog.Reflect(42 * time.Nanosecond)},
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"reflect duration":42
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"duration pointer": plog.Durationp(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil duration pointer": plog.Durationp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"any duration pointer": plog.Any(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			d := 42 * time.Nanosecond
			return map[string]json.Marshaler{"reflect duration pointer": plog.Reflect(&d)}
		}(),
		want:     "42ns",
		wantText: "42ns",
		wantJSON: `{
			"reflect duration pointer":42
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any struct": plog.Any(Struct{Name: "John Doe", Age: 42})},
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			s := Struct{Name: "John Doe", Age: 42}
			return map[string]json.Marshaler{"any struct pointer": plog.Any(&s)}
		}(),
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"struct reflect": plog.Reflect(Struct{Name: "John Doe", Age: 42})},
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			s := Struct{Name: "John Doe", Age: 42}
			return map[string]json.Marshaler{"struct reflect pointer": plog.Reflect(&s)}
		}(),
		want:     "{John Doe 42}",
		wantText: "{John Doe 42}",
		wantJSON: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"raw json": plog.Raw([]byte(`{"foo":"bar"}`))},
		want:     `{"foo":"bar"}`,
		wantText: `{"foo":"bar"}`,
		wantJSON: `{
			"raw json":{"foo":"bar"}
		}`,
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw malformed json object": plog.Raw([]byte(`xyz{"foo":"bar"}`))},
		want:      `xyz{"foo":"bar"}`,
		wantText:  `xyz{"foo":"bar"}`,
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw malformed json key/value": plog.Raw([]byte(`{"foo":"bar""}`))},
		want:      `{"foo":"bar""}`,
		wantText:  `{"foo":"bar""}`,
		wantError: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:      line(),
		input:     map[string]json.Marshaler{"raw json with unescaped null byte": plog.Raw(append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...))},
		want:      "{\"foo\":\"\u0000xyz\"}",
		wantText:  "{\"foo\":\"\u0000xyz\"}",
		wantError: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"raw nil": plog.Raw(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"raw nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any byte array": plog.Any([3]byte{'f', 'o', 'o'})},
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"any byte array pointer": plog.Any(&a)}
		}(),
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"any byte array pointer to nil": plog.Any(a)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"any byte array pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect byte array": plog.Reflect([3]byte{'f', 'o', 'o'})},
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			a := [3]byte{'f', 'o', 'o'}
			return map[string]json.Marshaler{"reflect byte array pointer": plog.Reflect(&a)}
		}(),
		want:     "[102 111 111]",
		wantText: "[102 111 111]",
		wantJSON: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var a *[3]byte
			return map[string]json.Marshaler{"reflect byte array pointer to nil": plog.Reflect(a)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"any untyped nil": plog.Any(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"any untyped nil":null
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"reflect untyped nil": plog.Reflect(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect untyped nil":null
		}`,
	},
}

func TestMarshalBytesp(t *testing.T) {
	testMarshal(t, MarshalBytespTests)
}
