/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

//type Watch struct {
//	cli web.WebsocketClientConn
//}

//func NewWatch(ctx context.Context, c *Config, cli web.WebsocketClient) (*Watch, error) {
//	conn, err := cli.Create(ctx, "ws://"+c.Addr+c.ApiPath, func(o web.WebsocketClientOption) {
//		if len(c.AuthToken) > 0 {
//			o.Header(AuthTokenHeaderName, c.AuthToken)
//		}
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &Watch{
//		cli: conn,
//	}, nil
//}
//
//func (v *Watch) Open() error {
//	return v.cli.Run()
//}
//
//func (v *Watch) Close() {
//	v.cli.Close()
//}
