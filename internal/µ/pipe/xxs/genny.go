// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file uses geanny to pull the type specific generic code

//go:generate -command genny genny -pkg $GOPACKAGE

//go:generate genny	-in ../../../../../pipe/.generate.xx/01-any-mode.go	-out stream.genny		gen "anyThing=S mode=supply anyMode=StreamOfStates"

package pipe
