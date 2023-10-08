/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package db

import "go.osspkg.com/goppy/plugins"

var Plugin = plugins.Plugin{
	Config: &Config{},
	Inject: New,
}
