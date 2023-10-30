/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"context"
	"net/http"
)

func (v *_client) Get(ctx context.Context, key string) (string, error) {
	data := make(EntitiesKV, 0, 1)
	data = append(data, EntityKV{
		Key: key,
	})
	if err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KV), &data, &data); err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", errRequestEmpty
	}
	return data[0].Value.String(), nil
}

func (v *_client) Set(ctx context.Context, key, value string) error {
	data := make(EntitiesKV, 0, 1)
	data = append(data, EntityKV{
		Key:   key,
		Value: []byte(value),
	})
	return v.cli.Call(ctx, http.MethodPut, v.buildUri(PathApiV1KV), &data, nil)
}

func (v *_client) Delete(ctx context.Context, key string) error {
	data := make(EntitiesKV, 0, 1)
	data = append(data, EntityKV{
		Key: key,
	})
	return v.cli.Call(ctx, http.MethodDelete, v.buildUri(PathApiV1KV), &data, nil)
}

func (v *_client) Search(ctx context.Context, prefix string) ([]EntityKV, error) {
	data := EntityKV{
		Key: prefix,
	}
	var result EntitiesKV
	err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KVSearch), &data, &result)
	if len(result) == 0 {
		return nil, errRequestEmpty
	}
	return result, err
}

func (v *_client) List(ctx context.Context, prefix string) ([]EntityKV, error) {
	data := EntityKV{
		Key: prefix,
	}
	var result EntitiesKV
	err := v.cli.Call(ctx, http.MethodGet, v.buildUri(PathApiV1KVList), &data, &result)
	return result, err
}
