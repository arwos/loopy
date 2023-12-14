/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"net/http"

	"go.osspkg.com/goppy/web"
	"go.osspkg.com/goppy/ws/server"
	"go.osspkg.com/goppy/xlog"
)

type Props struct {
	Decode func(in interface{}) error
	Encode func(in interface{})
	Log    func(key string, err error, suffix string)
}

func (v *AppV1) ProxyWS(call func(p *Props) error) func(server.Response, server.Request, server.Meta) error {
	return func(w server.Response, r server.Request, _ server.Meta) error {
		return call(&Props{
			Decode: func(in interface{}) error {
				return r.Decode(in)
			},
			Encode: func(in interface{}) {
				w.Encode(in)
			},
			Log: func(key string, err error, suffix string) {
				v.log.WithFields(xlog.Fields{
					"key": key,
					"err": err.Error(),
				}).Warnf("ws %s key error", suffix)
			},
		})
	}
}

func (v *AppV1) ProxyRest(call func(p *Props) error) func(ctx web.Context) {
	return func(ctx web.Context) {
		err := call(&Props{
			Decode: func(in interface{}) error {
				return ctx.BindJSON(in)
			},
			Encode: func(in interface{}) {
				ctx.JSON(http.StatusOK, in)
			},
			Log: func(key string, err error, suffix string) {
				v.log.WithFields(xlog.Fields{
					"key": key,
					"err": err.Error(),
				}).Warnf("rest %s key error", suffix)
			},
		})
		if err != nil {
			ctx.Error(http.StatusBadRequest, err)
		}
	}
}
