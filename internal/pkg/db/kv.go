/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package db

import (
	"bytes"
	"fmt"

	"go.arwos.org/loopy/internal/pkg/utils"
	"go.etcd.io/bbolt"
)

var (
	kvBucket = []byte("kv")
	kvSep    = []byte("/")
)

type EntityKV struct {
	Key   []byte
	Value []byte
}

func (v EntityKV) Buckets() [][]byte {
	return utils.BucketsFromKey(v.Key, kvSep)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func eachKV(buk *bbolt.Bucket, bukName []byte, bukPath []byte) []EntityKV {
	nbuk := buk.Bucket(bukName)
	if nbuk == nil {
		return nil
	}
	bukPath = utils.BytesExtend(bukPath, bukName, kvSep)
	result := make([]EntityKV, 0, 10)
	cur := nbuk.Cursor()
	for k, v := cur.First(); k != nil; k, v = cur.Next() {
		if v == nil {
			result = append(result, eachKV(nbuk, k, bukPath)...)
			continue
		}
		result = append(result, EntityKV{
			Key:   utils.BytesExtend(bukPath, k),
			Value: v,
		})
	}
	return result
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *DB) SetKV(item EntityKV) error {
	return v.db.Update(func(tx *bbolt.Tx) error {
		buk, err := tx.CreateBucketIfNotExists(kvBucket)
		if err != nil {
			return fmt.Errorf("create kv bucket: %w", err)
		}
		buks := item.Buckets()
		key := buks[len(buks)-1]
		for _, bukName := range buks[:len(buks)-1] {
			buk, err = buk.CreateBucketIfNotExists(bukName)
			if err != nil {
				return fmt.Errorf("create %s bucket: %w", string(bukName), err)
			}
		}
		return buk.Put(key, item.Value)
	})
}

func (v *DB) GetKV(item *EntityKV) error {
	return v.db.View(func(tx *bbolt.Tx) error {
		buk := tx.Bucket(kvBucket)
		if buk == nil {
			return fmt.Errorf("kv bucket not found")
		}
		buks := item.Buckets()
		key := buks[len(buks)-1]
		for _, bukName := range buks[:len(buks)-1] {
			buk = buk.Bucket(bukName)
			if buk == nil {
				return fmt.Errorf("key not found")
			}
		}
		item.Value = buk.Get(key)
		return nil
	})
}

func (v *DB) SearchKV(prefix []byte) ([]EntityKV, error) {
	result := make([]EntityKV, 0, 10)
	err := v.db.View(func(tx *bbolt.Tx) error {
		buk := tx.Bucket(kvBucket)
		if buk == nil {
			return fmt.Errorf("kv bucket not found")
		}
		buks := utils.BucketsFromKey(prefix, kvSep)
		bukPath := utils.BucketKeyPath(buks[:len(buks)-1], kvSep)
		prefix = buks[len(buks)-1]
		for _, bukName := range buks[:len(buks)-1] {
			buk = buk.Bucket(bukName)
			if buk == nil {
				return fmt.Errorf("key not found")
			}
		}
		cur := buk.Cursor()
		for k, v := cur.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cur.Next() {
			if v == nil {
				result = append(result, eachKV(buk, k, bukPath)...)
				continue
			}
			result = append(result, EntityKV{
				Key:   utils.BytesExtend(bukPath, k),
				Value: v,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *DB) ListKV(prefix []byte) ([]EntityKV, error) {
	prefix = bytes.TrimRight(prefix, string(kvSep))
	result := make([]EntityKV, 0, 10)
	err := v.db.View(func(tx *bbolt.Tx) error {
		buk := tx.Bucket(kvBucket)
		if buk == nil {
			return fmt.Errorf("kv bucket not found")
		}
		buks := utils.BucketsFromKey(prefix, kvSep)
		bukPath := utils.BucketKeyPath(buks, kvSep)
		for _, bukName := range buks {
			buk = buk.Bucket(bukName)
			if buk == nil {
				return nil
			}
		}
		cur := buk.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			if v == nil {
				result = append(result, EntityKV{
					Key:   utils.BytesExtend(bukPath, k, kvSep),
					Value: nil,
				})
				continue
			}
			result = append(result, EntityKV{
				Key:   utils.BytesExtend(bukPath, k),
				Value: v,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *DB) DelKV(item EntityKV) error {
	return v.db.Update(func(tx *bbolt.Tx) error {
		buk, err := tx.CreateBucketIfNotExists(kvBucket)
		if err != nil {
			return fmt.Errorf("create kv bucket: %w", err)
		}
		buks := item.Buckets()
		key := buks[len(buks)-1]
		for _, bukName := range buks[:len(buks)-1] {
			buk = buk.Bucket(bukName)
			if buk == nil {
				return nil
			}
		}

		if err = buk.Delete(key); err != nil {
			return err
		}

		return nil
	})
}
