// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/pprint/plog"
)

var MarshalErrorpTests = []marshalTests{
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"error pointer": plog.Errorp(&err)}
		}(),
		want:     "something went wrong",
		wantText: "something went wrong",
		wantJSON: `{
			"error pointer":"something went wrong"
		}`,
	},
	{
		line:     line(),
		input:    map[string]json.Marshaler{"nil error pointer": plog.Errorp(nil)},
		want:     "null",
		wantText: "null",
		wantJSON: `{
			"nil error pointer":null
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"any error pointer": plog.Any(&err)}
		}(),
		want:     "something went wrong",
		wantText: "something went wrong",
		wantJSON: `{
			"any error pointer":"something went wrong"
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			err2 := &err
			return map[string]json.Marshaler{"any twice/nested pointer to error": plog.Any(&err2)}
		}(),
		want:     "{something went wrong}",
		wantText: "{something went wrong}",
		wantJSON: `{
			"any twice/nested pointer to error":{}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			return map[string]json.Marshaler{"reflect error pointer": plog.Reflect(&err)}
		}(),
		want:     "{something went wrong}",
		wantText: "{something went wrong}",
		wantJSON: `{
			"reflect error pointer":{}
		}`,
	},
	{
		line: line(),
		input: func() map[string]json.Marshaler {
			err := errors.New("something went wrong")
			err2 := &err
			return map[string]json.Marshaler{"reflect twice/nested pointer to error": plog.Reflect(&err2)}
		}(),
		want:     "{something went wrong}",
		wantText: "{something went wrong}",
		wantJSON: `{
			"reflect twice/nested pointer to error":{}
		}`,
	},
}

func TestErrorpMarshal(t *testing.T) {
	testMarshal(t, MarshalErrorpTests)
}
