/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"net/http"

	"go.osspkg.com/goppy/plugins/web"
)

func (v *AppV1) KVSetV1(ctx web.Context) {
	var data KVModel
	if err := ctx.BindJSON(&data); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	if err := v.db.SetKV(data.toKVItem()); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	ctx.String(http.StatusOK, DoneResponse)
}

func (v *AppV1) KVGetV1(ctx web.Context) {
	data := &KVModel{}
	if err := ctx.BindJSON(data); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	kiv := data.toKVItem()
	if err := v.db.GetKV(&kiv); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	data.fromKVItem(kiv)
	ctx.JSON(http.StatusOK, data)
}

func (v *AppV1) KVDelV1(ctx web.Context) {
	var data KVModel
	if err := ctx.BindJSON(&data); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	if err := v.db.DelKV(data.toKVItem()); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	ctx.String(http.StatusOK, DoneResponse)
}

func (v *AppV1) KVSearchV1(ctx web.Context) {
	var data KVModel
	if err := ctx.BindJSON(&data); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	list, err := v.db.SearchKV(data.toKVItem().Key)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	result := make([]KVModel, 0, len(list))
	for _, item := range list {
		kvi := &KVModel{}
		kvi.fromKVItem(item)
		result = append(result, *kvi)
	}
	ctx.JSON(http.StatusOK, result)
}

func (v *AppV1) KVListV1(ctx web.Context) {
	var data KVModel
	if err := ctx.BindJSON(&data); err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	list, err := v.db.ListKV(data.toKVItem().Key)
	if err != nil {
		ctx.Error(http.StatusBadRequest, err)
		return
	}
	result := make([]KVModel, 0, len(list))
	for _, item := range list {
		kvi := &KVModel{}
		kvi.fromKVItem(item)
		result = append(result, *kvi)
	}
	ctx.JSON(http.StatusOK, result)
}
