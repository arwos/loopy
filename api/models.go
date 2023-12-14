/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

//go:generate easyjson

//easyjson:json
type (
	EntitiesKV []EntityKV
	EntityKV   struct {
		Key string  `json:"k"`
		Val *string `json:"v"`
	}
)

func (v EntityKV) ValueStrOrNull() string {
	if v.Val == nil {
		return "null"
	}
	return *v.Val
}

func (v EntityKV) Value() string {
	if v.Val == nil {
		return ""
	}
	return *v.Val
}
