/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

//go:generate easyjson

import (
	"go.arwos.org/loopy/api"
	"go.arwos.org/loopy/internal/pkg/db"
)

//easyjson:json
type KVModel struct {
	Key   string       `json:"key"`
	Value api.RawValue `json:"value,omitempty"`
}

func (v *KVModel) toKVItem() db.KVItem {
	kvi := db.KVItem{
		Key: []byte(v.Key),
	}
	if v.Value != nil {
		kvi.Value = v.Value
	}
	return kvi
}
func (v *KVModel) fromKVItem(item db.KVItem) {
	v.Key = string(item.Key)
	v.Value = item.Value
}
