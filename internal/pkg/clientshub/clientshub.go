/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package clientshub

import (
	"strings"

	"go.osspkg.com/goppy/iosync"
)

type Hub struct {
	cid  map[string]*Maps
	keys map[string]*Maps
	mux  iosync.Lock
}

func New() *Hub {
	return &Hub{
		cid:  make(map[string]*Maps, 100),
		keys: make(map[string]*Maps, 100),
		mux:  iosync.NewLock(),
	}
}

func (v *Hub) Add(cid string, keys []string) {
	v.mux.Lock(func() {
		cm, ok := v.cid[cid]
		if !ok {
			cm = NewMaps()
			v.cid[cid] = cm
		}
		cm.Set(keys...)

		for _, key := range keys {
			km, ok := v.keys[key]
			if !ok {
				km = NewMaps()
				v.keys[key] = km
			}
			km.Set(cid)
		}
	})
}

func (v *Hub) Del(cid string, keys []string) {
	v.mux.Lock(func() {
		cm, ok := v.cid[cid]
		if !ok {
			return
		}
		cm.Del(keys...)

		for _, key := range keys {
			km, ok := v.keys[key]
			if !ok {
				continue
			}
			km.Del(cid)
		}
	})
}

func (v *Hub) GetClients(key string) (out []string) {
	v.mux.RLock(func() {
		for k, km := range v.keys {
			if key == k || strings.HasPrefix(key, k) {
				out = append(out, km.Get()...)
			}
		}
	})
	return
}

func (v *Hub) HasClient(cid string) (out bool) {
	v.mux.RLock(func() {
		_, out = v.cid[cid]
	})
	return
}

func (v *Hub) DelClient(cid string) {
	v.mux.Lock(func() {
		cm, ok := v.cid[cid]
		if !ok {
			return
		}
		keys := cm.Get()
		delete(v.cid, cid)

		for _, key := range keys {
			km, ok := v.keys[key]
			if !ok {
				continue
			}
			km.Del(cid)
			if km.IsEmpty() {
				delete(v.keys, key)
			}
		}
	})
}
