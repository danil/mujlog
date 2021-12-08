// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/pprint/plog"
)

var KVTests = []struct {
	line      string
	input     plog.KV
	want      string
	error     error
	benchmark bool
}{
	{
		line:  line(),
		input: plog.StringBool("bool true", true),
		want: `{
			"bool true":true
		}`,
	},
	{
		line:  line(),
		input: plog.StringBool("bool false", false),
		want: `{
			"bool false":false
		}`,
	},
	{
		line:  line(),
		input: plog.StringBools("bools true false", true, false),
		want: `{
			"bools true false":[true,false]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBools("without bools"),
		want: `{
			"without bools":[]
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any bool false", false),
		want: `{
			"any bool false":false
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect bool false", false),
		want: `{
			"reflect bool false":false
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.StringBoolp("bool pointer to true", &b)
		}(),
		want: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := false
			return plog.StringBoolp("bool pointer to false", &b)
		}(),
		want: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:  line(),
		input: plog.StringBoolp("bool nil pointer", nil),
		want: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.StringAny("any bool pointer to true", &b)
		}(),
		want: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			b2 := &b
			return plog.StringAny("any twice/nested pointer to bool true", &b2)
		}(),
		want: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.StringReflect("reflect bool pointer to true", &b)
		}(),
		want: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			b2 := &b
			return plog.StringReflect("reflect bool twice/nested pointer to true", &b2)
		}(),
		want: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var b *bool
			return plog.StringReflect("reflect bool pointer to nil", b)
		}(),
		want: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes", []byte("Hello, Wörld!")...),
		want: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes with quote", []byte(`Hello, "World"!`)...),
		want: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes quote", []byte(`"Hello, World!"`)...),
		want: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes nested quote", []byte(`"Hello, "World"!"`)...),
		want: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes json", []byte(`{"foo":"bar"}`)...),
		want: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("bytes json quote", []byte(`"{"foo":"bar"}"`)...),
		want: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytes("empty bytes", []byte{}...),
		want: `{
			"empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var p []byte
			return plog.StringBytes("nil bytes", p...)
		}(),
		want: `{
			"nil bytes":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of byte slices", []byte("Hello, Wörld!"), []byte("Hello, World!")),
		want: `{
			"slice of byte slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of byte slices with quote", []byte(`Hello, "Wörld"!`), []byte(`Hello, "World"!`)),
		want: `{
			"slice of byte slices with quote":["Hello, \"Wörld\"!","Hello, \"World\"!"]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("quoted slice of byte slices", []byte(`"Hello, Wörld!"`), []byte(`"Hello, World!"`)),
		want: `{
			"quoted slice of byte slices":["\"Hello, Wörld!\"","\"Hello, World!\""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of byte slices with nested quote", []byte(`"Hello, "Wörld"!"`), []byte(`"Hello, "World"!"`)),
		want: `{
			"slice of byte slices with nested quote":["\"Hello, \"Wörld\"!\"","\"Hello, \"World\"!\""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of byte slices with json", []byte(`{"foo":"bar"}`), []byte(`{"baz":"xyz"}`)),
		want: `{
			"slice of byte slices with json":["{\"foo\":\"bar\"}","{\"baz\":\"xyz\"}"]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of byte slices with quoted json", []byte(`"{"foo":"bar"}"`), []byte(`"{"baz":"xyz"}"`)),
		want: `{
			"slice of byte slices with quoted json":["\"{\"foo\":\"bar\"}\"","\"{\"baz\":\"xyz\"}\""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of empty byte slices", []byte{}, []byte{}),
		want: `{
			"slice of empty byte slices":["",""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("slice of nil byte slices", nil, nil),
		want: `{
			"slice of nil byte slices":[null,null]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytess("empty slice of nil slices"),
		want: `{
			"empty slice of nil slices":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any bytes", []byte("Hello, Wörld!")),
		want: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any empty bytes", []byte{}),
		want: `{
			"any empty bytes":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect bytes", []byte("Hello, Wörld!")),
		want: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect empty bytes", []byte{}),
		want: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.StringBytesp("bytes pointer", &p)
		}(),
		want: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.StringBytesp("empty bytes pointer", &p)
		}(),
		want: `{
			"empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytesp("nil bytes pointer", nil),
		want: `{
			"nil bytes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p, p2 := []byte("Hello, Wörld!"), []byte("Hello, World!")
			return plog.StringBytessp("slice of byte pointer slices", &p, &p2)
		}(),
		want: `{
			"slice of byte pointer slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p, p2 := []byte{}, []byte{}
			return plog.StringBytessp("slice of empty byte pointer slices", &p, &p2)
		}(),
		want: `{
			"slice of empty byte pointer slices":["",""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytessp("slice of nil byte pointer slices", nil, nil),
		want: `{
			"slice of nil byte pointer slices":[null,null]
		}`,
	},
	{
		line:  line(),
		input: plog.StringBytessp("empty slice of byte pointer slices"),
		want: `{
			"empty slice of byte pointer slices":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.StringAny("any bytes pointer", &p)
		}(),
		want: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.StringAny("any empty bytes pointer", &p)
		}(),
		want: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.StringReflect("reflect bytes pointer", &p)
		}(),
		want: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.StringReflect("reflect empty bytes pointer", &p)
		}(),
		want: `{
			"reflect empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringComplex128("complex128", complex(1, 23)),
		want: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any complex128", complex(1, 23)),
		want: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect complex128", complex(1, 23)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.StringComplex128p("complex128 pointer", &c)
		}(),
		want: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringComplex128p("nil complex128 pointer", nil),
		want: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.StringAny("any complex128 pointer", &c)
		}(),
		want: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.StringReflect("reflect complex128 pointer", &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: plog.StringComplex64("complex64", complex(3, 21)),
		want: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any complex64", complex(3, 21)),
		want: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect complex64", complex(3, 21)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: plog.StringError("error", errors.New("something went wrong")),
		want: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: plog.StringError("nil error", nil),
		want: `{
			"nil error":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringErrors("errors", errors.New("something went wrong"), errors.New("wrong")),
		want: `{
			"errors":["something went wrong","wrong"]
		}`,
	},
	{
		line:  line(),
		input: plog.StringErrors("nil errors", nil, nil),
		want: `{
			"nil errors":[null,null]
		}`,
	},
	{
		line:  line(),
		input: plog.StringErrors("without errors"),
		want: `{
			"without errors":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any error", errors.New("something went wrong")),
		want: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect error", errors.New("something went wrong")),
		want: `{
			"reflect error":{}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.StringComplex64p("complex64 pointer", &c)
		}(),
		want: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.StringComplex64p("nil complex64 pointer", nil),
		want: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.StringAny("any complex64 pointer", &c)
		}(),
		want: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.StringReflect("reflect complex64 pointer", &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
	{
		line:  line(),
		input: plog.StringFloat32("float32", 4.2),
		want: `{
			"float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat32("high precision float32", 0.123456789),
		want: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat32("zero float32", 0),
		want: `{
			"zero float32":0
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any float32", 4.2),
		want: `{
			"any float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any zero float32", 0),
		want: `{
			"any zero float32":0
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect float32", 4.2),
		want: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect zero float32", 0),
		want: `{
			"reflect zero float32":0
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.StringFloat32p("float32 pointer", &f)
		}(),
		want: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 0.123456789
			return plog.StringFloat32p("high precision float32 pointer", &f)
		}(),
		want: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat32p("float32 nil pointer", nil),
		want: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.StringAny("any float32 pointer", &f)
		}(),
		want: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.StringReflect("reflect float32 pointer", &f)
		}(),
		want: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f *float32
			return plog.StringReflect("reflect float32 pointer to nil", f)
		}(),
		want: `{
			"reflect float32 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat64("float64", 4.2),
		want: `{
			"float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat64("high precision float64", 0.123456789),
		want: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat64("zero float64", 0),
		want: `{
			"zero float64":0
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any float64", 4.2),
		want: `{
			"any float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any zero float64", 0),
		want: `{
			"any zero float64":0
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect float64", 4.2),
		want: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect zero float64", 0),
		want: `{
			"reflect zero float64":0
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.StringFloat64p("float64 pointer", &f)
		}(),
		want: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 0.123456789
			return plog.StringFloat64p("high precision float64 pointer", &f)
		}(),
		want: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.StringFloat64p("float64 nil pointer", nil),
		want: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.StringAny("any float64 pointer", &f)
		}(),
		want: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.StringReflect("reflect float64 pointer", &f)
		}(),
		want: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f *float64
			return plog.StringReflect("reflect float64 pointer to nil", f)
		}(),
		want: `{
			"reflect float64 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringInt("int", 42),
		want: `{
			"int":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any int", 42),
		want: `{
			"any int":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect int", 42),
		want: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.StringIntp("int pointer", &i)
		}(),
		want: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.StringAny("any int pointer", &i)
		}(),
		want: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.StringReflect("reflect int pointer", &i)
		}(),
		want: `{
			"reflect int pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringInt16("int16", 42),
		want: `{
			"int16":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any int16", 42),
		want: `{
			"any int16":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect int16", 42),
		want: `{
			"reflect int16":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.StringInt16p("int16 pointer", &i)
		}(),
		want: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.StringAny("any int16 pointer", &i)
		}(),
		want: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.StringReflect("reflect int16 pointer", &i)
		}(),
		want: `{
			"reflect int16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringInt32("int32", 42),
		want: `{
			"int32":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any int32", 42),
		want: `{
			"any int32":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect int32", 42),
		want: `{
			"reflect int32":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.StringInt32p("int32 pointer", &i)
		}(),
		want: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.StringAny("any int32 pointer", &i)
		}(),
		want: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.StringReflect("reflect int32 pointer", &i)
		}(),
		want: `{
			"reflect int32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringInt64("int64", 42),
		want: `{
			"int64":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any int64", 42),
		want: `{
			"any int64":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect int64", 42),
		want: `{
			"reflect int64":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.StringInt64p("int64 pointer", &i)
		}(),
		want: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.StringAny("any int64 pointer", &i)
		}(),
		want: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.StringReflect("reflect int64 pointer", &i)
		}(),
		want: `{
			"reflect int64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringInt8("int8", 42),
		want: `{
			"int8":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any int8", 42),
		want: `{
			"any int8":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect int8", 42),
		want: `{
			"reflect int8":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.StringInt8p("int8 pointer", &i)
		}(),
		want: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.StringAny("any int8 pointer", &i)
		}(),
		want: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.StringReflect("reflect int8 pointer", &i)
		}(),
		want: `{
			"reflect int8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringRunes("runes", []rune("Hello, Wörld!")...),
		want: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringRunes("empty runes", []rune{}...),
		want: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var p []rune
			return plog.StringRunes("nil runes", p...)
		}(),
		want: `{
			"nil runes":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringRunes("rune slice with zero rune", []rune{rune(0)}...),
		want: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any runes", []rune("Hello, Wörld!")),
		want: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any empty runes", []rune{}),
		want: `{
			"any empty runes":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any rune slice with zero rune", []rune{rune(0)}),
		want: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect runes", []rune("Hello, Wörld!")),
		want: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect empty runes", []rune{}),
		want: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect rune slice with zero rune", []rune{rune(0)}),
		want: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.StringRunesp("runes pointer", &p)
		}(),
		want: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.StringRunesp("empty runes pointer", &p)
		}(),
		want: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringRunesp("nil runes pointer", nil),
		want: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.StringAny("any runes pointer", &p)
		}(),
		want: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.StringAny("any empty runes pointer", &p)
		}(),
		want: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.StringReflect("reflect runes pointer", &p)
		}(),
		want: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.StringReflect("reflect empty runes pointer", &p)
		}(),
		want: `{
			"reflect empty runes pointer":[]
		}`,
	},
	{
		line:  line(),
		input: plog.StringString("string", "Hello, Wörld!"),
		want: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringString("empty string", ""),
		want: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringString("string with zero byte", string(byte(0))),
		want: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.StringStrings("strings", "Hello, Wörld!", "Hello, World!"),
		want: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:  line(),
		input: plog.StringStrings("nil strings"),
		want: `{
			"nil strings":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringStrings("empty strings", "", ""),
		want: `{
			"empty strings":["",""]
		}`,
	},
	{
		line:  line(),
		input: plog.StringStrings("without strings"),
		want: `{
			"without strings":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringStrings("strings with zero byte", string(byte(0)), string(byte(0))),
		want: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
	},

	{
		line:  line(),
		input: plog.StringAny("any string", "Hello, Wörld!"),
		want: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any empty string", ""),
		want: `{
			"any empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any string with zero byte", string(byte(0))),
		want: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect string", "Hello, Wörld!"),
		want: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect empty string", ""),
		want: `{
			"reflect empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect string with zero byte", string(byte(0))),
		want: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.StringStringp("string pointer", &p)
		}(),
		want: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.StringStringp("empty string pointer", &p)
		}(),
		want: `{
			"empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringStringp("nil string pointer", nil),
		want: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.StringAny("any string pointer", &p)
		}(),
		want: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.StringAny("any empty string pointer", &p)
		}(),
		want: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.StringReflect("reflect string pointer", &p)
		}(),
		want: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.StringReflect("reflect empty string pointer", &p)
		}(),
		want: `{
			"reflect empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint("uint", 42),
		want: `{
			"uint":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any uint", 42),
		want: `{
			"any uint":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect uint", 42),
		want: `{
			"reflect uint":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.StringUintp("uint pointer", &i)
		}(),
		want: `{
			"uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUintp("nil uint pointer", nil),
		want: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.StringAny("any uint pointer", &i)
		}(),
		want: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.StringReflect("reflect uint pointer", &i)
		}(),
		want: `{
			"reflect uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint16("uint16", 42),
		want: `{
			"uint16":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any uint16", 42),
		want: `{
			"any uint16":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect uint16", 42),
		want: `{
			"reflect uint16":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.StringUint16p("uint16 pointer", &i)
		}(),
		want: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint16p("uint16 pointer", nil),
		want: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.StringAny("any uint16 pointer", &i)
		}(),
		want: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.StringReflect("reflect uint16 pointer", &i)
		}(),
		want: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i *uint16
			return plog.StringReflect("reflect uint16 pointer to nil", i)
		}(),
		want: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint32("uint32", 42),
		want: `{
			"uint32":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any uint32", 42),
		want: `{
			"any uint32":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect uint32", 42),
		want: `{
			"reflect uint32":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.StringUint32p("uint32 pointer", &i)
		}(),
		want: `{
			"uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint32p("nil uint32 pointer", nil),
		want: `{
			"nil uint32 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.StringAny("any uint32 pointer", &i)
		}(),
		want: `{
			"any uint32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.StringReflect("reflect uint32 pointer", &i)
		}(),
		want: `{
			"reflect uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint64("uint64", 42),
		want: `{
			"uint64":42
		}`,
	},

	{
		line:  line(),
		input: plog.StringAny("any uint64", 42),
		want: `{
			"any uint64":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect uint64", 42),
		want: `{
			"reflect uint64":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.StringUint64p("uint64 pointer", &i)
		}(),
		want: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint64p("nil uint64 pointer", nil),
		want: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.StringAny("any uint64 pointer", &i)
		}(),
		want: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.StringReflect("reflect uint64 pointer", &i)
		}(),
		want: `{
			"reflect uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint8("uint8", 42),
		want: `{
			"uint8":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any uint8", 42),
		want: `{
			"any uint8":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect uint8", 42),
		want: `{
			"reflect uint8":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.StringUint8p("uint8 pointer", &i)
		}(),
		want: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUint8p("nil uint8 pointer", nil),
		want: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.StringAny("any uint8 pointer", &i)
		}(),
		want: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.StringReflect("reflect uint8 pointer", &i)
		}(),
		want: `{
			"reflect uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUintptr("uintptr", 42),
		want: `{
			"uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: plog.StringAny("any uintptr", 42),
		want: `{
			"any uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: plog.StringReflect("reflect uintptr", 42),
		want: `{
			"reflect uintptr":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.StringUintptrp("uintptr pointer", &i)
		}(),
		want: `{
			"uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringUintptrp("nil uintptr pointer", nil),
		want: `{
			"nil uintptr pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.StringAny("any uintptr pointer", &i)
		}(),
		want: `{
			"any uintptr pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.StringReflect("reflect uintptr pointer", &i)
		}(),
		want: `{
			"reflect uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringTime("time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.StringTimep("time pointer", &t)
		}(),
		want: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var t *time.Time
			return plog.StringTimep("nil time pointer", t)
		}(),
		want: `{
			"nil time pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			return plog.StringFunc("function", func() plog.KV {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return plog.Time(t)
			})
		}(),
		want: `{
			"function":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.StringAny("any time pointer", &t)
		}(),
		want: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.StringReflect("reflect time pointer", &t)
		}(),
		want: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.StringDuration("duration", 42*time.Nanosecond),
		want: `{
			"duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any duration", 42*time.Nanosecond),
		want: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect duration", 42*time.Nanosecond),
		want: `{
			"reflect duration":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.StringDurationp("duration pointer", &d)
		}(),
		want: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.StringDurationp("nil duration pointer", nil),
		want: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.StringAny("any duration pointer", &d)
		}(),
		want: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.StringReflect("reflect duration pointer", &d)
		}(),
		want: `{
			"reflect duration pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any struct", Struct{Name: "John Doe", Age: 42}),
		want: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return plog.StringAny("any struct pointer", &s)
		}(),
		want: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("struct reflect", Struct{Name: "John Doe", Age: 42}),
		want: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return plog.StringReflect("struct reflect pointer", &s)
		}(),
		want: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: plog.StringRaw("raw json", []byte(`{"foo":"bar"}`)),
		want: `{
			"raw json":{"foo":"bar"}
		}`,
	},
	{
		line:  line(),
		input: plog.StringRaw("raw malformed json object", []byte(`xyz{"foo":"bar"}`)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:  line(),
		input: plog.StringRaw("raw malformed json key/value", []byte(`{"foo":"bar""}`)),
		error: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:  line(),
		input: plog.StringRaw("raw json with unescaped null byte", append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:  line(),
		input: plog.StringRaw("raw nil", nil),
		want: `{
			"raw nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any byte array", [3]byte{'f', 'o', 'o'}),
		want: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			a := [3]byte{'f', 'o', 'o'}
			return plog.StringAny("any byte array pointer", &a)
		}(),
		want: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var a *[3]byte
			return plog.StringAny("any byte array pointer to nil", a)
		}(),
		want: `{
			"any byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect byte array", [3]byte{'f', 'o', 'o'}),
		want: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			a := [3]byte{'f', 'o', 'o'}
			return plog.StringReflect("reflect byte array pointer", &a)
		}(),
		want: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var a *[3]byte
			return plog.StringReflect("reflect byte array pointer to nil", a)
		}(),
		want: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringAny("any untyped nil", nil),
		want: `{
			"any untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.StringReflect("reflect untyped nil", nil),
		want: `{
			"reflect untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextBool(plog.String("bool true"), true),
		want: `{
			"bool true":true
		}`,
	},
	{
		line:  line(),
		input: plog.TextBool(plog.String("bool false"), false),
		want: `{
			"bool false":false
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any bool false"), false),
		want: `{
			"any bool false":false
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("reflect bool false"), false),
		want: `{
			"reflect bool false":false
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.TextBoolp(plog.String("bool pointer to true"), &b)
		}(),
		want: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := false
			return plog.TextBoolp(plog.String("bool pointer to false"), &b)
		}(),
		want: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:  line(),
		input: plog.TextBoolp(plog.String("bool nil pointer"), nil),
		want: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.TextAny(plog.String("any bool pointer to true"), &b)
		}(),
		want: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			b2 := &b
			return plog.TextAny(plog.String("any twice/nested pointer to bool true"), &b2)
		}(),
		want: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			return plog.TextReflect(plog.String("reflect bool pointer to true"), &b)
		}(),
		want: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			b := true
			b2 := &b
			return plog.TextReflect(plog.String("reflect bool twice/nested pointer to true"), &b2)
		}(),
		want: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var b *bool
			return plog.TextReflect(plog.String("reflect bool pointer to nil"), b)
		}(),
		want: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes"), []byte("Hello, Wörld!")...),
		want: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes with quote"), []byte(`Hello, "World"!`)...),
		want: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes quote"), []byte(`"Hello, World!"`)...),
		want: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes nested quote"), []byte(`"Hello, "World"!"`)...),
		want: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes json"), []byte(`{"foo":"bar"}`)...),
		want: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("bytes json quote"), []byte(`"{"foo":"bar"}"`)...),
		want: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytes(plog.String("empty bytes"), []byte{}...),
		want: `{
			"empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var p []byte
			return plog.TextBytes(plog.String("nil bytes"), p...)
		}(),
		want: `{
			"nil bytes":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any bytes"), []byte("Hello, Wörld!")),
		want: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any empty bytes"), []byte{}),
		want: `{
			"any empty bytes":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect bytes"), []byte("Hello, Wörld!")),
		want: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect empty bytes"), []byte{}),
		want: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.TextBytesp(plog.String("bytes pointer"), &p)
		}(),
		want: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.TextBytesp(plog.String("empty bytes pointer"), &p)
		}(),
		want: `{
			"empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextBytesp(plog.String("nil bytes pointer"), nil),
		want: `{
			"nil bytes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.TextAny(plog.String("any bytes pointer"), &p)
		}(),
		want: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.TextAny(plog.String("any empty bytes pointer"), &p)
		}(),
		want: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte("Hello, Wörld!")
			return plog.TextReflect(plog.String("reflect bytes pointer"), &p)
		}(),
		want: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []byte{}
			return plog.TextReflect(plog.String("reflect empty bytes pointer"), &p)
		}(),
		want: `{
			"reflect empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextComplex128(plog.String("complex128"), complex(1, 23)),
		want: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any complex128"), complex(1, 23)),
		want: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect complex128"), complex(1, 23)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.TextComplex128p(plog.String("complex128 pointer"), &c)
		}(),
		want: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextComplex128p(plog.String("nil complex128 pointer"), nil),
		want: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.TextAny(plog.String("any complex128 pointer"), &c)
		}(),
		want: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex128 = complex(1, 23)
			return plog.TextReflect(plog.String("reflect complex128 pointer"), &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: plog.TextComplex64(plog.String("complex64"), complex(3, 21)),
		want: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any complex64"), complex(3, 21)),
		want: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect complex64"), complex(3, 21)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: plog.TextError(plog.String("error"), errors.New("something went wrong")),
		want: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: plog.TextError(plog.String("nil error"), nil),
		want: `{
			"nil error":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any error"), errors.New("something went wrong")),
		want: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect error"), errors.New("something went wrong")),
		want: `{
			"reflect error":{}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.TextComplex64p(plog.String("complex64 pointer"), &c)
		}(),
		want: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: plog.TextComplex64p(plog.String("nil complex64 pointer"), nil),
		want: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.TextAny(plog.String("any complex64 pointer"), &c)
		}(),
		want: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var c complex64 = complex(1, 23)
			return plog.TextReflect(plog.String("reflect complex64 pointer"), &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
	{
		line:  line(),
		input: plog.TextFloat32(plog.String("float32"), 4.2),
		want: `{
			"float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat32(plog.String("high precision float32"), 0.123456789),
		want: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat32(plog.String("zero float32"), 0),
		want: `{
			"zero float32":0
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any float32"), 4.2),
		want: `{
			"any float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any zero float32"), 0),
		want: `{
			"any zero float32":0
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect float32"), 4.2),
		want: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect zero float32"), 0),
		want: `{
			"reflect zero float32":0
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.TextFloat32p(plog.String("float32 pointer"), &f)
		}(),
		want: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 0.123456789
			return plog.TextFloat32p(plog.String("high precision float32 pointer"), &f)
		}(),
		want: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat32p(plog.String("float32 nil pointer"), nil),
		want: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.TextAny(plog.String("any float32 pointer"), &f)
		}(),
		want: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float32 = 4.2
			return plog.TextReflect(plog.String("reflect float32 pointer"), &f)
		}(),
		want: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f *float32
			return plog.TextReflect(plog.String("reflect float32 pointer to nil"), f)
		}(),
		want: `{
			"reflect float32 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat64(plog.String("float64"), 4.2),
		want: `{
			"float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat64(plog.String("high precision float64"), 0.123456789),
		want: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat64(plog.String("zero float64"), 0),
		want: `{
			"zero float64":0
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any float64"), 4.2),
		want: `{
			"any float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any zero float64"), 0),
		want: `{
			"any zero float64":0
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect float64"), 4.2),
		want: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect zero float64"), 0),
		want: `{
			"reflect zero float64":0
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.TextFloat64p(plog.String("float64 pointer"), &f)
		}(),
		want: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 0.123456789
			return plog.TextFloat64p(plog.String("high precision float64 pointer"), &f)
		}(),
		want: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: plog.TextFloat64p(plog.String("float64 nil pointer"), nil),
		want: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.TextAny(plog.String("any float64 pointer"), &f)
		}(),
		want: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f float64 = 4.2
			return plog.TextReflect(plog.String("reflect float64 pointer"), &f)
		}(),
		want: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var f *float64
			return plog.TextReflect(plog.String("reflect float64 pointer to nil"), f)
		}(),
		want: `{
			"reflect float64 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextInt(plog.String("int"), 42),
		want: `{
			"int":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any int"), 42),
		want: `{
			"any int":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect int"), 42),
		want: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.TextIntp(plog.String("int pointer"), &i)
		}(),
		want: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.TextAny(plog.String("any int pointer"), &i)
		}(),
		want: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int = 42
			return plog.TextReflect(plog.String("reflect int pointer"), &i)
		}(),
		want: `{
			"reflect int pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextInt16(plog.String("int16"), 42),
		want: `{
			"int16":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any int16"), 42),
		want: `{
			"any int16":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect int16"), 42),
		want: `{
			"reflect int16":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.TextInt16p(plog.String("int16 pointer"), &i)
		}(),
		want: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.TextAny(plog.String("any int16 pointer"), &i)
		}(),
		want: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int16 = 42
			return plog.TextReflect(plog.String("reflect int16 pointer"), &i)
		}(),
		want: `{
			"reflect int16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextInt32(plog.String("int32"), 42),
		want: `{
			"int32":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any int32"), 42),
		want: `{
			"any int32":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect int32"), 42),
		want: `{
			"reflect int32":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.TextInt32p(plog.String("int32 pointer"), &i)
		}(),
		want: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.TextAny(plog.String("any int32 pointer"), &i)
		}(),
		want: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int32 = 42
			return plog.TextReflect(plog.String("reflect int32 pointer"), &i)
		}(),
		want: `{
			"reflect int32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextInt64(plog.String("int64"), 42),
		want: `{
			"int64":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any int64"), 42),
		want: `{
			"any int64":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect int64"), 42),
		want: `{
			"reflect int64":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.TextInt64p(plog.String("int64 pointer"), &i)
		}(),
		want: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.TextAny(plog.String("any int64 pointer"), &i)
		}(),
		want: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int64 = 42
			return plog.TextReflect(plog.String("reflect int64 pointer"), &i)
		}(),
		want: `{
			"reflect int64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextInt8(plog.String("int8"), 42),
		want: `{
			"int8":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any int8"), 42),
		want: `{
			"any int8":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect int8"), 42),
		want: `{
			"reflect int8":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.TextInt8p(plog.String("int8 pointer"), &i)
		}(),
		want: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.TextAny(plog.String("any int8 pointer"), &i)
		}(),
		want: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i int8 = 42
			return plog.TextReflect(plog.String("reflect int8 pointer"), &i)
		}(),
		want: `{
			"reflect int8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextRunes(plog.String("runes"), []rune("Hello, Wörld!")...),
		want: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextRunes(plog.String("empty runes"), []rune{}...),
		want: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var p []rune
			return plog.TextRunes(plog.String("nil runes"), p...)
		}(),
		want: `{
			"nil runes":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextRunes(plog.String("rune slice with zero rune"), []rune{rune(0)}...),
		want: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any runes"), []rune("Hello, Wörld!")),
		want: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any empty runes"), []rune{}),
		want: `{
			"any empty runes":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any rune slice with zero rune"), []rune{rune(0)}),
		want: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect runes"), []rune("Hello, Wörld!")),
		want: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect empty runes"), []rune{}),
		want: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect rune slice with zero rune"), []rune{rune(0)}),
		want: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.TextRunesp(plog.String("runes pointer"), &p)
		}(),
		want: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.TextRunesp(plog.String("empty runes pointer"), &p)
		}(),
		want: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextRunesp(plog.String("nil runes pointer"), nil),
		want: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.TextAny(plog.String("any runes pointer"), &p)
		}(),
		want: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.TextAny(plog.String("any empty runes pointer"), &p)
		}(),
		want: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune("Hello, Wörld!")
			return plog.TextReflect(plog.String("reflect runes pointer"), &p)
		}(),
		want: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := []rune{}
			return plog.TextReflect(plog.String("reflect empty runes pointer"), &p)
		}(),
		want: `{
			"reflect empty runes pointer":[]
		}`,
	},
	{
		line:  line(),
		input: plog.TextText(plog.String("string"), plog.String("Hello, Wörld!")),
		want: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextText(plog.String("empty string"), plog.String("")),
		want: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextText(plog.String("string with zero byte"), plog.String((string(byte(0))))),
		want: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.TextString(plog.String("string"), "Hello, Wörld!"),
		want: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextString(plog.String("empty string"), ""),
		want: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextString(plog.String("string with zero byte"), string(byte(0))),
		want: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any string"), "Hello, Wörld!"),
		want: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any empty string"), ""),
		want: `{
			"any empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any string with zero byte"), string(byte(0))),
		want: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect string"), "Hello, Wörld!"),
		want: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect empty string"), ""),
		want: `{
			"reflect empty string":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect string with zero byte"), string(byte(0))),
		want: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.TextStringp(plog.String("string pointer"), &p)
		}(),
		want: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.TextStringp(plog.String("empty string pointer"), &p)
		}(),
		want: `{
			"empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextStringp(plog.String("nil string pointer"), nil),
		want: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.TextAny(plog.String("any string pointer"), &p)
		}(),
		want: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.TextAny(plog.String("any empty string pointer"), &p)
		}(),
		want: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := "Hello, Wörld!"
			return plog.TextReflect(plog.String("reflect string pointer"), &p)
		}(),
		want: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			p := ""
			return plog.TextReflect(plog.String("reflect empty string pointer"), &p)
		}(),
		want: `{
			"reflect empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint(plog.String("uint"), 42),
		want: `{
			"uint":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any uint"), 42),
		want: `{
			"any uint":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uint"), 42),
		want: `{
			"reflect uint":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.TextUintp(plog.String("uint pointer"), &i)
		}(),
		want: `{
			"uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUintp(plog.String("nil uint pointer"), nil),
		want: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.TextAny(plog.String("any uint pointer"), &i)
		}(),
		want: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint = 42
			return plog.TextReflect(plog.String("reflect uint pointer"), &i)
		}(),
		want: `{
			"reflect uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint16(plog.String("uint16"), 42),
		want: `{
			"uint16":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any uint16"), 42),
		want: `{
			"any uint16":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uint16"), 42),
		want: `{
			"reflect uint16":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.TextUint16p(plog.String("uint16 pointer"), &i)
		}(),
		want: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint16p(plog.String("uint16 pointer"), nil),
		want: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.TextAny(plog.String("any uint16 pointer"), &i)
		}(),
		want: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint16 = 42
			return plog.TextReflect(plog.String("reflect uint16 pointer"), &i)
		}(),
		want: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i *uint16
			return plog.TextReflect(plog.String("reflect uint16 pointer to nil"), i)
		}(),
		want: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint32(plog.String("uint32"), 42),
		want: `{
			"uint32":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any uint32"), 42),
		want: `{
			"any uint32":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uint32"), 42),
		want: `{
			"reflect uint32":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.TextUint32p(plog.String("uint32 pointer"), &i)
		}(),
		want: `{
			"uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint32p(plog.String("nil uint32 pointer"), nil),
		want: `{
			"nil uint32 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.TextAny(plog.String("any uint32 pointer"), &i)
		}(),
		want: `{
			"any uint32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint32 = 42
			return plog.TextReflect(plog.String("reflect uint32 pointer"), &i)
		}(),
		want: `{
			"reflect uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint64(plog.String("uint64"), 42),
		want: `{
			"uint64":42
		}`,
	},

	{
		line:  line(),
		input: plog.TextAny(plog.String("any uint64"), 42),
		want: `{
			"any uint64":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uint64"), 42),
		want: `{
			"reflect uint64":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.TextUint64p(plog.String("uint64 pointer"), &i)
		}(),
		want: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint64p(plog.String("nil uint64 pointer"), nil),
		want: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.TextAny(plog.String("any uint64 pointer"), &i)
		}(),
		want: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint64 = 42
			return plog.TextReflect(plog.String("reflect uint64 pointer"), &i)
		}(),
		want: `{
			"reflect uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint8(plog.String("uint8"), 42),
		want: `{
			"uint8":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any uint8"), 42),
		want: `{
			"any uint8":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uint8"), 42),
		want: `{
			"reflect uint8":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.TextUint8p(plog.String("uint8 pointer"), &i)
		}(),
		want: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUint8p(plog.String("nil uint8 pointer"), nil),
		want: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.TextAny(plog.String("any uint8 pointer"), &i)
		}(),
		want: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uint8 = 42
			return plog.TextReflect(plog.String("reflect uint8 pointer"), &i)
		}(),
		want: `{
			"reflect uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUintptr(plog.String("uintptr"), 42),
		want: `{
			"uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: plog.TextAny(plog.String("any uintptr"), 42),
		want: `{
			"any uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect uintptr"), 42),
		want: `{
			"reflect uintptr":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.TextUintptrp(plog.String("uintptr pointer"), &i)
		}(),
		want: `{
			"uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextUintptrp(plog.String("nil uintptr pointer"), nil),
		want: `{
			"nil uintptr pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.TextAny(plog.String("any uintptr pointer"), &i)
		}(),
		want: `{
			"any uintptr pointer":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var i uintptr = 42
			return plog.TextReflect(plog.String("reflect uintptr pointer"), &i)
		}(),
		want: `{
			"reflect uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextTime(plog.String("time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		want: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.TextTimep(plog.String("time pointer"), &t)
		}(),
		want: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var t *time.Time
			return plog.TextTimep(plog.String("nil time pointer"), t)
		}(),
		want: `{
			"nil time pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			return plog.TextFunc(plog.String("function"), func() json.Marshaler {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return plog.Time(t)
			})
		}(),
		want: `{
			"function":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.TextAny(plog.String("any time pointer"), &t)
		}(),
		want: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return plog.TextReflect(plog.String("reflect time pointer"), &t)
		}(),
		want: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: plog.TextDuration(plog.String("duration"), 42*time.Nanosecond),
		want: `{
			"duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any duration"), 42*time.Nanosecond),
		want: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect duration"), 42*time.Nanosecond),
		want: `{
			"reflect duration":42
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.TextDurationp(plog.String("duration pointer"), &d)
		}(),
		want: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:  line(),
		input: plog.TextDurationp(plog.String("nil duration pointer"), nil),
		want: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.TextAny(plog.String("any duration pointer"), &d)
		}(),
		want: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			d := 42 * time.Nanosecond
			return plog.TextReflect(plog.String("reflect duration pointer"), &d)
		}(),
		want: `{
			"reflect duration pointer":42
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any struct"), Struct{Name: "John Doe", Age: 42}),
		want: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return plog.TextAny(plog.String("any struct pointer"), &s)
		}(),
		want: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("struct reflect"), Struct{Name: "John Doe", Age: 42}),
		want: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return plog.TextReflect(plog.String("struct reflect pointer"), &s)
		}(),
		want: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: plog.TextRaw(plog.String("raw json"), []byte(`{"foo":"bar"}`)),
		want: `{
			"raw json":{"foo":"bar"}
		}`,
	},
	{
		line:  line(),
		input: plog.TextRaw(plog.String("raw malformed json object"), []byte(`xyz{"foo":"bar"}`)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:  line(),
		input: plog.TextRaw(plog.String("raw malformed json key/value"), []byte(`{"foo":"bar""}`)),
		error: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:  line(),
		input: plog.TextRaw(plog.String("raw json with unescaped null byte"), append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:  line(),
		input: plog.TextRaw(plog.String("raw nil"), nil),
		want: `{
			"raw nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any byte array"), [3]byte{'f', 'o', 'o'}),
		want: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			a := [3]byte{'f', 'o', 'o'}
			return plog.TextAny(plog.String("any byte array pointer"), &a)
		}(),
		want: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var a *[3]byte
			return plog.TextAny(plog.String("any byte array pointer to nil"), a)
		}(),
		want: `{
			"any byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect byte array"), [3]byte{'f', 'o', 'o'}),
		want: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			a := [3]byte{'f', 'o', 'o'}
			return plog.TextReflect(plog.String("reflect byte array pointer"), &a)
		}(),
		want: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() plog.KV {
			var a *[3]byte
			return plog.TextReflect(plog.String("reflect byte array pointer to nil"), a)
		}(),
		want: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextAny(plog.String("any untyped nil"), nil),
		want: `{
			"any untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: plog.TextReflect(plog.String("reflect untyped nil"), nil),
		want: `{
			"reflect untyped nil":null
		}`,
	},
}

func TestKV(t *testing.T) {
	for _, tt := range KVTests {
		tt := tt
		t.Run(tt.line+"/"+fmt.Sprint(tt.input), func(t *testing.T) {
			t.Parallel()

			txt, err := tt.input.MarshalText()
			if err != nil {
				t.Fatalf("encoding marshal text error: %s", err)
			}

			m := map[string]json.Marshaler{string(txt): tt.input}

			jsn, err := json.Marshal(m)

			if fmt.Sprint(err) != fmt.Sprint(tt.error) {
				t.Fatalf("unwant marshal error, want: %s, recieved: %s %s", tt.error, err, tt.line)
			}

			if err == nil {
				ja := jsonassert.New(testprinter{t: t, link: tt.line})
				ja.Assertf(string(jsn), tt.want)
			}
		})
	}
}
