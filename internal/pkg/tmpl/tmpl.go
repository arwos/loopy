package tmpl

import (
	"context"
	"text/template"

	"go.osspkg.com/goppy/sdk/iosync"
	"go.osspkg.com/goppy/sdk/log"
)

type Tmpl struct {
	keys    map[string]string
	tmpls   []*Item
	keyH    func(keys ...string)
	log     log.Logger
	trigger chan struct{}
	mux     iosync.Lock
}

func New(l log.Logger) *Tmpl {
	return &Tmpl{
		keys:    make(map[string]string, 100),
		tmpls:   make([]*Item, 0, 2),
		keyH:    func(keys ...string) {},
		log:     l,
		trigger: make(chan struct{}, 2),
		mux:     iosync.NewLock(),
	}
}

func (v *Tmpl) KeysHandler(call func(keys ...string)) {
	v.keyH = call
}

func (v *Tmpl) SetKey(key, value string) {
	v.mux.Lock(func() {
		v.keys[key] = value
	})
	v.sendTrigger()
}

func (v *Tmpl) funcs() template.FuncMap {
	return template.FuncMap{
		"key": func(k string) (out string) {
			v.mux.RLock(func() {
				var ok bool
				out, ok = v.keys[k]
				if ok {
					return
				}
				out = ""
				v.keyH(k)
			})
			return
		},
	}
}

func (v *Tmpl) Add(in, out string) error {
	t, err := NewItem(in, out, v.funcs())
	if err != nil {
		return err
	}
	v.tmpls = append(v.tmpls, t)
	return nil
}

func (v *Tmpl) sendTrigger() {
	select {
	case v.trigger <- struct{}{}:
	default:
	}
}

func (v *Tmpl) Watch(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-v.trigger:
				for _, tmpl := range v.tmpls {
					if err := tmpl.Update(); err != nil {
						v.log.WithError("err", err).Errorf("update template")
					}
				}
			}
		}
	}()
	v.sendTrigger()
}
