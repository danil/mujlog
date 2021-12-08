// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog

import (
	"encoding"
	"encoding/json"
	"fmt"
	"time"
)

// kvm is a key-value pair implements json/text marshaler.
type kvm struct {
	K encoding.TextMarshaler
	V json.Marshaler
}

func (kv kvm) MarshalText() (text []byte, err error) { return kv.K.MarshalText() }
func (kv kvm) MarshalJSON() ([]byte, error)          { return kv.V.MarshalJSON() }

func StringBool(k string, v bool) kvm {
	return kvm{K: String(k), V: Bool(v)}
}

func StringBools(k string, v ...bool) kvm {
	return kvm{K: String(k), V: Bools(v...)}
}

func StringBoolp(k string, v *bool) kvm {
	return kvm{K: String(k), V: Boolp(v)}
}

func StringBytes(k string, v ...byte) kvm {
	return kvm{K: String(k), V: Bytes(v...)}
}

func StringBytess(k string, v ...[]byte) kvm {
	return kvm{K: String(k), V: Bytess(v...)}
}

func StringBytesp(k string, v *[]byte) kvm {
	return kvm{K: String(k), V: Bytesp(v)}
}

func StringBytessp(k string, v ...*[]byte) kvm {
	return kvm{K: String(k), V: Bytesps(v...)}
}

func StringComplex128(k string, v complex128) kvm {
	return kvm{K: String(k), V: Complex128(v)}
}

func StringComplex128p(k string, v *complex128) kvm {
	return kvm{K: String(k), V: Complex128p(v)}
}

func StringComplex64(k string, v complex64) kvm {
	return kvm{K: String(k), V: Complex64(v)}
}

func StringComplex64p(k string, v *complex64) kvm {
	return kvm{K: String(k), V: Complex64p(v)}
}

func StringError(k string, v error) kvm {
	return kvm{K: String(k), V: Error(v)}
}

func StringErrors(k string, v ...error) kvm {
	return kvm{K: String(k), V: Errors(v...)}
}

func StringFloat32(k string, v float32) kvm {
	return kvm{K: String(k), V: Float32(v)}
}

func StringFloat32p(k string, v *float32) kvm {
	return kvm{K: String(k), V: Float32p(v)}
}

func StringFloat64(k string, v float64) kvm {
	return kvm{K: String(k), V: Float64(v)}
}

func StringFloat64p(k string, v *float64) kvm {
	return kvm{K: String(k), V: Float64p(v)}
}

func StringInt(k string, v int) kvm {
	return kvm{K: String(k), V: Int(v)}
}

func StringIntp(k string, v *int) kvm {
	return kvm{K: String(k), V: Intp(v)}
}

func StringInt16(k string, v int16) kvm {
	return kvm{K: String(k), V: Int16(v)}
}

func StringInt16p(k string, v *int16) kvm {
	return kvm{K: String(k), V: Int16p(v)}
}

func StringInt32(k string, v int32) kvm {
	return kvm{K: String(k), V: Int32(v)}
}

func StringInt32p(k string, v *int32) kvm {
	return kvm{K: String(k), V: Int32p(v)}
}

func StringInt64(k string, v int64) kvm {
	return kvm{K: String(k), V: Int64(v)}
}

func StringInt64p(k string, v *int64) kvm {
	return kvm{K: String(k), V: Int64p(v)}
}

func StringInt8(k string, v int8) kvm {
	return kvm{K: String(k), V: Int8(v)}
}

func StringInt8p(k string, v *int8) kvm {
	return kvm{K: String(k), V: Int8p(v)}
}

func StringRunes(k string, v ...rune) kvm {
	return kvm{K: String(k), V: Runes(v...)}
}

func StringRunesp(k string, v *[]rune) kvm {
	return kvm{K: String(k), V: Runesp(v)}
}

func StringString(k string, v string) kvm {
	return kvm{K: String(k), V: String(v)}
}

func StringStrings(k string, v ...string) kvm {
	return kvm{K: String(k), V: Strings(v...)}
}

func StringStringp(k string, v *string) kvm {
	return kvm{K: String(k), V: Stringp(v)}
}

func StringUint(k string, v uint) kvm {
	return kvm{K: String(k), V: Uint(v)}
}

func StringUintp(k string, v *uint) kvm {
	return kvm{K: String(k), V: Uintp(v)}
}

func StringUint16(k string, v uint16) kvm {
	return kvm{K: String(k), V: Uint16(v)}
}

func StringUint16p(k string, v *uint16) kvm {
	return kvm{K: String(k), V: Uint16p(v)}
}

func StringUint32(k string, v uint32) kvm {
	return kvm{K: String(k), V: Uint32(v)}
}

func StringUint32p(k string, v *uint32) kvm {
	return kvm{K: String(k), V: Uint32p(v)}
}

func StringUint64(k string, v uint64) kvm {
	return kvm{K: String(k), V: Uint64(v)}
}

func StringUint64p(k string, v *uint64) kvm {
	return kvm{K: String(k), V: Uint64p(v)}
}

func StringUint8(k string, v uint8) kvm {
	return kvm{K: String(k), V: Uint8(v)}
}

func StringUint8p(k string, v *uint8) kvm {
	return kvm{K: String(k), V: Uint8p(v)}
}

func StringUintptr(k string, v uintptr) kvm {
	return kvm{K: String(k), V: Uintptr(v)}
}

func StringUintptrp(k string, v *uintptr) kvm {
	return kvm{K: String(k), V: Uintptrp(v)}
}

func StringDuration(k string, v time.Duration) kvm {
	return kvm{K: String(k), V: Duration(v)}
}

func StringDurationp(k string, v *time.Duration) kvm {
	return kvm{K: String(k), V: Durationp(v)}
}

func StringTime(k string, v time.Time) kvm {
	return kvm{K: String(k), V: Time(v)}
}

func StringTimep(k string, v *time.Time) kvm {
	return kvm{K: String(k), V: Timep(v)}
}

func StringFunc(k string, v func() KV) kvm {
	return kvm{K: String(k), V: Func(v)}
}

func StringRaw(k string, v []byte) kvm {
	return kvm{K: String(k), V: Raw(v)}
}

func StringAny(k string, v interface{}) kvm {
	return kvm{K: String(k), V: Any(v)}
}

func StringReflect(k string, v interface{}) kvm {
	return kvm{K: String(k), V: Reflect(v)}
}

func TextBool(k encoding.TextMarshaler, v bool) kvm {
	return kvm{K: k, V: Bool(v)}
}

func TextBoolp(k encoding.TextMarshaler, v *bool) kvm {
	return kvm{K: k, V: Boolp(v)}
}

func TextBytes(k encoding.TextMarshaler, v ...byte) kvm {
	return kvm{K: k, V: Bytes(v...)}
}

func TextBytesp(k encoding.TextMarshaler, v *[]byte) kvm {
	return kvm{K: k, V: Bytesp(v)}
}

func TextComplex128(k encoding.TextMarshaler, v complex128) kvm {
	return kvm{K: k, V: Complex128(v)}
}

func TextComplex128p(k encoding.TextMarshaler, v *complex128) kvm {
	return kvm{K: k, V: Complex128p(v)}
}

func TextComplex64(k encoding.TextMarshaler, v complex64) kvm {
	return kvm{K: k, V: Complex64(v)}
}

func TextComplex64p(k encoding.TextMarshaler, v *complex64) kvm {
	return kvm{K: k, V: Complex64p(v)}
}

func TextError(k encoding.TextMarshaler, v error) kvm {
	return kvm{K: k, V: Error(v)}
}

func TextFloat32(k encoding.TextMarshaler, v float32) kvm {
	return kvm{K: k, V: Float32(v)}
}

func TextFloat32p(k encoding.TextMarshaler, v *float32) kvm {
	return kvm{K: k, V: Float32p(v)}
}

func TextFloat64(k encoding.TextMarshaler, v float64) kvm {
	return kvm{K: k, V: Float64(v)}
}

func TextFloat64p(k encoding.TextMarshaler, v *float64) kvm {
	return kvm{K: k, V: Float64p(v)}
}

func TextInt(k encoding.TextMarshaler, v int) kvm {
	return kvm{K: k, V: Int(v)}
}

func TextIntp(k encoding.TextMarshaler, v *int) kvm {
	return kvm{K: k, V: Intp(v)}
}

func TextInt16(k encoding.TextMarshaler, v int16) kvm {
	return kvm{K: k, V: Int16(v)}
}

func TextInt16p(k encoding.TextMarshaler, v *int16) kvm {
	return kvm{K: k, V: Int16p(v)}
}

func TextInt32(k encoding.TextMarshaler, v int32) kvm {
	return kvm{K: k, V: Int32(v)}
}

func TextInt32p(k encoding.TextMarshaler, v *int32) kvm {
	return kvm{K: k, V: Int32p(v)}
}

func TextInt64(k encoding.TextMarshaler, v int64) kvm {
	return kvm{K: k, V: Int64(v)}
}

func TextInt64p(k encoding.TextMarshaler, v *int64) kvm {
	return kvm{K: k, V: Int64p(v)}
}

func TextInt8(k encoding.TextMarshaler, v int8) kvm {
	return kvm{K: k, V: Int8(v)}
}

func TextInt8p(k encoding.TextMarshaler, v *int8) kvm {
	return kvm{K: k, V: Int8p(v)}
}

func TextRunes(k encoding.TextMarshaler, v ...rune) kvm {
	return kvm{K: k, V: Runes(v...)}
}

func TextRunesp(k encoding.TextMarshaler, v *[]rune) kvm {
	return kvm{K: k, V: Runesp(v)}
}

func TextText(k, v encoding.TextMarshaler) kvm {
	return kvm{K: k, V: Text(v)}
}

func TextString(k encoding.TextMarshaler, v string) kvm {
	return kvm{K: k, V: String(v)}
}

func TextStringp(k encoding.TextMarshaler, v *string) kvm {
	return kvm{K: k, V: Stringp(v)}
}

func TextUint(k encoding.TextMarshaler, v uint) kvm {
	return kvm{K: k, V: Uint(v)}
}

func TextUintp(k encoding.TextMarshaler, v *uint) kvm {
	return kvm{K: k, V: Uintp(v)}
}

func TextUint16(k encoding.TextMarshaler, v uint16) kvm {
	return kvm{K: k, V: Uint16(v)}
}

func TextUint16p(k encoding.TextMarshaler, v *uint16) kvm {
	return kvm{K: k, V: Uint16p(v)}
}

func TextUint32(k encoding.TextMarshaler, v uint32) kvm {
	return kvm{K: k, V: Uint32(v)}
}

func TextUint32p(k encoding.TextMarshaler, v *uint32) kvm {
	return kvm{K: k, V: Uint32p(v)}
}

func TextUint64(k encoding.TextMarshaler, v uint64) kvm {
	return kvm{K: k, V: Uint64(v)}
}

func TextUint64p(k encoding.TextMarshaler, v *uint64) kvm {
	return kvm{K: k, V: Uint64p(v)}
}

func TextUint8(k encoding.TextMarshaler, v uint8) kvm {
	return kvm{K: k, V: Uint8(v)}
}

func TextUint8p(k encoding.TextMarshaler, v *uint8) kvm {
	return kvm{K: k, V: Uint8p(v)}
}

func TextUintptr(k encoding.TextMarshaler, v uintptr) kvm {
	return kvm{K: k, V: Uintptr(v)}
}

func TextUintptrp(k encoding.TextMarshaler, v *uintptr) kvm {
	return kvm{K: k, V: Uintptrp(v)}
}

func TextDuration(k encoding.TextMarshaler, v time.Duration) kvm {
	return kvm{K: k, V: Duration(v)}
}

func TextDurationp(k encoding.TextMarshaler, v *time.Duration) kvm {
	return kvm{K: k, V: Durationp(v)}
}

func TextTime(k encoding.TextMarshaler, v time.Time) kvm {
	return kvm{K: k, V: Time(v)}
}

func TextTimep(k encoding.TextMarshaler, v *time.Time) kvm {
	return kvm{K: k, V: Timep(v)}
}

func TextFunc(k encoding.TextMarshaler, v func() json.Marshaler) kvm {
	return kvm{K: k, V: v()}
}

func TextRaw(k encoding.TextMarshaler, v []byte) kvm {
	return kvm{K: k, V: Raw(v)}
}

func TextAny(k encoding.TextMarshaler, v interface{}) kvm {
	return kvm{K: k, V: Any(v)}
}

func TextReflect(k encoding.TextMarshaler, v interface{}) kvm {
	return kvm{K: k, V: Reflect(v)}
}

// kvl is a key-value pair implements the Leveler interface
// in addition to the KV interface (text/json marshalers).
// Level method intends to indicate severity level.
// For example syslog levels: "0" emergency;
//               						  "1" alert;
//               						  "2" critical;
//               						  "3" error;
//               						  "4" warning;
//               						  "5" notice;
//               						  "6" informational;
//               						  "7" debug;
// (https://en.wikipedia.org/wiki/Syslog#Severity_level).
type kvl struct {
	K encoding.TextMarshaler
	V json.Marshaler
	S fmt.Stringer
}

func (kv kvl) MarshalText() (text []byte, err error) { return kv.K.MarshalText() }
func (kv kvl) MarshalJSON() ([]byte, error)          { return kv.V.MarshalJSON() }
func (kv kvl) Level() string                         { return kv.S.String() }

func StringLevel(k string, v string) kvl {
	return kvl{K: String(k), V: String(v), S: String(v)}
}
