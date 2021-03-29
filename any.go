// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log0

import (
	"encoding"
	"encoding/json"
	"time"
)

// Any returns stringer/JSON marshaler interface implementation for the any type.

func Any(v interface{}) anyV { return anyV{V: v} }

type anyV struct{ V interface{} }

func (v anyV) String() string {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).String()
	case *bool:
		return Boolp(x).String()
	case []byte:
		return Bytes(x).String()
	case *[]byte:
		return Bytesp(x).String()
	case complex128:
		return Complex128(x).String()
	case *complex128:
		return Complex128p(x).String()
	case complex64:
		return Complex64(x).String()
	case *complex64:
		return Complex64p(x).String()
	case error:
		return Error(x).String()
	case float32:
		return Float32(x).String()
	case *float32:
		return Float32p(x).String()
	case float64:
		return Float64(x).String()
	case *float64:
		return Float64p(x).String()
	case int:
		return Int(x).String()
	case *int:
		return Intp(x).String()
	case int16:
		return Int16(x).String()
	case *int16:
		return Int16p(x).String()
	case int32:
		return Int32(x).String()
	case *int32:
		return Int32p(x).String()
	case int64:
		return Int64(x).String()
	case *int64:
		return Int64p(x).String()
	case int8:
		return Int8(x).String()
	case *int8:
		return Int8p(x).String()
	case []rune:
		return Runes(x).String()
	case *[]rune:
		return Runesp(x).String()
	case string:
		return String(x).String()
	case *string:
		return Stringp(x).String()
	case uint:
		return Uint(x).String()
	case *uint:
		return Uintp(x).String()
	case uint16:
		return Uint16(x).String()
	case *uint16:
		return Uint16p(x).String()
	case uint32:
		return Uint32(x).String()
	case *uint32:
		return Uint32p(x).String()
	case uint64:
		return Uint64(x).String()
	case *uint64:
		return Uint64p(x).String()
	case uint8:
		return Uint8(x).String()
	case *uint8:
		return Uint8p(x).String()
	case uintptr:
		return Uintptr(x).String()
	case *uintptr:
		return Uintptrp(x).String()
	case time.Time:
		return Time(x).String()
	case *time.Time:
		return Timep(x).String()
	case time.Duration:
		return Duration(x).String()
	case *time.Duration:
		return Durationp(x).String()
	case encoding.TextMarshaler:
		return Text(x).String()
	case json.Marshaler:
		b, _ := x.MarshalJSON()
		return string(b)
	default:
		return Reflect(x).String()
	}
}

func (v anyV) MarshalText() ([]byte, error) {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).MarshalText()
	case *bool:
		return Boolp(x).MarshalText()
	case []byte:
		return Bytes(x).MarshalText()
	case *[]byte:
		return Bytesp(x).MarshalText()
	case complex128:
		return Complex128(x).MarshalText()
	case *complex128:
		return Complex128p(x).MarshalText()
	case complex64:
		return Complex64(x).MarshalText()
	case *complex64:
		return Complex64p(x).MarshalText()
	case error:
		return Error(x).MarshalText()
	case float32:
		return Float32(x).MarshalText()
	case *float32:
		return Float32p(x).MarshalText()
	case float64:
		return Float64(x).MarshalText()
	case *float64:
		return Float64p(x).MarshalText()
	case int:
		return Int(x).MarshalText()
	case *int:
		return Intp(x).MarshalText()
	case int16:
		return Int16(x).MarshalText()
	case *int16:
		return Int16p(x).MarshalText()
	case int32:
		return Int32(x).MarshalText()
	case *int32:
		return Int32p(x).MarshalText()
	case int64:
		return Int64(x).MarshalText()
	case *int64:
		return Int64p(x).MarshalText()
	case int8:
		return Int8(x).MarshalText()
	case *int8:
		return Int8p(x).MarshalText()
	case []rune:
		return Runes(x).MarshalText()
	case *[]rune:
		return Runesp(x).MarshalText()
	case string:
		return String(x).MarshalText()
	case *string:
		return Stringp(x).MarshalText()
	case uint:
		return Uint(x).MarshalText()
	case *uint:
		return Uintp(x).MarshalText()
	case uint16:
		return Uint16(x).MarshalText()
	case *uint16:
		return Uint16p(x).MarshalText()
	case uint32:
		return Uint32(x).MarshalText()
	case *uint32:
		return Uint32p(x).MarshalText()
	case uint64:
		return Uint64(x).MarshalText()
	case *uint64:
		return Uint64p(x).MarshalText()
	case uint8:
		return Uint8(x).MarshalText()
	case *uint8:
		return Uint8p(x).MarshalText()
	case uintptr:
		return Uintptr(x).MarshalText()
	case *uintptr:
		return Uintptrp(x).MarshalText()
	case time.Time:
		return Time(x).MarshalText()
	case *time.Time:
		return Timep(x).MarshalText()
	case time.Duration:
		return Duration(x).MarshalText()
	case *time.Duration:
		return Durationp(x).MarshalText()
	case encoding.TextMarshaler:
		return x.MarshalText()
	default:
		return Reflect(x).MarshalText()
	}
}

func (v anyV) MarshalJSON() ([]byte, error) {
	switch x := v.V.(type) {
	case bool:
		return Bool(x).MarshalJSON()
	case *bool:
		return Boolp(x).MarshalJSON()
	case []byte:
		return Bytes(x).MarshalJSON()
	case *[]byte:
		return Bytesp(x).MarshalJSON()
	case complex128:
		return Complex128(x).MarshalJSON()
	case *complex128:
		return Complex128p(x).MarshalJSON()
	case complex64:
		return Complex64(x).MarshalJSON()
	case *complex64:
		return Complex64p(x).MarshalJSON()
	case error:
		return Error(x).MarshalJSON()
	case float32:
		return Float32(x).MarshalJSON()
	case *float32:
		return Float32p(x).MarshalJSON()
	case float64:
		return Float64(x).MarshalJSON()
	case *float64:
		return Float64p(x).MarshalJSON()
	case int:
		return Int(x).MarshalJSON()
	case *int:
		return Intp(x).MarshalJSON()
	case int16:
		return Int16(x).MarshalJSON()
	case *int16:
		return Int16p(x).MarshalJSON()
	case int32:
		return Int32(x).MarshalJSON()
	case *int32:
		return Int32p(x).MarshalJSON()
	case int64:
		return Int64(x).MarshalJSON()
	case *int64:
		return Int64p(x).MarshalJSON()
	case int8:
		return Int8(x).MarshalJSON()
	case *int8:
		return Int8p(x).MarshalJSON()
	case []rune:
		return Runes(x).MarshalJSON()
	case *[]rune:
		return Runesp(x).MarshalJSON()
	case string:
		return String(x).MarshalJSON()
	case *string:
		return Stringp(x).MarshalJSON()
	case uint:
		return Uint(x).MarshalJSON()
	case *uint:
		return Uintp(x).MarshalJSON()
	case uint16:
		return Uint16(x).MarshalJSON()
	case *uint16:
		return Uint16p(x).MarshalJSON()
	case uint32:
		return Uint32(x).MarshalJSON()
	case *uint32:
		return Uint32p(x).MarshalJSON()
	case uint64:
		return Uint64(x).MarshalJSON()
	case *uint64:
		return Uint64p(x).MarshalJSON()
	case uint8:
		return Uint8(x).MarshalJSON()
	case *uint8:
		return Uint8p(x).MarshalJSON()
	case uintptr:
		return Uintptr(x).MarshalJSON()
	case *uintptr:
		return Uintptrp(x).MarshalJSON()
	case time.Time:
		return Time(x).MarshalJSON()
	case *time.Time:
		return Timep(x).MarshalJSON()
	case time.Duration:
		return Duration(x).MarshalJSON()
	case *time.Duration:
		return Durationp(x).MarshalJSON()
	case encoding.TextMarshaler:
		return Text(x).MarshalJSON()
	case json.Marshaler:
		return x.MarshalJSON()
	default:
		return Reflect(x).MarshalJSON()
	}
}
