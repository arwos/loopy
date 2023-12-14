/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package utils

import "bytes"

func BucketsFromKey(v []byte, sep []byte) [][]byte {
	return bytes.Split(v, sep)
}

func BucketKeyPath(v [][]byte, sep []byte) []byte {
	return append(bytes.Join(v, sep), sep...)
}

func BytesExtend(b ...[]byte) []byte {
	count := 0
	for _, item := range b {
		count += len(item)
	}
	result := make([]byte, 0, count)
	for _, item := range b {
		result = append(result, item...)
	}
	return result
}
