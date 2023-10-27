/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"context"
	"net/http"
)

func (v *Client) Set(ctx context.Context, key, value string) (string, error) {
	data := KVModel{
		Key:   key,
		Value: []byte(value),
	}
	var result string
	err := v.cli.Call(ctx, http.MethodPut, v.buildUri(PathApiV1KV), &data, &result)
	return result, err
}

func (v *Client) Get(ctx context.Context, key string) (string, error) {
	data := &KVModel{
		Key: key,
	}
	err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KV), data, data)
	return data.Value.String(), err
}

func (v *Client) Delete(ctx context.Context, key string) (string, error) {
	data := KVModel{
		Key: key,
	}
	var result string
	err := v.cli.Call(ctx, http.MethodDelete, v.buildUri(PathApiV1KV), &data, &result)
	return result, err
}

func (v *Client) Search(ctx context.Context, prefix string) ([]KVModel, error) {
	data := KVModel{
		Key: prefix,
	}
	var result KVListModel
	err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KVSearch), &data, &result)
	return result, err
}

func (v *Client) List(ctx context.Context, prefix string) ([]KVModel, error) {
	data := KVModel{
		Key: prefix,
	}
	var result KVListModel
	err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KVList), &data, &result)
	return result, err
}
