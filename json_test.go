// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plog_test

import (
	"encoding/json"
	"testing"

	"github.com/pprint/pfmt"
	"github.com/pprint/plog"
)

func TestMarshalJSON(t *testing.T) {
	tests := []marshalTest{
		{
			line:     line(),
			input:    map[string]json.Marshaler{"kv slice": pfmt.JSON(plog.StringString("foo", "bar"), plog.StringInt("xyz", 42))},
			want:     `foo "bar" xyz 42`,
			wantText: `foo "bar" xyz 42`,
			wantJSON: `{
			"kv slice":{"foo":"bar","xyz":42}
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"without jsons": pfmt.JSON()},
			want:     "",
			wantText: "",
			wantJSON: `{
			"without jsons":null
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of empty jsons": pfmt.JSON(pfmt.String(""), pfmt.String(""))},
			want:     ``,
			wantText: ``,
			wantJSON: `{
			"slice of empty jsons":{}
		}`,
		},
		{
			line:     line(),
			input:    map[string]json.Marshaler{"slice of json nils": pfmt.JSON(nil, nil)},
			want:     "",
			wantText: "",
			wantJSON: `{
			"slice of json nils":{}
		}`,
		},
	}

	testMarshal(t, tests)
}
