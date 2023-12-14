/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package utils

import (
	"reflect"
	"testing"
)

func TestBucketKeyPath(t *testing.T) {
	type args struct {
		v   [][]byte
		sep []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Case1",
			args: args{
				v:   append(make([][]byte, 0), []byte("k1/k2"), []byte("k3/k4")),
				sep: []byte("/"),
			},
			want: []byte("k1/k2/k3/k4/"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketKeyPath(tt.args.v, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketKeyPath() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestBucketsFromKey(t *testing.T) {
	type args struct {
		v   []byte
		sep []byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Case1",
			args: args{
				v:   []byte("k1/k2/k3/k4/"),
				sep: []byte("/"),
			},
			want: append(make([][]byte, 0), []byte("k1"), []byte("k2"), []byte("k3"), []byte("k4"), []byte("")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketsFromKey(tt.args.v, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketsFromKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesExtend(t *testing.T) {
	type args struct {
		b [][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Case1",
			args: args{
				b: append(make([][]byte, 0), []byte("k1"), []byte("k2"), []byte("k3"), []byte("k4"), []byte("")),
			},
			want: []byte("k1k2k3k4"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesExtend(tt.args.b...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesExtend() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
