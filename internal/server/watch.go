package server

import (
	"go.arwos.org/loopy/api"
	"go.osspkg.com/goppy/sdk/netutil/websocket"
)

func (v *AppV1) Broadcast(key string, eid websocket.EventID, m interface{}) {
	cids := v.hub.GetClients(key)
	if len(cids) == 0 {
		return
	}
	v.bh(eid, m, cids...)
}

func (v *AppV1) KVWatchV1(w websocket.Response, r websocket.Request, m websocket.Meta) error {
	var data EntitiesKV
	if err := r.Decode(&data); err != nil {
		return err
	}
	if len(data) == 0 {
		return errRequestEmpty
	}

	keys := make([]string, 0, len(data))
	for _, datum := range data {
		keys = append(keys, datum.Key)
	}

	if !v.hub.HasClient(m.ConnectID()) {
		m.OnClose(func(cid string) {
			v.hub.DelClient(cid)
		})
	}

	for i := 0; i < len(data); i++ {
		entity := data[i].ToDB()
		if err := v.db.GetKV(&entity); err != nil {
			(&data[i]).UseEmptyValue()
		} else {
			(&data[i]).FromDB(entity)
		}
	}
	v.hub.Add(m.ConnectID(), keys)
	w.EncodeEvent(api.EventKVWatchValue, data)

	return nil
}
