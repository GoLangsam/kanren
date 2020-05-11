// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GoLangsam/do/cli/cancel"
)

// ===========================================================================

func main() {

	flagParse()

	ctx, _ := cancel.WithCancel()
	doneFn := cancel.Done(ctx)
	_ = doneFn

	if n > 0 {
		sample(ctx, n)
	} else {
		for i := 1; i <= max; i++ {
			sample(ctx, i)
		}
	}

	if x {
		fmt.Println("about to leave ...")
		<-time.After(time.Millisecond * 100)

		if ctx.Err() != nil {
			see("Early graceful exit!", tab, "reason:", tab, ctx.Err().Error())
		}

		see()

		os.Exit(1) // to see leaking goroutines, if any
	}
}

// ===========================================================================
