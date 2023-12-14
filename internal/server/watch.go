/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"go.arwos.org/loopy/api"
	"go.osspkg.com/goppy/ws/event"
	"go.osspkg.com/goppy/ws/server"
)

func (v *AppV1) Broadcast(key string, eid event.Id, m interface{}) {
	cids := v.hub.GetClients(key)
	if len(cids) == 0 {
		return
	}
	v.bh(eid, m, cids...)
}

func (v *AppV1) KVWatchV1(w server.Response, r server.Request, m server.Meta) error {
	var req EntitiesKV
	if err := r.Decode(&req); err != nil {
		return err
	}
	if len(req) == 0 {
		return errRequestEmpty
	}

	keys := make([]string, 0, len(req))
	for _, datum := range req {
		keys = append(keys, datum.Key)
	}

	if !v.hub.HasClient(m.ConnectID()) {
		m.OnClose(func(cid string) {
			v.hub.DelClient(cid)
		})
	}

	resp := make(EntitiesKV, 0, len(req)*2)
	for i := 0; i < len(req); i++ {
		entity := req[i].ToDB()
		list, err := v.db.SearchKV(entity.Key)
		if err != nil {
			continue
		}
		for _, kv := range list {
			item := EntityKV{}
			item.FromDB(kv)
			resp = append(resp, item)
		}
	}
	v.hub.Add(m.ConnectID(), keys)
	w.EncodeEvent(api.EventKVWatchValue, resp)

	return nil
}
