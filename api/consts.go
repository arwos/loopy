/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

const (
	PathApiV1 = "/api/watch/v1"
)

const (
	PathApiV1KV       = "/api/kv/v1"
	PathApiV1KVSearch = "/api/kv/search/v1"
	PathApiV1KVList   = "/api/kv/list/v1"
)

const (
	AuthTokenHeaderName = "x-loop-auth"
)

const (
	EventKVSet = 100001
	EventKVGet = 100002
	EventKVDel = 100003
)
