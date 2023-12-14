/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package cli

import (
	"os"
	"strings"
	"time"

	"go.arwos.org/loopy/api"
	"go.arwos.org/loopy/internal/pkg/tmpl"
	"go.osspkg.com/goppy/console"
	"go.osspkg.com/goppy/iosync"
	"go.osspkg.com/goppy/syscall"
	"go.osspkg.com/goppy/xc"
	"go.osspkg.com/goppy/xlog"
)

func CommandTemplate() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("template", "Update template with Loop")
		setter.Flag(func(fs console.FlagsSetter) {
			fs.StringVar("server", "127.0.0.1:8080", "Set Loopy address")
			fs.StringVar("pid", "/tmp/loopy_template.pid", "Set Loopy PID file")
			fs.IntVar("uptime", 5, "Template update check interval")
			fs.Bool("ssl", "Set to use SSL for Loopy address")
		})
		setter.ExecFunc(func(args []string, server, pidfile string, uptime int64, ssl bool) {
			console.FatalIfErr(syscall.Pid(pidfile), "Write PID file %s", pidfile)

			logger := xlog.Default()
			logger.SetOutput(os.Stdout)
			logger.SetLevel(xlog.LevelDebug)
			logger.SetFormatter(xlog.NewFormatString())

			wg := iosync.NewGroup()
			closeCtx := xc.New()
			openCtx := xc.New()

			cli, err := api.NewWatch(closeCtx.Context(), &api.Config{SSL: ssl, HostPort: server}, logger)
			console.FatalIfErr(err, "Connect to server %s", server)
			cli.AfterOpened(openCtx.Close)
			cli.AfterClosed(closeCtx.Close)

			t := tmpl.New(logger)
			cli.KeyHandler(func(e api.EntitiesKV) {
				for _, kv := range e {
					t.SetKey(kv.Key, kv.Val)
				}
			})
			wg.Background(func() {
				if err0 := cli.Open(); err0 != nil {
					console.Errorf("Open connect to Loopy: %v", err0.Error())
				}
			})
			<-openCtx.Done()
			t.KeysHandler(func(key string) {
				cli.KeySubscribe(key)
			})

			for _, arg := range args {
				f := strings.Split(arg, ":")
				if len(f) != 2 {
					console.Fatalf("invalid file path [%s], use format: src:dst", arg)
				}
				console.FatalIfErr(t.Add(f[0], f[1]), "Add file path template")
			}
			go syscall.OnStop(closeCtx.Close)
			t.Watch(closeCtx, time.Duration(uptime)*time.Second)
			cli.Close()
			wg.Wait()
		})
	})
}
