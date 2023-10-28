/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"go.arwos.org/loopy/internal/cli"
	"go.osspkg.com/goppy/sdk/console"
)

func main() {

	app := console.New("loopcli", "LoopClient")
	app.RootCommand(console.NewCommand(func(setter console.CommandSetter) {
		setter.ExecFunc(func(_ []string) {
			console.Rawf("Hello! I am LoopClient ;-)")
		})
	}))
	app.AddCommand(
		cli.CommandKVCommon(),
		cli.CommandTemplate(),
	)
	app.Exec()
}
