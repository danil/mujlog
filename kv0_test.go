// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/danil/equal4"
	"github.com/danil/log0"
	"github.com/kinbiko/jsonassert"
)

var KVTestCases = []struct {
	line      int
	input     log0.KV
	expected  string
	error     error
	benchmark bool
}{
	{
		line:  line(),
		input: log0.StringBool("bool true", true),
		expected: `{
			"bool true":true
		}`,
	},
	{
		line:  line(),
		input: log0.StringBool("bool false", false),
		expected: `{
			"bool false":false
		}`,
	},
	{
		line:  line(),
		input: log0.StringBools("bools true false", true, false),
		expected: `{
			"bools true false":[true,false]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBools("without bools"),
		expected: `{
			"without bools":[]
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any bool false", false),
		expected: `{
			"any bool false":false
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect bool false", false),
		expected: `{
			"reflect bool false":false
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.StringBoolp("bool pointer to true", &b)
		}(),
		expected: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := false
			return log0.StringBoolp("bool pointer to false", &b)
		}(),
		expected: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:  line(),
		input: log0.StringBoolp("bool nil pointer", nil),
		expected: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.StringAny("any bool pointer to true", &b)
		}(),
		expected: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			b2 := &b
			return log0.StringAny("any twice/nested pointer to bool true", &b2)
		}(),
		expected: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.StringReflect("reflect bool pointer to true", &b)
		}(),
		expected: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			b2 := &b
			return log0.StringReflect("reflect bool twice/nested pointer to true", &b2)
		}(),
		expected: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var b *bool
			return log0.StringReflect("reflect bool pointer to nil", b)
		}(),
		expected: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes", []byte("Hello, Wörld!")...),
		expected: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes with quote", []byte(`Hello, "World"!`)...),
		expected: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes quote", []byte(`"Hello, World!"`)...),
		expected: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes nested quote", []byte(`"Hello, "World"!"`)...),
		expected: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes json", []byte(`{"foo":"bar"}`)...),
		expected: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("bytes json quote", []byte(`"{"foo":"bar"}"`)...),
		expected: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytes("empty bytes", []byte{}...),
		expected: `{
			"empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var p []byte
			return log0.StringBytes("nil bytes", p...)
		}(),
		expected: `{
			"nil bytes":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of byte slices", []byte("Hello, Wörld!"), []byte("Hello, World!")),
		expected: `{
			"slice of byte slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of byte slices with quote", []byte(`Hello, "Wörld"!`), []byte(`Hello, "World"!`)),
		expected: `{
			"slice of byte slices with quote":["Hello, \"Wörld\"!","Hello, \"World\"!"]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("quoted slice of byte slices", []byte(`"Hello, Wörld!"`), []byte(`"Hello, World!"`)),
		expected: `{
			"quoted slice of byte slices":["\"Hello, Wörld!\"","\"Hello, World!\""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of byte slices with nested quote", []byte(`"Hello, "Wörld"!"`), []byte(`"Hello, "World"!"`)),
		expected: `{
			"slice of byte slices with nested quote":["\"Hello, \"Wörld\"!\"","\"Hello, \"World\"!\""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of byte slices with json", []byte(`{"foo":"bar"}`), []byte(`{"baz":"xyz"}`)),
		expected: `{
			"slice of byte slices with json":["{\"foo\":\"bar\"}","{\"baz\":\"xyz\"}"]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of byte slices with quoted json", []byte(`"{"foo":"bar"}"`), []byte(`"{"baz":"xyz"}"`)),
		expected: `{
			"slice of byte slices with quoted json":["\"{\"foo\":\"bar\"}\"","\"{\"baz\":\"xyz\"}\""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of empty byte slices", []byte{}, []byte{}),
		expected: `{
			"slice of empty byte slices":["",""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("slice of nil byte slices", nil, nil),
		expected: `{
			"slice of nil byte slices":[null,null]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytess("empty slice of nil slices"),
		expected: `{
			"empty slice of nil slices":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any bytes", []byte("Hello, Wörld!")),
		expected: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any empty bytes", []byte{}),
		expected: `{
			"any empty bytes":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect bytes", []byte("Hello, Wörld!")),
		expected: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect empty bytes", []byte{}),
		expected: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.StringBytesp("bytes pointer", &p)
		}(),
		expected: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.StringBytesp("empty bytes pointer", &p)
		}(),
		expected: `{
			"empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytesp("nil bytes pointer", nil),
		expected: `{
			"nil bytes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p, p2 := []byte("Hello, Wörld!"), []byte("Hello, World!")
			return log0.StringBytessp("slice of byte pointer slices", &p, &p2)
		}(),
		expected: `{
			"slice of byte pointer slices":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p, p2 := []byte{}, []byte{}
			return log0.StringBytessp("slice of empty byte pointer slices", &p, &p2)
		}(),
		expected: `{
			"slice of empty byte pointer slices":["",""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytessp("slice of nil byte pointer slices", nil, nil),
		expected: `{
			"slice of nil byte pointer slices":[null,null]
		}`,
	},
	{
		line:  line(),
		input: log0.StringBytessp("empty slice of byte pointer slices"),
		expected: `{
			"empty slice of byte pointer slices":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.StringAny("any bytes pointer", &p)
		}(),
		expected: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.StringAny("any empty bytes pointer", &p)
		}(),
		expected: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.StringReflect("reflect bytes pointer", &p)
		}(),
		expected: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.StringReflect("reflect empty bytes pointer", &p)
		}(),
		expected: `{
			"reflect empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringComplex128("complex128", complex(1, 23)),
		expected: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any complex128", complex(1, 23)),
		expected: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect complex128", complex(1, 23)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.StringComplex128p("complex128 pointer", &c)
		}(),
		expected: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringComplex128p("nil complex128 pointer", nil),
		expected: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.StringAny("any complex128 pointer", &c)
		}(),
		expected: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.StringReflect("reflect complex128 pointer", &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: log0.StringComplex64("complex64", complex(3, 21)),
		expected: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any complex64", complex(3, 21)),
		expected: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect complex64", complex(3, 21)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: log0.StringError("error", errors.New("something went wrong")),
		expected: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: log0.StringError("nil error", nil),
		expected: `{
			"nil error":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringErrors("errors", errors.New("something went wrong"), errors.New("wrong")),
		expected: `{
			"errors":["something went wrong","wrong"]
		}`,
	},
	{
		line:  line(),
		input: log0.StringErrors("nil errors", nil, nil),
		expected: `{
			"nil errors":[null,null]
		}`,
	},
	{
		line:  line(),
		input: log0.StringErrors("without errors"),
		expected: `{
			"without errors":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any error", errors.New("something went wrong")),
		expected: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect error", errors.New("something went wrong")),
		expected: `{
			"reflect error":{}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.StringComplex64p("complex64 pointer", &c)
		}(),
		expected: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.StringComplex64p("nil complex64 pointer", nil),
		expected: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.StringAny("any complex64 pointer", &c)
		}(),
		expected: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.StringReflect("reflect complex64 pointer", &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
	{
		line:  line(),
		input: log0.StringFloat32("float32", 4.2),
		expected: `{
			"float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat32("high precision float32", 0.123456789),
		expected: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat32("zero float32", 0),
		expected: `{
			"zero float32":0
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any float32", 4.2),
		expected: `{
			"any float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any zero float32", 0),
		expected: `{
			"any zero float32":0
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect float32", 4.2),
		expected: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect zero float32", 0),
		expected: `{
			"reflect zero float32":0
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.StringFloat32p("float32 pointer", &f)
		}(),
		expected: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 0.123456789
			return log0.StringFloat32p("high precision float32 pointer", &f)
		}(),
		expected: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat32p("float32 nil pointer", nil),
		expected: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.StringAny("any float32 pointer", &f)
		}(),
		expected: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.StringReflect("reflect float32 pointer", &f)
		}(),
		expected: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f *float32
			return log0.StringReflect("reflect float32 pointer to nil", f)
		}(),
		expected: `{
			"reflect float32 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat64("float64", 4.2),
		expected: `{
			"float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat64("high precision float64", 0.123456789),
		expected: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat64("zero float64", 0),
		expected: `{
			"zero float64":0
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any float64", 4.2),
		expected: `{
			"any float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any zero float64", 0),
		expected: `{
			"any zero float64":0
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect float64", 4.2),
		expected: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect zero float64", 0),
		expected: `{
			"reflect zero float64":0
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.StringFloat64p("float64 pointer", &f)
		}(),
		expected: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 0.123456789
			return log0.StringFloat64p("high precision float64 pointer", &f)
		}(),
		expected: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.StringFloat64p("float64 nil pointer", nil),
		expected: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.StringAny("any float64 pointer", &f)
		}(),
		expected: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.StringReflect("reflect float64 pointer", &f)
		}(),
		expected: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f *float64
			return log0.StringReflect("reflect float64 pointer to nil", f)
		}(),
		expected: `{
			"reflect float64 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringInt("int", 42),
		expected: `{
			"int":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any int", 42),
		expected: `{
			"any int":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect int", 42),
		expected: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.StringIntp("int pointer", &i)
		}(),
		expected: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.StringAny("any int pointer", &i)
		}(),
		expected: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.StringReflect("reflect int pointer", &i)
		}(),
		expected: `{
			"reflect int pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringInt16("int16", 42),
		expected: `{
			"int16":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any int16", 42),
		expected: `{
			"any int16":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect int16", 42),
		expected: `{
			"reflect int16":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.StringInt16p("int16 pointer", &i)
		}(),
		expected: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.StringAny("any int16 pointer", &i)
		}(),
		expected: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.StringReflect("reflect int16 pointer", &i)
		}(),
		expected: `{
			"reflect int16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringInt32("int32", 42),
		expected: `{
			"int32":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any int32", 42),
		expected: `{
			"any int32":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect int32", 42),
		expected: `{
			"reflect int32":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.StringInt32p("int32 pointer", &i)
		}(),
		expected: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.StringAny("any int32 pointer", &i)
		}(),
		expected: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.StringReflect("reflect int32 pointer", &i)
		}(),
		expected: `{
			"reflect int32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringInt64("int64", 42),
		expected: `{
			"int64":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any int64", 42),
		expected: `{
			"any int64":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect int64", 42),
		expected: `{
			"reflect int64":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.StringInt64p("int64 pointer", &i)
		}(),
		expected: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.StringAny("any int64 pointer", &i)
		}(),
		expected: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.StringReflect("reflect int64 pointer", &i)
		}(),
		expected: `{
			"reflect int64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringInt8("int8", 42),
		expected: `{
			"int8":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any int8", 42),
		expected: `{
			"any int8":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect int8", 42),
		expected: `{
			"reflect int8":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.StringInt8p("int8 pointer", &i)
		}(),
		expected: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.StringAny("any int8 pointer", &i)
		}(),
		expected: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.StringReflect("reflect int8 pointer", &i)
		}(),
		expected: `{
			"reflect int8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringRunes("runes", []rune("Hello, Wörld!")...),
		expected: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringRunes("empty runes", []rune{}...),
		expected: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var p []rune
			return log0.StringRunes("nil runes", p...)
		}(),
		expected: `{
			"nil runes":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringRunes("rune slice with zero rune", []rune{rune(0)}...),
		expected: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any runes", []rune("Hello, Wörld!")),
		expected: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any empty runes", []rune{}),
		expected: `{
			"any empty runes":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any rune slice with zero rune", []rune{rune(0)}),
		expected: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect runes", []rune("Hello, Wörld!")),
		expected: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect empty runes", []rune{}),
		expected: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect rune slice with zero rune", []rune{rune(0)}),
		expected: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.StringRunesp("runes pointer", &p)
		}(),
		expected: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.StringRunesp("empty runes pointer", &p)
		}(),
		expected: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringRunesp("nil runes pointer", nil),
		expected: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.StringAny("any runes pointer", &p)
		}(),
		expected: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.StringAny("any empty runes pointer", &p)
		}(),
		expected: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.StringReflect("reflect runes pointer", &p)
		}(),
		expected: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.StringReflect("reflect empty runes pointer", &p)
		}(),
		expected: `{
			"reflect empty runes pointer":[]
		}`,
	},
	{
		line:  line(),
		input: log0.StringString("string", "Hello, Wörld!"),
		expected: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringString("empty string", ""),
		expected: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringString("string with zero byte", string(byte(0))),
		expected: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.StringStrings("strings", "Hello, Wörld!", "Hello, World!"),
		expected: `{
			"strings":["Hello, Wörld!","Hello, World!"]
		}`,
	},
	{
		line:  line(),
		input: log0.StringStrings("nil strings"),
		expected: `{
			"nil strings":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringStrings("empty strings", "", ""),
		expected: `{
			"empty strings":["",""]
		}`,
	},
	{
		line:  line(),
		input: log0.StringStrings("without strings"),
		expected: `{
			"without strings":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringStrings("strings with zero byte", string(byte(0)), string(byte(0))),
		expected: `{
			"strings with zero byte":["\u0000","\u0000"]
		}`,
	},

	{
		line:  line(),
		input: log0.StringAny("any string", "Hello, Wörld!"),
		expected: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any empty string", ""),
		expected: `{
			"any empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any string with zero byte", string(byte(0))),
		expected: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect string", "Hello, Wörld!"),
		expected: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect empty string", ""),
		expected: `{
			"reflect empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect string with zero byte", string(byte(0))),
		expected: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.StringStringp("string pointer", &p)
		}(),
		expected: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.StringStringp("empty string pointer", &p)
		}(),
		expected: `{
			"empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringStringp("nil string pointer", nil),
		expected: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.StringAny("any string pointer", &p)
		}(),
		expected: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.StringAny("any empty string pointer", &p)
		}(),
		expected: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.StringReflect("reflect string pointer", &p)
		}(),
		expected: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.StringReflect("reflect empty string pointer", &p)
		}(),
		expected: `{
			"reflect empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint("uint", 42),
		expected: `{
			"uint":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any uint", 42),
		expected: `{
			"any uint":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect uint", 42),
		expected: `{
			"reflect uint":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.StringUintp("uint pointer", &i)
		}(),
		expected: `{
			"uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUintp("nil uint pointer", nil),
		expected: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.StringAny("any uint pointer", &i)
		}(),
		expected: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.StringReflect("reflect uint pointer", &i)
		}(),
		expected: `{
			"reflect uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint16("uint16", 42),
		expected: `{
			"uint16":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any uint16", 42),
		expected: `{
			"any uint16":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect uint16", 42),
		expected: `{
			"reflect uint16":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.StringUint16p("uint16 pointer", &i)
		}(),
		expected: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint16p("uint16 pointer", nil),
		expected: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.StringAny("any uint16 pointer", &i)
		}(),
		expected: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.StringReflect("reflect uint16 pointer", &i)
		}(),
		expected: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i *uint16
			return log0.StringReflect("reflect uint16 pointer to nil", i)
		}(),
		expected: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint32("uint32", 42),
		expected: `{
			"uint32":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any uint32", 42),
		expected: `{
			"any uint32":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect uint32", 42),
		expected: `{
			"reflect uint32":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.StringUint32p("uint32 pointer", &i)
		}(),
		expected: `{
			"uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint32p("nil uint32 pointer", nil),
		expected: `{
			"nil uint32 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.StringAny("any uint32 pointer", &i)
		}(),
		expected: `{
			"any uint32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.StringReflect("reflect uint32 pointer", &i)
		}(),
		expected: `{
			"reflect uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint64("uint64", 42),
		expected: `{
			"uint64":42
		}`,
	},

	{
		line:  line(),
		input: log0.StringAny("any uint64", 42),
		expected: `{
			"any uint64":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect uint64", 42),
		expected: `{
			"reflect uint64":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.StringUint64p("uint64 pointer", &i)
		}(),
		expected: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint64p("nil uint64 pointer", nil),
		expected: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.StringAny("any uint64 pointer", &i)
		}(),
		expected: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.StringReflect("reflect uint64 pointer", &i)
		}(),
		expected: `{
			"reflect uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint8("uint8", 42),
		expected: `{
			"uint8":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any uint8", 42),
		expected: `{
			"any uint8":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect uint8", 42),
		expected: `{
			"reflect uint8":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.StringUint8p("uint8 pointer", &i)
		}(),
		expected: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUint8p("nil uint8 pointer", nil),
		expected: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.StringAny("any uint8 pointer", &i)
		}(),
		expected: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.StringReflect("reflect uint8 pointer", &i)
		}(),
		expected: `{
			"reflect uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUintptr("uintptr", 42),
		expected: `{
			"uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: log0.StringAny("any uintptr", 42),
		expected: `{
			"any uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: log0.StringReflect("reflect uintptr", 42),
		expected: `{
			"reflect uintptr":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.StringUintptrp("uintptr pointer", &i)
		}(),
		expected: `{
			"uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringUintptrp("nil uintptr pointer", nil),
		expected: `{
			"nil uintptr pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.StringAny("any uintptr pointer", &i)
		}(),
		expected: `{
			"any uintptr pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.StringReflect("reflect uintptr pointer", &i)
		}(),
		expected: `{
			"reflect uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringTime("time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect time", time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.StringTimep("time pointer", &t)
		}(),
		expected: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var t *time.Time
			return log0.StringTimep("nil time pointer", t)
		}(),
		expected: `{
			"nil time pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			return log0.StringFunc("function", func() log0.KV {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return log0.Time(t)
			})
		}(),
		expected: `{
			"function":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.StringAny("any time pointer", &t)
		}(),
		expected: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.StringReflect("reflect time pointer", &t)
		}(),
		expected: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.StringDuration("duration", 42*time.Nanosecond),
		expected: `{
			"duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any duration", 42*time.Nanosecond),
		expected: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect duration", 42*time.Nanosecond),
		expected: `{
			"reflect duration":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.StringDurationp("duration pointer", &d)
		}(),
		expected: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.StringDurationp("nil duration pointer", nil),
		expected: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.StringAny("any duration pointer", &d)
		}(),
		expected: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.StringReflect("reflect duration pointer", &d)
		}(),
		expected: `{
			"reflect duration pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any struct", Struct{Name: "John Doe", Age: 42}),
		expected: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return log0.StringAny("any struct pointer", &s)
		}(),
		expected: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("struct reflect", Struct{Name: "John Doe", Age: 42}),
		expected: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return log0.StringReflect("struct reflect pointer", &s)
		}(),
		expected: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: log0.StringRaw("raw json", []byte(`{"foo":"bar"}`)),
		expected: `{
			"raw json":{"foo":"bar"}
		}`,
	},
	{
		line:  line(),
		input: log0.StringRaw("raw malformed json object", []byte(`xyz{"foo":"bar"}`)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:  line(),
		input: log0.StringRaw("raw malformed json key/value", []byte(`{"foo":"bar""}`)),
		error: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:  line(),
		input: log0.StringRaw("raw json with unescaped null byte", append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:  line(),
		input: log0.StringRaw("raw nil", nil),
		expected: `{
			"raw nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any byte array", [3]byte{'f', 'o', 'o'}),
		expected: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			a := [3]byte{'f', 'o', 'o'}
			return log0.StringAny("any byte array pointer", &a)
		}(),
		expected: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var a *[3]byte
			return log0.StringAny("any byte array pointer to nil", a)
		}(),
		expected: `{
			"any byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect byte array", [3]byte{'f', 'o', 'o'}),
		expected: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			a := [3]byte{'f', 'o', 'o'}
			return log0.StringReflect("reflect byte array pointer", &a)
		}(),
		expected: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var a *[3]byte
			return log0.StringReflect("reflect byte array pointer to nil", a)
		}(),
		expected: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringAny("any untyped nil", nil),
		expected: `{
			"any untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.StringReflect("reflect untyped nil", nil),
		expected: `{
			"reflect untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextBool(log0.String("bool true"), true),
		expected: `{
			"bool true":true
		}`,
	},
	{
		line:  line(),
		input: log0.TextBool(log0.String("bool false"), false),
		expected: `{
			"bool false":false
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any bool false"), false),
		expected: `{
			"any bool false":false
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("reflect bool false"), false),
		expected: `{
			"reflect bool false":false
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.TextBoolp(log0.String("bool pointer to true"), &b)
		}(),
		expected: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := false
			return log0.TextBoolp(log0.String("bool pointer to false"), &b)
		}(),
		expected: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:  line(),
		input: log0.TextBoolp(log0.String("bool nil pointer"), nil),
		expected: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.TextAny(log0.String("any bool pointer to true"), &b)
		}(),
		expected: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			b2 := &b
			return log0.TextAny(log0.String("any twice/nested pointer to bool true"), &b2)
		}(),
		expected: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			return log0.TextReflect(log0.String("reflect bool pointer to true"), &b)
		}(),
		expected: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			b := true
			b2 := &b
			return log0.TextReflect(log0.String("reflect bool twice/nested pointer to true"), &b2)
		}(),
		expected: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var b *bool
			return log0.TextReflect(log0.String("reflect bool pointer to nil"), b)
		}(),
		expected: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes"), []byte("Hello, Wörld!")...),
		expected: `{
			"bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes with quote"), []byte(`Hello, "World"!`)...),
		expected: `{
			"bytes with quote":"Hello, \"World\"!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes quote"), []byte(`"Hello, World!"`)...),
		expected: `{
			"bytes quote":"\"Hello, World!\""
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes nested quote"), []byte(`"Hello, "World"!"`)...),
		expected: `{
			"bytes nested quote":"\"Hello, \"World\"!\""
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes json"), []byte(`{"foo":"bar"}`)...),
		expected: `{
			"bytes json":"{\"foo\":\"bar\"}"
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("bytes json quote"), []byte(`"{"foo":"bar"}"`)...),
		expected: `{
			"bytes json quote":"\"{\"foo\":\"bar\"}\""
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytes(log0.String("empty bytes"), []byte{}...),
		expected: `{
			"empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var p []byte
			return log0.TextBytes(log0.String("nil bytes"), p...)
		}(),
		expected: `{
			"nil bytes":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any bytes"), []byte("Hello, Wörld!")),
		expected: `{
			"any bytes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any empty bytes"), []byte{}),
		expected: `{
			"any empty bytes":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect bytes"), []byte("Hello, Wörld!")),
		expected: `{
			"reflect bytes":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect empty bytes"), []byte{}),
		expected: `{
			"reflect empty bytes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.TextBytesp(log0.String("bytes pointer"), &p)
		}(),
		expected: `{
			"bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.TextBytesp(log0.String("empty bytes pointer"), &p)
		}(),
		expected: `{
			"empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextBytesp(log0.String("nil bytes pointer"), nil),
		expected: `{
			"nil bytes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.TextAny(log0.String("any bytes pointer"), &p)
		}(),
		expected: `{
			"any bytes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.TextAny(log0.String("any empty bytes pointer"), &p)
		}(),
		expected: `{
			"any empty bytes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte("Hello, Wörld!")
			return log0.TextReflect(log0.String("reflect bytes pointer"), &p)
		}(),
		expected: `{
			"reflect bytes pointer":"SGVsbG8sIFfDtnJsZCE="
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []byte{}
			return log0.TextReflect(log0.String("reflect empty bytes pointer"), &p)
		}(),
		expected: `{
			"reflect empty bytes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextComplex128(log0.String("complex128"), complex(1, 23)),
		expected: `{
			"complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any complex128"), complex(1, 23)),
		expected: `{
			"any complex128":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect complex128"), complex(1, 23)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.TextComplex128p(log0.String("complex128 pointer"), &c)
		}(),
		expected: `{
			"complex128 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextComplex128p(log0.String("nil complex128 pointer"), nil),
		expected: `{
			"nil complex128 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.TextAny(log0.String("any complex128 pointer"), &c)
		}(),
		expected: `{
			"any complex128 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex128 = complex(1, 23)
			return log0.TextReflect(log0.String("reflect complex128 pointer"), &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: log0.TextComplex64(log0.String("complex64"), complex(3, 21)),
		expected: `{
			"complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any complex64"), complex(3, 21)),
		expected: `{
			"any complex64":"3+21i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect complex64"), complex(3, 21)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex128"),
	},
	{
		line:  line(),
		input: log0.TextError(log0.String("error"), errors.New("something went wrong")),
		expected: `{
			"error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: log0.TextError(log0.String("nil error"), nil),
		expected: `{
			"nil error":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any error"), errors.New("something went wrong")),
		expected: `{
			"any error":"something went wrong"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect error"), errors.New("something went wrong")),
		expected: `{
			"reflect error":{}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.TextComplex64p(log0.String("complex64 pointer"), &c)
		}(),
		expected: `{
			"complex64 pointer":"1+23i"
		}`,
	},
	{
		line:  line(),
		input: log0.TextComplex64p(log0.String("nil complex64 pointer"), nil),
		expected: `{
			"nil complex64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.TextAny(log0.String("any complex64 pointer"), &c)
		}(),
		expected: `{
			"any complex64 pointer":"1+23i"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var c complex64 = complex(1, 23)
			return log0.TextReflect(log0.String("reflect complex64 pointer"), &c)
		}(),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: json: unsupported type: complex64"),
	},
	{
		line:  line(),
		input: log0.TextFloat32(log0.String("float32"), 4.2),
		expected: `{
			"float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat32(log0.String("high precision float32"), 0.123456789),
		expected: `{
			"high precision float32":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat32(log0.String("zero float32"), 0),
		expected: `{
			"zero float32":0
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any float32"), 4.2),
		expected: `{
			"any float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any zero float32"), 0),
		expected: `{
			"any zero float32":0
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect float32"), 4.2),
		expected: `{
			"reflect float32":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect zero float32"), 0),
		expected: `{
			"reflect zero float32":0
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.TextFloat32p(log0.String("float32 pointer"), &f)
		}(),
		expected: `{
			"float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 0.123456789
			return log0.TextFloat32p(log0.String("high precision float32 pointer"), &f)
		}(),
		expected: `{
			"high precision float32 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat32p(log0.String("float32 nil pointer"), nil),
		expected: `{
			"float32 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.TextAny(log0.String("any float32 pointer"), &f)
		}(),
		expected: `{
			"any float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float32 = 4.2
			return log0.TextReflect(log0.String("reflect float32 pointer"), &f)
		}(),
		expected: `{
			"reflect float32 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f *float32
			return log0.TextReflect(log0.String("reflect float32 pointer to nil"), f)
		}(),
		expected: `{
			"reflect float32 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat64(log0.String("float64"), 4.2),
		expected: `{
			"float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat64(log0.String("high precision float64"), 0.123456789),
		expected: `{
			"high precision float64":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat64(log0.String("zero float64"), 0),
		expected: `{
			"zero float64":0
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any float64"), 4.2),
		expected: `{
			"any float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any zero float64"), 0),
		expected: `{
			"any zero float64":0
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect float64"), 4.2),
		expected: `{
			"reflect float64":4.2
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect zero float64"), 0),
		expected: `{
			"reflect zero float64":0
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.TextFloat64p(log0.String("float64 pointer"), &f)
		}(),
		expected: `{
			"float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 0.123456789
			return log0.TextFloat64p(log0.String("high precision float64 pointer"), &f)
		}(),
		expected: `{
			"high precision float64 pointer":0.123456789
		}`,
	},
	{
		line:  line(),
		input: log0.TextFloat64p(log0.String("float64 nil pointer"), nil),
		expected: `{
			"float64 nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.TextAny(log0.String("any float64 pointer"), &f)
		}(),
		expected: `{
			"any float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f float64 = 4.2
			return log0.TextReflect(log0.String("reflect float64 pointer"), &f)
		}(),
		expected: `{
			"reflect float64 pointer":4.2
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var f *float64
			return log0.TextReflect(log0.String("reflect float64 pointer to nil"), f)
		}(),
		expected: `{
			"reflect float64 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextInt(log0.String("int"), 42),
		expected: `{
			"int":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any int"), 42),
		expected: `{
			"any int":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect int"), 42),
		expected: `{
			"reflect int":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.TextIntp(log0.String("int pointer"), &i)
		}(),
		expected: `{
			"int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.TextAny(log0.String("any int pointer"), &i)
		}(),
		expected: `{
			"any int pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int = 42
			return log0.TextReflect(log0.String("reflect int pointer"), &i)
		}(),
		expected: `{
			"reflect int pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextInt16(log0.String("int16"), 42),
		expected: `{
			"int16":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any int16"), 42),
		expected: `{
			"any int16":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect int16"), 42),
		expected: `{
			"reflect int16":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.TextInt16p(log0.String("int16 pointer"), &i)
		}(),
		expected: `{
			"int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.TextAny(log0.String("any int16 pointer"), &i)
		}(),
		expected: `{
			"any int16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int16 = 42
			return log0.TextReflect(log0.String("reflect int16 pointer"), &i)
		}(),
		expected: `{
			"reflect int16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextInt32(log0.String("int32"), 42),
		expected: `{
			"int32":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any int32"), 42),
		expected: `{
			"any int32":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect int32"), 42),
		expected: `{
			"reflect int32":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.TextInt32p(log0.String("int32 pointer"), &i)
		}(),
		expected: `{
			"int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.TextAny(log0.String("any int32 pointer"), &i)
		}(),
		expected: `{
			"any int32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int32 = 42
			return log0.TextReflect(log0.String("reflect int32 pointer"), &i)
		}(),
		expected: `{
			"reflect int32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextInt64(log0.String("int64"), 42),
		expected: `{
			"int64":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any int64"), 42),
		expected: `{
			"any int64":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect int64"), 42),
		expected: `{
			"reflect int64":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.TextInt64p(log0.String("int64 pointer"), &i)
		}(),
		expected: `{
			"int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.TextAny(log0.String("any int64 pointer"), &i)
		}(),
		expected: `{
			"any int64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int64 = 42
			return log0.TextReflect(log0.String("reflect int64 pointer"), &i)
		}(),
		expected: `{
			"reflect int64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextInt8(log0.String("int8"), 42),
		expected: `{
			"int8":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any int8"), 42),
		expected: `{
			"any int8":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect int8"), 42),
		expected: `{
			"reflect int8":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.TextInt8p(log0.String("int8 pointer"), &i)
		}(),
		expected: `{
			"int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.TextAny(log0.String("any int8 pointer"), &i)
		}(),
		expected: `{
			"any int8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i int8 = 42
			return log0.TextReflect(log0.String("reflect int8 pointer"), &i)
		}(),
		expected: `{
			"reflect int8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextRunes(log0.String("runes"), []rune("Hello, Wörld!")...),
		expected: `{
			"runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextRunes(log0.String("empty runes"), []rune{}...),
		expected: `{
			"empty runes":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var p []rune
			return log0.TextRunes(log0.String("nil runes"), p...)
		}(),
		expected: `{
			"nil runes":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextRunes(log0.String("rune slice with zero rune"), []rune{rune(0)}...),
		expected: `{
			"rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any runes"), []rune("Hello, Wörld!")),
		expected: `{
			"any runes":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any empty runes"), []rune{}),
		expected: `{
			"any empty runes":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any rune slice with zero rune"), []rune{rune(0)}),
		expected: `{
			"any rune slice with zero rune":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect runes"), []rune("Hello, Wörld!")),
		expected: `{
			"reflect runes":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect empty runes"), []rune{}),
		expected: `{
			"reflect empty runes":[]
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect rune slice with zero rune"), []rune{rune(0)}),
		expected: `{
			"reflect rune slice with zero rune":[0]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.TextRunesp(log0.String("runes pointer"), &p)
		}(),
		expected: `{
			"runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.TextRunesp(log0.String("empty runes pointer"), &p)
		}(),
		expected: `{
			"empty runes pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextRunesp(log0.String("nil runes pointer"), nil),
		expected: `{
			"nil runes pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.TextAny(log0.String("any runes pointer"), &p)
		}(),
		expected: `{
			"any runes pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.TextAny(log0.String("any empty runes pointer"), &p)
		}(),
		expected: `{
			"any empty runes pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune("Hello, Wörld!")
			return log0.TextReflect(log0.String("reflect runes pointer"), &p)
		}(),
		expected: `{
			"reflect runes pointer":[72,101,108,108,111,44,32,87,246,114,108,100,33]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := []rune{}
			return log0.TextReflect(log0.String("reflect empty runes pointer"), &p)
		}(),
		expected: `{
			"reflect empty runes pointer":[]
		}`,
	},
	{
		line:  line(),
		input: log0.TextText(log0.String("string"), log0.String("Hello, Wörld!")),
		expected: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextText(log0.String("empty string"), log0.String("")),
		expected: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextText(log0.String("string with zero byte"), log0.String((string(byte(0))))),
		expected: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.TextString(log0.String("string"), "Hello, Wörld!"),
		expected: `{
			"string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextString(log0.String("empty string"), ""),
		expected: `{
			"empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextString(log0.String("string with zero byte"), string(byte(0))),
		expected: `{
			"string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any string"), "Hello, Wörld!"),
		expected: `{
			"any string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any empty string"), ""),
		expected: `{
			"any empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any string with zero byte"), string(byte(0))),
		expected: `{
			"any string with zero byte":"\u0000"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect string"), "Hello, Wörld!"),
		expected: `{
			"reflect string":"Hello, Wörld!"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect empty string"), ""),
		expected: `{
			"reflect empty string":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect string with zero byte"), string(byte(0))),
		expected: `{
			"reflect string with zero byte":"\u0000"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.TextStringp(log0.String("string pointer"), &p)
		}(),
		expected: `{
			"string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.TextStringp(log0.String("empty string pointer"), &p)
		}(),
		expected: `{
			"empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextStringp(log0.String("nil string pointer"), nil),
		expected: `{
			"nil string pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.TextAny(log0.String("any string pointer"), &p)
		}(),
		expected: `{
			"any string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.TextAny(log0.String("any empty string pointer"), &p)
		}(),
		expected: `{
			"any empty string pointer":""
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := "Hello, Wörld!"
			return log0.TextReflect(log0.String("reflect string pointer"), &p)
		}(),
		expected: `{
			"reflect string pointer":"Hello, Wörld!"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			p := ""
			return log0.TextReflect(log0.String("reflect empty string pointer"), &p)
		}(),
		expected: `{
			"reflect empty string pointer":""
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint(log0.String("uint"), 42),
		expected: `{
			"uint":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any uint"), 42),
		expected: `{
			"any uint":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uint"), 42),
		expected: `{
			"reflect uint":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.TextUintp(log0.String("uint pointer"), &i)
		}(),
		expected: `{
			"uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUintp(log0.String("nil uint pointer"), nil),
		expected: `{
			"nil uint pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.TextAny(log0.String("any uint pointer"), &i)
		}(),
		expected: `{
			"any uint pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint = 42
			return log0.TextReflect(log0.String("reflect uint pointer"), &i)
		}(),
		expected: `{
			"reflect uint pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint16(log0.String("uint16"), 42),
		expected: `{
			"uint16":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any uint16"), 42),
		expected: `{
			"any uint16":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uint16"), 42),
		expected: `{
			"reflect uint16":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.TextUint16p(log0.String("uint16 pointer"), &i)
		}(),
		expected: `{
			"uint16 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint16p(log0.String("uint16 pointer"), nil),
		expected: `{
			"uint16 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.TextAny(log0.String("any uint16 pointer"), &i)
		}(),
		expected: `{
			"any uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint16 = 42
			return log0.TextReflect(log0.String("reflect uint16 pointer"), &i)
		}(),
		expected: `{
			"reflect uint16 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i *uint16
			return log0.TextReflect(log0.String("reflect uint16 pointer to nil"), i)
		}(),
		expected: `{
			"reflect uint16 pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint32(log0.String("uint32"), 42),
		expected: `{
			"uint32":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any uint32"), 42),
		expected: `{
			"any uint32":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uint32"), 42),
		expected: `{
			"reflect uint32":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.TextUint32p(log0.String("uint32 pointer"), &i)
		}(),
		expected: `{
			"uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint32p(log0.String("nil uint32 pointer"), nil),
		expected: `{
			"nil uint32 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.TextAny(log0.String("any uint32 pointer"), &i)
		}(),
		expected: `{
			"any uint32 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint32 = 42
			return log0.TextReflect(log0.String("reflect uint32 pointer"), &i)
		}(),
		expected: `{
			"reflect uint32 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint64(log0.String("uint64"), 42),
		expected: `{
			"uint64":42
		}`,
	},

	{
		line:  line(),
		input: log0.TextAny(log0.String("any uint64"), 42),
		expected: `{
			"any uint64":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uint64"), 42),
		expected: `{
			"reflect uint64":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.TextUint64p(log0.String("uint64 pointer"), &i)
		}(),
		expected: `{
			"uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint64p(log0.String("nil uint64 pointer"), nil),
		expected: `{
			"nil uint64 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.TextAny(log0.String("any uint64 pointer"), &i)
		}(),
		expected: `{
			"any uint64 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint64 = 42
			return log0.TextReflect(log0.String("reflect uint64 pointer"), &i)
		}(),
		expected: `{
			"reflect uint64 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint8(log0.String("uint8"), 42),
		expected: `{
			"uint8":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any uint8"), 42),
		expected: `{
			"any uint8":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uint8"), 42),
		expected: `{
			"reflect uint8":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.TextUint8p(log0.String("uint8 pointer"), &i)
		}(),
		expected: `{
			"uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUint8p(log0.String("nil uint8 pointer"), nil),
		expected: `{
			"nil uint8 pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.TextAny(log0.String("any uint8 pointer"), &i)
		}(),
		expected: `{
			"any uint8 pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uint8 = 42
			return log0.TextReflect(log0.String("reflect uint8 pointer"), &i)
		}(),
		expected: `{
			"reflect uint8 pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUintptr(log0.String("uintptr"), 42),
		expected: `{
			"uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: log0.TextAny(log0.String("any uintptr"), 42),
		expected: `{
			"any uintptr":42
		}`,
	},
	// FIXME: use var x uintptr = 42
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect uintptr"), 42),
		expected: `{
			"reflect uintptr":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.TextUintptrp(log0.String("uintptr pointer"), &i)
		}(),
		expected: `{
			"uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextUintptrp(log0.String("nil uintptr pointer"), nil),
		expected: `{
			"nil uintptr pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.TextAny(log0.String("any uintptr pointer"), &i)
		}(),
		expected: `{
			"any uintptr pointer":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var i uintptr = 42
			return log0.TextReflect(log0.String("reflect uintptr pointer"), &i)
		}(),
		expected: `{
			"reflect uintptr pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextTime(log0.String("time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"any time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect time"), time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)),
		expected: `{
			"reflect time":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.TextTimep(log0.String("time pointer"), &t)
		}(),
		expected: `{
			"time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var t *time.Time
			return log0.TextTimep(log0.String("nil time pointer"), t)
		}(),
		expected: `{
			"nil time pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			return log0.TextFunc(log0.String("function"), func() json.Marshaler {
				t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
				return log0.Time(t)
			})
		}(),
		expected: `{
			"function":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.TextAny(log0.String("any time pointer"), &t)
		}(),
		expected: `{
			"any time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			t := time.Date(1970, time.January, 1, 0, 0, 0, 42, time.UTC)
			return log0.TextReflect(log0.String("reflect time pointer"), &t)
		}(),
		expected: `{
			"reflect time pointer":"1970-01-01T00:00:00.000000042Z"
		}`,
	},
	{
		line:  line(),
		input: log0.TextDuration(log0.String("duration"), 42*time.Nanosecond),
		expected: `{
			"duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any duration"), 42*time.Nanosecond),
		expected: `{
			"any duration":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect duration"), 42*time.Nanosecond),
		expected: `{
			"reflect duration":42
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.TextDurationp(log0.String("duration pointer"), &d)
		}(),
		expected: `{
			"duration pointer":"42ns"
		}`,
	},
	{
		line:  line(),
		input: log0.TextDurationp(log0.String("nil duration pointer"), nil),
		expected: `{
			"nil duration pointer":null
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.TextAny(log0.String("any duration pointer"), &d)
		}(),
		expected: `{
			"any duration pointer":"42ns"
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			d := 42 * time.Nanosecond
			return log0.TextReflect(log0.String("reflect duration pointer"), &d)
		}(),
		expected: `{
			"reflect duration pointer":42
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any struct"), Struct{Name: "John Doe", Age: 42}),
		expected: `{
			"any struct": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return log0.TextAny(log0.String("any struct pointer"), &s)
		}(),
		expected: `{
			"any struct pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("struct reflect"), Struct{Name: "John Doe", Age: 42}),
		expected: `{
			"struct reflect": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			s := Struct{Name: "John Doe", Age: 42}
			return log0.TextReflect(log0.String("struct reflect pointer"), &s)
		}(),
		expected: `{
			"struct reflect pointer": {
				"Name":"John Doe",
				"Age":42
			}
		}`,
	},
	{
		line:  line(),
		input: log0.TextRaw(log0.String("raw json"), []byte(`{"foo":"bar"}`)),
		expected: `{
			"raw json":{"foo":"bar"}
		}`,
	},
	{
		line:  line(),
		input: log0.TextRaw(log0.String("raw malformed json object"), []byte(`xyz{"foo":"bar"}`)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character 'x' looking for beginning of value"),
	},
	{
		line:  line(),
		input: log0.TextRaw(log0.String("raw malformed json key/value"), []byte(`{"foo":"bar""}`)),
		error: errors.New(`json: error calling MarshalJSON for type json.Marshaler: invalid character '"' after object key:value pair`),
	},
	{
		line:  line(),
		input: log0.TextRaw(log0.String("raw json with unescaped null byte"), append([]byte(`{"foo":"`), append([]byte{0}, []byte(`xyz"}`)...)...)),
		error: errors.New("json: error calling MarshalJSON for type json.Marshaler: invalid character '\\x00' in string literal"),
	},
	{
		line:  line(),
		input: log0.TextRaw(log0.String("raw nil"), nil),
		expected: `{
			"raw nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any byte array"), [3]byte{'f', 'o', 'o'}),
		expected: `{
			"any byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			a := [3]byte{'f', 'o', 'o'}
			return log0.TextAny(log0.String("any byte array pointer"), &a)
		}(),
		expected: `{
			"any byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var a *[3]byte
			return log0.TextAny(log0.String("any byte array pointer to nil"), a)
		}(),
		expected: `{
			"any byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect byte array"), [3]byte{'f', 'o', 'o'}),
		expected: `{
			"reflect byte array":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			a := [3]byte{'f', 'o', 'o'}
			return log0.TextReflect(log0.String("reflect byte array pointer"), &a)
		}(),
		expected: `{
			"reflect byte array pointer":[102,111,111]
		}`,
	},
	{
		line: line(),
		input: func() log0.KV {
			var a *[3]byte
			return log0.TextReflect(log0.String("reflect byte array pointer to nil"), a)
		}(),
		expected: `{
			"reflect byte array pointer to nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextAny(log0.String("any untyped nil"), nil),
		expected: `{
			"any untyped nil":null
		}`,
	},
	{
		line:  line(),
		input: log0.TextReflect(log0.String("reflect untyped nil"), nil),
		expected: `{
			"reflect untyped nil":null
		}`,
	},
}

func TestKV(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range KVTestCases {
		tc := tc
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			txt, err := tc.input.MarshalText()
			if err != nil {
				t.Fatalf("encoding marshal text error: %s", err)
			}

			m := map[string]json.Marshaler{string(txt): tc.input}

			jsn, err := json.Marshal(m)

			if !equal4.ErrorEqual(err, tc.error) {
				t.Fatalf("unexpected marshal error, expected: %s, recieved: %s %s", tc.error, err, linkToExample)
			}

			if err == nil {
				ja := jsonassert.New(testprinter{t: t, link: linkToExample})
				ja.Assertf(string(jsn), tc.expected)
			}
		})
	}
}
