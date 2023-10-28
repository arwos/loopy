/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"go.osspkg.com/goppy/sdk/errors"
	"go.osspkg.com/goppy/sdk/netutil/websocket"
)

const (
	PathApiV1Watch = "/api/watch/v1"
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
	EventKVGet websocket.EventID = iota + 1
	EventKVSet
	EventKVDel
	EventKVWatch
	EventKVUnWatch
	EventKVWatchValue
)

var (
	errRequestEmpty = errors.New("request is empty")
)
