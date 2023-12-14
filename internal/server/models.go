/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

//go:generate easyjson

import (
	"go.arwos.org/loopy/internal/pkg/db"
)

//easyjson:json
type (
	EntitiesKV []EntityKV
	EntityKV   struct {
		Key   string  `json:"k"`
		Value *string `json:"v"`
	}
)

func (v EntityKV) ToDB() db.EntityKV {
	kvi := db.EntityKV{
		Key: []byte(v.Key),
	}
	if v.Value != nil {
		kvi.Value = []byte(*v.Value)
	}
	return kvi
}

func (v *EntityKV) FromDB(item db.EntityKV) {
	v.Key = string(item.Key)
	if len(item.Value) == 0 {
		v.Value = nil
	} else {
		s := string(item.Value)
		v.Value = &s
	}
}

func (v *EntityKV) UseEmptyValue() {
	if v.Value != nil {
		v.Value = nil
	}
}

//easyjson:json
type (
	EntitiesService []EntityService
	EntityService   struct {
		Name    string   `json:"n"`
		Address string   `json:"a,omitempty"`
		Tags    []string `json:"t,omitempty"`
		Health  string   `json:"h,omitempty"`
	}
)
