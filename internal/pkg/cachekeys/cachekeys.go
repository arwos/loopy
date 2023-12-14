/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package cachekeys

import "go.osspkg.com/goppy/iosync"

type (
	Value struct {
		Key   string
		Value *string
	}
	Map struct {
		data    map[string]*string
		bus     chan Value
		withBus bool
		mux     iosync.Lock
	}

	MapGetter interface {
		Get(key string) (value *string)
		Has(key string) (ok bool)
		Each(call func(key string, value *string))
	}

	MapSetter interface {
		Set(key string, value *string)
		Del(key string)
	}

	Mapper interface {
		MapGetter
		MapSetter
	}

	MapperWithBus interface {
		Bus() <-chan Value
		MapGetter
		MapSetter
	}
)

func NewWithBus() MapperWithBus {
	return &Map{
		data:    make(map[string]*string, 100),
		bus:     make(chan Value, 1000),
		withBus: true,
		mux:     iosync.NewLock(),
	}
}

func NewWithoutBus() Mapper {
	return &Map{
		data:    make(map[string]*string, 100),
		bus:     make(chan Value, 1000),
		withBus: false,
		mux:     iosync.NewLock(),
	}
}

func (v *Map) Bus() <-chan Value {
	return v.bus
}

func (v *Map) toBus(key string, value *string) {
	if !v.withBus {
		return
	}
	select {
	case v.bus <- Value{Key: key, Value: value}:
	default:
	}
}

func (v *Map) Set(key string, value *string) {
	v.mux.Lock(func() {
		v.data[key] = value
		v.toBus(key, value)
	})
}

func (v *Map) Get(key string) (value *string) {
	v.mux.RLock(func() {
		d, ok := v.data[key]
		if !ok {
			return
		}
		value = d
	})
	return
}

func (v *Map) Has(key string) (ok bool) {
	v.mux.RLock(func() {
		_, ok = v.data[key]
	})
	return
}

func (v *Map) Each(call func(key string, value *string)) {
	v.mux.RLock(func() {
		for key, val := range v.data {
			call(key, val)
		}
	})
	return
}

func (v *Map) Del(key string) {
	v.mux.Lock(func() {
		delete(v.data, key)
		v.toBus(key, nil)
	})
}
