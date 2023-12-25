# Loopy

Loopy is designed to store and manage keys and values that are used by various services and applications. It provides a reliable system for tracking changes and updating text templates in real time.

## Install with go

```shell
go install go.arwos.org/loopy/...@latest
```

## Server

Run server:

```shell
loppy --config=./config.yaml
```

## Console client

### Work with template

```shell
loopcli --server=127.0.0.1:9500 template \
    test_data/template.tmpl:test_data/template.out \
    test_data/template2.tmpl:test_data/template2.out
```

#### Template functions:

* __key__ - returns the key value or an empty string
```json
{{key "key/name"}}
```
* __key_or_default__ - returns the key value or default value
```json
{{key_or_default "key/name" "default value"}}
```
* __tree__ - returns a list of keys and value by prefix
```json
{{range $index, $data := tree "key/"}}
index: {{$index}} key: {{$data.Key}} val: {{$data.Value}}
{{end}}
```

### Work with keys

* Set key
```shell
loopcli --server=127.0.0.1:9500 kv set "key/name" "key_value"
```
* Get key
```shell
loopcli --server=127.0.0.1:9500 kv get "key/name"
```
* Delete key
```shell
loopcli --server=127.0.0.1:9500 kv del "key/name"
```
* List key by prefix (with empty value)
```shell
loopcli --server=127.0.0.1:9500 kv list "key/"
```
* Search key by prefix (without empty value)
```shell
loopcli --server=127.0.0.1:9500 kv search "key/"
```
