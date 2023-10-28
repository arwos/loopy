package server

import (
	"net/http"

	"go.osspkg.com/goppy/plugins/web"
	"go.osspkg.com/goppy/sdk/log"
	"go.osspkg.com/goppy/sdk/netutil/websocket"
)

type Props struct {
	Decode func(in interface{}) error
	Encode func(in interface{})
	Log    func(key string, err error, suffix string)
}

func (v *AppV1) ProxyWS(call func(p *Props) error) func(websocket.Response, websocket.Request, websocket.Meta) error {
	return func(w websocket.Response, r websocket.Request, _ websocket.Meta) error {
		return call(&Props{
			Decode: func(in interface{}) error {
				return r.Decode(in)
			},
			Encode: func(in interface{}) {
				w.Encode(in)
			},
			Log: func(key string, err error, suffix string) {
				v.log.WithFields(log.Fields{
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
				v.log.WithFields(log.Fields{
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
