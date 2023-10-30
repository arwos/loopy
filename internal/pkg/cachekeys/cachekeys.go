package cachekeys

import "go.osspkg.com/goppy/sdk/iosync"

type (
	Value struct {
		Key   string
		Value []byte
	}
	Map struct {
		data map[string][]byte
		bus  chan Value
		mux  iosync.Lock
	}
)

func New() *Map {
	return &Map{
		data: make(map[string][]byte, 100),
		bus:  make(chan Value, 1000),
		mux:  iosync.NewLock(),
	}
}

func (v *Map) Bus() <-chan Value {
	return v.bus
}

func (v *Map) toBus(key string, value []byte) {
	select {
	case v.bus <- Value{Key: key, Value: value}:
	default:
	}
}

func (v *Map) Set(key string, value []byte) {
	v.mux.Lock(func() {
		tmp := make([]byte, 0, len(value))
		copy(tmp, value)
		v.data[key] = tmp
		v.toBus(key, value)
	})
}

func (v *Map) Get(key string) (out []byte) {
	v.mux.RLock(func() {
		value, ok := v.data[key]
		if !ok {
			return
		}
		out = append(out, value...)
	})
	return
}

func (v *Map) Has(key string) (ok bool) {
	v.mux.RLock(func() {
		_, ok = v.data[key]
	})
	return
}

func (v *Map) Del(key string) {
	v.mux.Lock(func() {
		delete(v.data, key)
		v.toBus(key, nil)
	})
}
