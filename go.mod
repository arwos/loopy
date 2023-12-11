module go.arwos.org/loopy

go 1.20

require (
	github.com/mailru/easyjson v0.7.7
	go.arwos.org/loopy/api v0.0.0-00010101000000-000000000000
	go.etcd.io/bbolt v1.3.8
	go.osspkg.com/goppy v0.15.5
	go.osspkg.com/goppy/plugins v0.1.0
)

require (
	github.com/josharian/intern v1.0.0 // indirect
	go.osspkg.com/algorithms v1.3.0 // indirect
	go.osspkg.com/goppy/app v0.1.3 // indirect
	go.osspkg.com/goppy/console v0.1.0 // indirect
	go.osspkg.com/goppy/errors v0.1.0 // indirect
	go.osspkg.com/goppy/iosync v0.1.2 // indirect
	go.osspkg.com/goppy/syscall v0.1.0 // indirect
	go.osspkg.com/goppy/xc v0.1.0 // indirect
	go.osspkg.com/goppy/xlog v0.1.3 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.arwos.org/loopy/api => ./api
