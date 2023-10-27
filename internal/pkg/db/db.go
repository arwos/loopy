/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package db

import (
	"os"
	"time"

	"go.etcd.io/bbolt"
	"go.osspkg.com/goppy/sdk/app"
	"go.osspkg.com/goppy/sdk/iofile"
)

const databaseName = "loop.db"

type DB struct {
	db   *bbolt.DB
	conf DatabaseItem
}

func New(c *Config) *DB {
	return &DB{
		db:   nil,
		conf: c.Database,
	}
}

func (v *DB) Up(_ app.Context) error {
	if !iofile.Exist(v.conf.Folder) {
		if err := os.MkdirAll(v.conf.Folder, 0744); err != nil {
			return err
		}
	}
	filename := v.conf.Folder + "/" + databaseName
	db, err := bbolt.Open(filename, 0600, &bbolt.Options{
		Timeout:      3 * time.Second,
		NoGrowSync:   false,
		FreelistType: bbolt.FreelistArrayType,
	})
	if err != nil {
		return err
	}
	v.db = db
	return nil
}

func (v *DB) Down() error {
	if v.db != nil {
		return v.db.Close()
	}
	return nil
}
