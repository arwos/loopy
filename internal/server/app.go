/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"go.arwos.org/loopy/api"
	"go.arwos.org/loopy/internal/pkg/cachekeys"
	"go.arwos.org/loopy/internal/pkg/clientshub"
	"go.arwos.org/loopy/internal/pkg/db"
	"go.osspkg.com/goppy/ws/event"
	"go.osspkg.com/goppy/xc"
	"go.osspkg.com/goppy/xlog"
)

type AppV1 struct {
	db      *db.DB
	hub     *clientshub.Hub
	cacheKV cachekeys.MapperWithBus
	log     xlog.Logger
	bh      func(eid event.Id, m interface{}, cids ...string)
}

func New(db *db.DB, l xlog.Logger) *AppV1 {
	return &AppV1{
		db:      db,
		hub:     clientshub.New(),
		cacheKV: cachekeys.NewWithBus(),
		log:     l,
		bh:      func(eid event.Id, m interface{}, cids ...string) {},
	}
}

func (v *AppV1) SetBroadcastHandler(call func(eid event.Id, m interface{}, cids ...string)) {
	v.bh = call
}

func (v *AppV1) Up(ctx xc.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case value := <-v.cacheKV.Bus():
				entities := make(EntitiesKV, 0, 1)
				entities = append(entities, EntityKV{Key: value.Key, Value: value.Value})
				v.Broadcast(value.Key, api.EventKVWatchValue, entities)
			}
		}
	}()
	return nil
}

func (v *AppV1) Down() error {
	return nil
}
