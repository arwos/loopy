SHELL=/bin/bash

.PHONY: install
install:
	go install github.com/osspkg/devtool@latest

.PHONY: setup
setup:
	devtool setup-lib

.PHONY: lint
lint:
	devtool lint

.PHONY: work
work:
	go work use -r .
	go work sync

.PHONY: license
license:
	devtool license

.PHONY: build
build:
	devtool build --arch=amd64

.PHONY: tests
tests:
	devtool test

.PHONY: pre-commite
pre-commite: setup lint build tests

.PHONY: ci
ci: install setup lint build tests

run_local_server:
	go run cmd/loop/main.go --config=config/config.dev.yaml

run_local_cli:
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv set "k1/k2/k3" "{\"data\":\"aaa\"}"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv get "k1/k2/k3"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv del "k1/k2/k3"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv get "k1/k2/k3"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv set "k1/aaa/bbb" "{\"data\":\"bbb\"}"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv set "k1/aaaa" "{\"data\":\"aaaa\"}"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv set "k1/eeee" "{\"data\":\"eeee\"}"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv set "k1/bbb/ccc/ddd" "{\"data\":\"ddd\"}"
#	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv search "k1/"
	go run cmd/loopcli/main.go --server=127.0.0.1:8080 kv list "k1/bbb/"

