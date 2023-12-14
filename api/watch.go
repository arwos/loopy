/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"context"
	"net/url"

	"go.osspkg.com/goppy/ws/client"
	"go.osspkg.com/goppy/xlog"
)

type (
	_watch struct {
		conf *Config
		log  xlog.Logger
		cli  client.Client
	}
	Watch interface {
		Open() error
		AfterOpened(call func())
		Close()
		AfterClosed(call func())
		KeyHandler(call func(e EntitiesKV))
		KeySubscribe(keys ...string)
	}
)

func NewWatch(ctx context.Context, c *Config, l xlog.Logger) (Watch, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	w := &_watch{
		conf: c,
		log:  l,
	}
	opts := make([]func(client.Option), 0)
	if len(c.AuthToken) > 0 {
		opts = append(opts, func(co client.Option) {
			co.Header(AuthTokenHeaderName, c.AuthToken)
		})
	}
	w.cli = client.New(ctx, w.buildUri(PathApiV1Watch), l, opts...)
	return w, nil
}

func (v *_watch) Open() error {
	return v.cli.DialAndListen()
}

func (v *_watch) AfterOpened(call func()) {
	v.cli.OnOpen(func(cid string) {
		call()
	})
}

func (v *_watch) AfterClosed(call func()) {
	v.cli.OnClose(func(cid string) {
		call()
	})
}

func (v *_watch) Close() {
	v.cli.Close()
}

func (v *_watch) buildUri(path string) string {
	uri := &url.URL{
		Path:   path,
		Host:   v.conf.HostPort,
		Scheme: "ws",
	}
	if v.conf.SSL {
		uri.Scheme = "wss"
	}
	return uri.String()
}

func (v *_watch) KeyHandler(call func(e EntitiesKV)) {
	v.cli.SetHandler(func(w client.Request, r client.Response, m client.Meta) {
		var entities EntitiesKV
		if err := r.Decode(&entities); err != nil {
			v.log.WithError("err", err).Errorf("decode event")
			return
		}
		call(entities)
	}, EventKVWatchValue)
}

func (v *_watch) KeySubscribe(keys ...string) {
	var entities EntitiesKV
	for _, key := range keys {
		entities = append(entities, EntityKV{
			Key: key,
		})
	}
	v.cli.SendEvent(EventKVWatch, entities)
}
