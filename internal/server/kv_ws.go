/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

//func (v *AppV1) SetKV(ev web.WebsocketEventer, c web.WebsocketServerProcessor) error {
//	var data KVModel
//	if err := ev.Decode(&data); err != nil {
//		fmt.Println(err)
//		return err
//	}
//	fmt.Println(string(data.Key), string(data.Value))
//	if err := v.db.SetKV(data.toKVItem()); err != nil {
//		return err
//	}
//	c.EncodeEvent(ev, "ok")
//	return nil
//}
//
//func (v *AppV1) GetKV(ev web.WebsocketEventer, c web.WebsocketServerProcessor) error {
//	data := &KVModel{}
//	if err := ev.Decode(data); err != nil {
//		return err
//	}
//	kiv := data.toKVItem()
//	if err := v.db.GetKV(&kiv); err != nil {
//		return err
//	}
//	data.fromKVItem(kiv)
//	c.EncodeEvent(ev, &data)
//	return nil
//}
//
//func (v *AppV1) DelKV(ev web.WebsocketEventer, c web.WebsocketServerProcessor) error {
//	var data KVModel
//	if err := ev.Decode(&data); err != nil {
//		return err
//	}
//	if err := v.db.DelKV(data.toKVItem()); err != nil {
//		return err
//	}
//	c.EncodeEvent(ev, "ok")
//	return nil
//}
