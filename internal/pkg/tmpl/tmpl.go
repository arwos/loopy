/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package tmpl

import (
	"context"
	"strings"
	"time"

	"go.arwos.org/loopy/internal/pkg/cachekeys"
	"go.osspkg.com/goppy/iosync"
	"go.osspkg.com/goppy/routine"
	"go.osspkg.com/goppy/xc"
	"go.osspkg.com/goppy/xlog"
)

type Tmpl struct {
	keys    cachekeys.Mapper
	tmpls   map[string]*Item
	handler func(key string)
	trigger chan struct{}
	log     xlog.Logger
	mux     iosync.Lock
	wg      iosync.Group
}

func New(l xlog.Logger) *Tmpl {
	return &Tmpl{
		keys:    cachekeys.NewWithoutBus(),
		tmpls:   make(map[string]*Item, 10),
		handler: func(key string) {},
		log:     l,
		trigger: make(chan struct{}, 2),
		mux:     iosync.NewLock(),
		wg:      iosync.NewGroup(),
	}
}

func (v *Tmpl) KeysHandler(call func(key string)) {
	v.handler = call
}

func (v *Tmpl) SetKey(key string, value *string) {
	v.mux.Lock(func() {
		v.keys.Set(key, value)
	})
	v.sendTrigger()
}

func (v *Tmpl) Query(key string) {
	v.handler(key)
}

func (v *Tmpl) Keys() cachekeys.MapGetter {
	return v.keys
}

func (v *Tmpl) Add(in, out string) (err error) {
	v.mux.Lock(func() {
		var t *Item
		if t, err = NewItem(in, out, v); err != nil {
			return
		}
		if err = t.Prepare(); err != nil {
			v.log.WithError("err", err).Errorf("prepare template")
			return
		}
		v.tmpls[in+":"+out] = t
	})

	return
}

func (v *Tmpl) sendTrigger() {
	select {
	case v.trigger <- struct{}{}:
	default:
	}
}

func (v *Tmpl) Watch(ctx xc.Context, interval time.Duration) {
	v.wg.Background(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-v.trigger:
				v.mux.RLock(func() {
					for _, tmpl := range v.tmpls {
						if err := tmpl.Update(); err != nil {
							v.log.WithError("err", err).Errorf("update template")
						}
					}
				})
			}
		}
	})
	routine.Interval(ctx.Context(), interval, func(ctx context.Context) {
		var (
			links []string
		)
		v.mux.RLock(func() {
			for link, tmpl := range v.tmpls {
				if tmpl.IsChanged() {
					v.log.WithField("path", link).Infof("reload template")
					links = append(links, link)
				}
			}
		})
		for _, link := range links {
			f := strings.Split(link, ":")
			if len(f) != 2 {
				v.log.WithField("path", link).Errorf("invalid template path")
			}
			err := v.Add(f[0], f[1])
			if err != nil {
				v.log.WithError("err", err).Errorf("reload template")
				continue
			}
		}
	})
	v.wg.Wait()
}
