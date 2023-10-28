/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package cli

import (
	"context"
	"strings"

	"go.arwos.org/loopy/api"
	"go.arwos.org/loopy/internal/pkg/tmpl"
	"go.osspkg.com/goppy/sdk/console"
	"go.osspkg.com/goppy/sdk/iosync"
	"go.osspkg.com/goppy/sdk/log"
	"go.osspkg.com/goppy/sdk/syscall"
)

func CommandTemplate() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("template", "Update template with Loop KV")
		setter.Flag(func(fs console.FlagsSetter) {
			fs.StringVar("server", "127.0.0.1:8080", "Set Loopy address")
			fs.Bool("ssl", "Set use ssl for Loopy address")
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			logger := log.Default()
			wg := iosync.NewGroup()
			ctx, cncl := context.WithCancel(context.TODO())
			cli, err := api.NewWatch(ctx, &api.Config{SSL: ssl, HostPort: server}, logger)
			console.FatalIfErr(err, "Connect to server %s", server)
			t := tmpl.New(logger)
			cli.KeyHandler(func(e api.EntitiesKV) {
				for _, kv := range e {
					t.SetKey(kv.Key, string(kv.Value))
				}
			})
			wg.Background(func() {
				if err := cli.Open(); err != nil {
					console.Errorf("Open connect to Loopy: %v", err.Error())
				}
				cncl()
			})
			t.KeysHandler(func(keys ...string) {
				cli.KeySubscribe(keys...)
			})
			for _, arg := range args {
				f := strings.Split(arg, ":")
				if len(f) != 2 {
					console.Fatalf("invalid file path [%s], use format: src:dst", arg)
				}
				console.FatalIfErr(t.Add(f[0], f[1]), "Add file path template")
			}
			t.Watch(ctx)
			go syscall.OnStop(cncl)
			<-ctx.Done()
			wg.Wait()
		})
	})
}
