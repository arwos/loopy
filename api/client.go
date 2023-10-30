/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.osspkg.com/goppy/sdk/webutil"
)

type (
	_client struct {
		conf *Config
		cli  *webutil.ClientHttp
	}
	Client interface {
		Get(ctx context.Context, key string) (string, error)
		Set(ctx context.Context, key, value string) error
		Delete(ctx context.Context, key string) error
		Search(ctx context.Context, prefix string) ([]EntityKV, error)
		List(ctx context.Context, prefix string) ([]EntityKV, error)
	}
)

func NewKV(c *Config) (Client, error) {
	opts := make([]webutil.ClientHttpOption, 0, 2)
	if len(c.AuthToken) > 0 {
		opts = append(opts, webutil.ClientHttpOptionHeaders(AuthTokenHeaderName, c.AuthToken))
	}
	opts = append(opts, clientHttpOptionCodec())
	cli := webutil.NewClientHttp(opts...)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return &_client{
		cli:  cli,
		conf: c,
	}, nil
}

func (v *_client) buildUri(path string) string {
	uri := &url.URL{
		Path:   path,
		Host:   v.conf.HostPort,
		Scheme: "http",
	}
	if v.conf.SSL {
		uri.Scheme = "https"
	}
	return uri.String()
}

func clientHttpOptionCodec() webutil.ClientHttpOption {
	return webutil.ClientHttpOptionCodec(
		func(in interface{}) (body []byte, contentType string, err error) {
			switch v := in.(type) {
			case []byte:
				return v, "", nil
			case json.Marshaler:
				body, err = v.MarshalJSON()
				return body, "application/json; charset=utf-8", err
			case string:
				return []byte(v), "text/plain; charset=utf-8", err
			default:
				return nil, "", fmt.Errorf("unknown request format %T", in)
			}
		},
		func(code int, _ string, body []byte, out interface{}) error {
			switch v := out.(type) {
			case *[]byte:
				*v = append(*v, body...)
				return nil
			case json.Unmarshaler:
				return v.UnmarshalJSON(body)
			case *string:
				*v = string(body)
				return nil
			default:
				if code == 200 {
					return nil
				}
				return fmt.Errorf("%d %s\n%s", code, http.StatusText(code), string(body))
			}
		},
	)
}
