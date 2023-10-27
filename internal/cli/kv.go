/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package cli

import (
	"context"
	"errors"
	"time"

	"go.arwos.org/loopy/api"
	"go.osspkg.com/goppy/sdk/console"
)

func CommandKVCommon() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("kv", "Work with Loop KV")

		setter.AddCommand(
			CommandKVSet(),
			CommandKVGet(),
			CommandKVDel(),
			CommandKVSearch(),
			CommandKVList(),
		)
		setter.ExecFunc(func(_ []string) {
			console.Errorf("Use sub command. See help with flag --help")
		})
	})
}

func CommandKVSet() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("set", "Set data to Loop KV")
		setter.Flag(func(flagsSetter console.FlagsSetter) {
			flagsSetter.StringVar("server", "127.0.0.1:8080", "Set LoopServer address")
			flagsSetter.Bool("ssl", "Set use ssl for LoopServer address")
		})
		setter.ArgumentFunc(func(s []string) ([]string, error) {
			if len(s) != 2 {
				return nil, errors.New("You need to pass 2 arguments: key and value")
			}
			return s, nil
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			ctx, cncl := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cncl()
			apiCli, err := api.NewKV(&api.Config{HostPort: server, SSL: ssl})
			console.FatalIfErr(err, "kv init client")
			res, err := apiCli.Set(ctx, args[0], args[1])
			console.FatalIfErr(err, "kv set data")
			console.Rawf(res)
		})
	})
}

func CommandKVGet() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("get", "Get data from Loop KV")
		setter.Flag(func(flagsSetter console.FlagsSetter) {
			flagsSetter.StringVar("server", "127.0.0.1:8080", "Set LoopServer address")
			flagsSetter.Bool("ssl", "Set use ssl for LoopServer address")
		})
		setter.ArgumentFunc(func(s []string) ([]string, error) {
			if len(s) != 1 {
				return nil, errors.New("You need to pass 1 argument: key")
			}
			return s, nil
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			ctx, cncl := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cncl()
			apiCli, err := api.NewKV(&api.Config{HostPort: server, SSL: ssl})
			console.FatalIfErr(err, "kv init client")
			res, err := apiCli.Get(ctx, args[0])
			console.FatalIfErr(err, "kv get data")
			console.Rawf(res)
		})
	})
}

func CommandKVDel() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("del", "Delete data from Loop KV")
		setter.Flag(func(flagsSetter console.FlagsSetter) {
			flagsSetter.StringVar("server", "127.0.0.1:8080", "Set LoopServer address")
			flagsSetter.Bool("ssl", "Set use ssl for LoopServer address")
		})
		setter.ArgumentFunc(func(s []string) ([]string, error) {
			if len(s) != 1 {
				return nil, errors.New("You need to pass 1 argument: key")
			}
			return s, nil
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			ctx, cncl := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cncl()
			apiCli, err := api.NewKV(&api.Config{HostPort: server, SSL: ssl})
			console.FatalIfErr(err, "kv init client")
			res, err := apiCli.Delete(ctx, args[0])
			console.FatalIfErr(err, "kv delete data")
			console.Rawf(res)
		})
	})
}

func CommandKVSearch() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("search", "Search data in Loop KV")
		setter.Flag(func(flagsSetter console.FlagsSetter) {
			flagsSetter.StringVar("server", "127.0.0.1:8080", "Set LoopServer address")
			flagsSetter.Bool("ssl", "Set use ssl for LoopServer address")
		})
		setter.ArgumentFunc(func(s []string) ([]string, error) {
			if len(s) != 1 {
				return nil, errors.New("You need to pass 1 argument: key")
			}
			return s, nil
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			ctx, cncl := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cncl()
			apiCli, err := api.NewKV(&api.Config{HostPort: server, SSL: ssl})
			console.FatalIfErr(err, "kv init client")
			res, err := apiCli.Search(ctx, args[0])
			console.FatalIfErr(err, "kv search data")
			for _, re := range res {
				console.Rawf(">> %s", re.Key)
				console.Rawf(re.Value.String())
			}
		})
	})
}

func CommandKVList() console.CommandGetter {
	return console.NewCommand(func(setter console.CommandSetter) {
		setter.Setup("list", "List data in Loop KV")
		setter.Flag(func(flagsSetter console.FlagsSetter) {
			flagsSetter.StringVar("server", "127.0.0.1:8080", "Set LoopServer address")
			flagsSetter.Bool("ssl", "Set use ssl for LoopServer address")
		})
		setter.ArgumentFunc(func(s []string) ([]string, error) {
			if len(s) != 1 {
				return nil, errors.New("You need to pass 1 argument: key")
			}
			return s, nil
		})
		setter.ExecFunc(func(args []string, server string, ssl bool) {
			ctx, cncl := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cncl()
			apiCli, err := api.NewKV(&api.Config{HostPort: server, SSL: ssl})
			console.FatalIfErr(err, "kv init client")
			res, err := apiCli.List(ctx, args[0])
			console.FatalIfErr(err, "kv list data")
			for _, re := range res {
				console.Rawf(">> %s", re.Key)
				console.Rawf(re.Value.String())
			}
		})
	})
}
