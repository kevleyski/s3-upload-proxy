// Copyright 2017 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type cacheControlRules []cacheControlRule

type cacheControlRule struct {
	Extension string `json:"ext"`
	MaxAge    uint   `json:"maxAge"`
}

func (c *cacheControlRules) Set(value string) error {
	err := json.Unmarshal([]byte(value), c)
	return err
}

func (c cacheControlRules) headerValue(file string) *string {
	var (
		found  bool
		maxAge uint
	)

	ext := filepath.Ext(file)
	for _, rule := range c {
		if rule.Extension == ext {
			found = true
			maxAge = rule.MaxAge
			break
		}
	}

	if !found {
		return nil
	}
	cacheControl := fmt.Sprintf("max-age=%d", maxAge)
	return &cacheControl
}
