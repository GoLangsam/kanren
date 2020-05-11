// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/qqq/qqq.main.init.go"

package main

import "log"

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}
