/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"fmt"
)

//go:generate easyjson

//easyjson:json
type KVModel struct {
	Key   string   `json:"key"`
	Value RawValue `json:"value,omitempty"`
}

//easyjson:json
type KVListModel []KVModel

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type RawValue []byte

func (m RawValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *RawValue) UnmarshalJSON(data []byte) error {
	if m == nil {
		return fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (m RawValue) String() string {
	if m == nil {
		return "null"
	}
	return string(m)
}
