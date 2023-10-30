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

		c.SetBroadcastHandler(ws.SendEvent)

		ws.SetHandler(c.ProxyWS(c.KVGetV1), api.EventKVGet)
		ws.SetHandler(c.ProxyWS(c.KVSetV1), api.EventKVSet)
		ws.SetHandler(c.ProxyWS(c.KVDelV1), api.EventKVDel)
		ws.SetHandler(c.KVWatchV1, api.EventKVWatch)

		router.Get(api.PathApiV1Watch, func(ctx web.Context) {
			ws.Handling(ctx.Response(), ctx.Request())
		})

		router.Get(api.PathApiV1KV, c.ProxyRest(c.KVGetV1))
		router.Put(api.PathApiV1KV, c.ProxyRest(c.KVSetV1))
		router.Delete(api.PathApiV1KV, c.ProxyRest(c.KVDelV1))
		router.Get(api.PathApiV1KVSearch, c.ProxyRest(c.KVSearchV1))
		router.Get(api.PathApiV1KVList, c.ProxyRest(c.KVListV1))
	},
}
