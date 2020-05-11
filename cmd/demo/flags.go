// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
)

// ===========================================================================

var ( //flags
	n int
	x bool
)

func init() {
	flag.IntVar(&n, "n", 0, "# of sample: 0 = all")
	flag.BoolVar(&x, "x", true, "use os.Exit(1) to see leaking goroutines, if any")
	flag.Parse()
}

// ===========================================================================

func flagParse() {
	flag.Parse()
}

// ===========================================================================
