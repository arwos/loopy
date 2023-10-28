module go.arwos.org/loopy

go 1.20

require (
	github.com/mailru/easyjson v0.7.7
	go.arwos.org/loopy/api v0.0.0-00010101000000-000000000000
	go.etcd.io/bbolt v1.3.8
	go.osspkg.com/goppy v0.14.1
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	go.osspkg.com/algorithms v1.3.0 // indirect
	go.osspkg.com/static v1.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arwos.org/loopy/api => ./api
