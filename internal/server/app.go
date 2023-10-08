/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"go.arwos.org/loopy/internal/pkg/db"
	"go.osspkg.com/goppy/sdk/app"
)

type AppV1 struct {
	db *db.DB
}

func New(db *db.DB) *AppV1 {
	return &AppV1{
		db: db,
	}
}

func (v *AppV1) Up(ctx app.Context) error {
	return nil
}

func (v *AppV1) Down() error {
	return nil
}
