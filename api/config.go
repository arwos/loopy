/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package api

import (
	"fmt"
	"net"
	"regexp"
)

type Config struct {
	SSL       bool
	HostPort  string
	AuthToken string
}

var domainRegexp = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+){0,}$`)
var portRegexp = regexp.MustCompile(`^(?i)[1-9][0-9-]+$`)

func (v Config) Validate() error {
	h, p, err := net.SplitHostPort(v.HostPort)
	if err != nil {
		return err
	}
	if net.ParseIP(h) == nil && !domainRegexp.MatchString(h) {
		return fmt.Errorf("invalid host: %s", h)
	}
	if !portRegexp.MatchString(p) {
		return fmt.Errorf("invalid port: %s", p)
	}
	return nil
}
