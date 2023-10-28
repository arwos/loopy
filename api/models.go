/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"bytes"
	"fmt"
)

//go:generate easyjson

//easyjson:json
type EntityKV struct {
	Key   string   `json:"k"`
	Value RawValue `json:"v,omitempty"`
}

//easyjson:json
type EntitiesKV []EntityKV

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	esc   = []byte("\\\"")
	unesc = []byte("\"")
)

type RawValue []byte

func (m RawValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString("\"")
	buf.Write(bytes.ReplaceAll(m, unesc, esc))
	buf.WriteString("\"")
	return buf.Bytes(), nil
}

func (m *RawValue) UnmarshalJSON(data []byte) error {
	if m == nil {
		return fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	out := bytes.ReplaceAll(data[1:len(data)-1], esc, unesc)
	*m = append((*m)[0:0], out...)
	return nil
}

func (m RawValue) String() string {
	if m == nil {
		return "null"
	}
	return string(m)
}
