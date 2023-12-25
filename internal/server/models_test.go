/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"encoding/json"
	"testing"

	"go.osspkg.com/goppy/xtest"
)

func TestUnit_EntityKV(t *testing.T) {
	m := EntityKV{
		Key: "1",
	}
	b, err := json.Marshal(m)
	xtest.NoError(t, err)
	xtest.Equal(t, string(b), `{"k":"1","v":null}`)

	d := "123"
	m = EntityKV{
		Key:   "1",
		Value: &d,
	}
	b, err = json.Marshal(m)
	xtest.NoError(t, err)
	xtest.Equal(t, string(b), `{"k":"1","v":"123"}`)
}
