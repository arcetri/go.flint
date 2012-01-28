// Copyright 2011 go.flint authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package nmod

// #cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <ulong_extras.h>
// #include <nmod_mat.h>
import "C"

import (
	// "os"
	// "unsafe"
)

// An IntMat represents a matrix with integral entries. The
// zero value for a Mat represents the zero matrix.
type NmodMat struct {
	i    C.nmod_mat_t
	init bool
}

