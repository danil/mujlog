// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/plog"
)

var MarshalBoolpTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"bool pointer to true": plog.Boolp(&b)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := false
			return map[string]json.Marshaler{"bool pointer to false": plog.Boolp(&b)}
		}(),
		want:     "false",
		wantText: "false",
		wantJSON: `{
			"bool pointer to false":false
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"bool nil pointer": plog.Boolp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"bool nil pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"any bool pointer to true": plog.Any(&b)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"any twice/nested pointer to bool true": plog.Any(&b2)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"reflect bool pointer to true": plog.Reflect(&b)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"reflect bool twice/nested pointer to true": plog.Reflect(&b2)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var b *bool
			return map[string]json.Marshaler{"reflect bool pointer to nil": plog.Reflect(b)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect bool pointer to nil":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"any bool pointer to true": plog.Any(&b)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"any bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"any twice/nested pointer to bool true": plog.Any(&b2)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"any twice/nested pointer to bool true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			return map[string]json.Marshaler{"reflect bool pointer to true": plog.Reflect(&b)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"reflect bool pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			b := true
			b2 := &b
			return map[string]json.Marshaler{"reflect bool twice/nested pointer to true": plog.Reflect(&b2)}
		}(),
		want:     "true",
		wantText: "true",
		wantJSON: `{
			"reflect bool twice/nested pointer to true":true
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			var b *bool
			return map[string]json.Marshaler{"reflect bool pointer to nil": plog.Reflect(b)}
		}(),
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"reflect bool pointer to nil":null
		}`,
	},
}

func TestMarshalBoolp(t *testing.T) {
	testMarshal(t, MarshalBoolpTests)
}
