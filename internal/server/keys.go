/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

func (v *AppV1) KVGetV1(ctx *Props) error {
	var data EntitiesKV
	if err := ctx.Decode(&data); err != nil {
		return err
	}
	if len(data) == 0 {
		return errRequestEmpty
	}
	for i := 0; i < len(data); i++ {
		if v.cacheKV.Has(data[0].Key) {
			(&data[i]).Value = v.cacheKV.Get(data[0].Key)
			continue
		}
		entity := data[i].ToDB()
		if err := v.db.GetKV(&entity); err != nil {
			ctx.Log(data[i].Key, err, "get")
			(&data[i]).UseEmptyValue()
		} else {
			(&data[i]).FromDB(entity)
		}
	}
	ctx.Encode(data)
	return nil
}

func (v *AppV1) KVSetV1(ctx *Props) error {
	var data EntitiesKV
	if err := ctx.Decode(&data); err != nil {
		return err
	}
	if len(data) == 0 {
		return errRequestEmpty
	}
	for _, datum := range data {
		if err := v.db.SetKV(datum.ToDB()); err != nil {
			ctx.Log(datum.Key, err, "set")
			return err
		}
		v.cacheKV.Set(datum.Key, datum.Value)
	}
	ctx.Encode(doneResponse)
	return nil
}

func (v *AppV1) KVDelV1(ctx *Props) error {
	var data EntitiesKV
	if err := ctx.Decode(&data); err != nil {
		return err
	}
	if len(data) == 0 {
		return errRequestEmpty
	}
	for _, datum := range data {
		if err := v.db.DelKV(datum.ToDB()); err != nil {
			ctx.Log(datum.Key, err, "del")
			return err
		}
		v.cacheKV.Del(datum.Key)
		(&datum).UseEmptyValue()
	}
	ctx.Encode(doneResponse)
	return nil
}

func (v *AppV1) KVSearchV1(ctx *Props) error {
	var data EntityKV
	if err := ctx.Decode(&data); err != nil {
		return err
	}
	list, err := v.db.SearchKV(data.ToDB().Key)
	if err != nil {
		ctx.Log(data.Key, err, "search")
		return err
	}
	result := make(EntitiesKV, 0, len(list))
	for _, item := range list {
		entity := &EntityKV{}
		entity.FromDB(item)
		result = append(result, *entity)
	}
	ctx.Encode(result)
	return nil
}

func (v *AppV1) KVListV1(ctx *Props) error {
	var data EntityKV
	if err := ctx.Decode(&data); err != nil {
		return err
	}
	list, err := v.db.ListKV(data.ToDB().Key)
	if err != nil {
		ctx.Log(data.Key, err, "list")
		return err
	}
	result := make(EntitiesKV, 0, len(list))
	for _, item := range list {
		entity := &EntityKV{}
		entity.FromDB(item)
		result = append(result, *entity)
	}
	ctx.Encode(result)
	return nil
}
