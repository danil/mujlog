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

var MarshalErrorTests = []marshalTests{
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
}

func TestMarshalError(t *testing.T) {
	testMarshal(t, MarshalErrorTests)
}
