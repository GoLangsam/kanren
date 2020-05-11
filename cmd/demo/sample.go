// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	. "github.com/GoLangsam/kanren"
	"github.com/GoLangsam/sexpr"
)

// ===========================================================================

func sample(ctx context.Context, n int) {

	switch n {

	case 1:
		zebra(ctx, 0)
	case 2:
		zebra(ctx, 1)
	case 3:
		zebra(ctx, 1) // 2) // TODO: mode 2 runs forever
	case 4:
		e := NewS()
		x := e.Fresh("x")

		e1 := Equal(
			sexpr.NewSymbol("olive"),
			x,
		)
		e2 := Equal(
			sexpr.NewSymbol("oil"),
			x,
		)
		e1 = FAIL
		or := Disjoint(e1, e2)

		fmt.Println("#", n, " Oliv: ", e1(e.Clone()).String())
		fmt.Println("#", n, " Oil : ", e2(e.Clone()).String())
		fmt.Println("#", n, " .OR.: ", or(e.Clone()).String())

	default:
		fmt.Println("No such sample #", n, " - max =", max)

	}

}

const max = 4 // # of samples

// ===========================================================================
