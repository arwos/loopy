/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"go.arwos.org/loopy/api"
	"go.osspkg.com/goppy/plugins"
	"go.osspkg.com/goppy/plugins/web"
)

var Plugin = plugins.Plugin{
	Inject: New,
	Resolve: func(r web.RouterPool, c *AppV1, ws web.WebsocketServer) {
		router := r.Main()

		//ws.Event(c.SetKV, api.EventKVSet)
		//ws.Event(c.GetKV, api.EventKVGet)
		//ws.Event(c.DelKV, api.EventKVDel)
		//router.Get(api.PathApiV1, ws.Handling)

		router.Get(api.PathApiV1KV, c.KVGetV1)
		router.Put(api.PathApiV1KV, c.KVSetV1)
		router.Delete(api.PathApiV1KV, c.KVDelV1)
		router.Get(api.PathApiV1KVSearch, c.KVSearchV1)
		router.Get(api.PathApiV1KVList, c.KVListV1)
	},
}
