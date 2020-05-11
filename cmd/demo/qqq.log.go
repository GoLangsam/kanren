// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/do/qqq/qqq.log.go"

package main

import "log"

const (
	tab = "\t"
	eol = "\n"
)

var (
	see = log.Println
	sef = log.Printf

	die = log.Panicln
	dif = log.Panicf

	verbose bool
)

func qqq(args ...interface{}) {
	if verbose {
		see(args...)
	}
}

func qqf(format string, args ...interface{}) {
	if verbose {
		sef(format, args...)
	}
}
