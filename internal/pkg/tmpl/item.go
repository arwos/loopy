package tmpl

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type Item struct {
	src  string
	dst  string
	tmpl *template.Template
}

func NewItem(in, out string, funcs template.FuncMap) (*Item, error) {
	b, err := os.ReadFile(in)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", in, err)
	}
	t, err := template.New("_").Funcs(funcs).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", in, err)
	}
	return &Item{
		src:  in,
		dst:  out,
		tmpl: t,
	}, nil
}

func (v *Item) Update() error {
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
