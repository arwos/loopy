/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package tmpl

import (
	"strings"
	"text/template"
)

func TemplateFuncMap(a Action, s State) template.FuncMap {
	return template.FuncMap{
		"key":            KeyFunc(a, s),
		"key_or_default": KeyOrDefaultFunc(a, s),
		"tree":           TreeFunc(a, s),
	}
}

func KeyFunc(a Action, s State) func(string) string {
	return func(key string) string {
		if !s.IsPrepared() {
			a.Query(key)
			return ""
		}
		if value := a.Keys().Get(key); value != nil {
			return *value
		}
		return ""
	}
}

func KeyOrDefaultFunc(a Action, s State) func(string, string) string {
	return func(key string, def string) string {
		if !s.IsPrepared() {
			a.Query(key)
			return def
		}
		if value := a.Keys().Get(key); value != nil {
			return *value
		}
		return def
	}
}

type TreeModel struct {
	Key   string
	Value string
}

func TreeFunc(a Action, s State) func(string) []TreeModel {
	return func(prefix string) []TreeModel {
		data := make([]TreeModel, 0, 10)
		if !s.IsPrepared() {
			a.Query(prefix)
			return data
		}
		a.Keys().Each(func(key string, value *string) {
			if strings.HasPrefix(key, prefix) && value != nil {
				data = append(data, TreeModel{
					Key:   key,
					Value: *value,
				})
			}
		})
		return data
	}
}
