// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/GoLangsam/kanren/internal/µ/stat"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing = *stat.E
