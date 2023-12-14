/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package tmpl

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"go.arwos.org/loopy/internal/pkg/cachekeys"
	"go.osspkg.com/goppy/iosync"
)

type (
	Item struct {
		src   string
		dst   string
		mt    int64
		tmpl  *template.Template
		state iosync.Switch
	}

	Action interface {
		Query(key string)
		Keys() cachekeys.MapGetter
	}

	State interface {
		IsPrepared() bool
	}
)

func NewItem(in, out string, a Action) (*Item, error) {
	fi, err := os.Lstat(in)
	if err != nil {
		return nil, fmt.Errorf("get file info %s: %w", in, err)
	}
	b, err := os.ReadFile(in)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", in, err)
	}
	obj := &Item{
		src:   in,
		dst:   out,
		mt:    fi.ModTime().Unix(),
		state: iosync.NewSwitch(),
	}
	obj.tmpl, err = template.New("_").Funcs(TemplateFuncMap(a, obj)).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", in, err)
	}
	return obj, nil
}

func (v *Item) IsChanged() bool {
	fi, err := os.Lstat(v.src)
	if err != nil {
		return false
	}
	return v.mt < fi.ModTime().Unix()
}

func (v *Item) IsPrepared() bool {
	return v.state.IsOn()
}

func (v *Item) Prepare() error {
	var buf bytes.Buffer
	err := v.tmpl.Execute(&buf, nil)
	if err != nil {
		return fmt.Errorf("execute %s: %w", v.src, err)
	}
	v.state.On()
	return nil
}

func (v *Item) Update() error {
	if v.state.IsOff() {
		return fmt.Errorf("template is not prepared: %s", v.src)
	}
	var buf bytes.Buffer
	err := v.tmpl.Execute(&buf, nil)
	if err != nil {
		return fmt.Errorf("execute %s: %w", v.src, err)
	}
	err = os.WriteFile(v.dst, buf.Bytes(), 0755)
	if err != nil {
		return fmt.Errorf("write %s: %w", v.dst, err)
	}
	return nil
}
