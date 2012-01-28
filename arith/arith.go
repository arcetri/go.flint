// Copyright 2012 flint.go authors. All rights reserved.
// Use of this source code is governed by the GNU General
// Public License version 2 (or any later version).

package arith

// #cgo LDFLAGS: -lflint -lmpir -lmpfr -lm
// #include <stdlib.h>
// #include <flint.h>
// #include <fmpz.h>
// #include <arith.h>
import "C"

import "github.com/frithjof-schulze/flint.go/fmpz"

// RamanujanTau sets z to the Ramanujan tau function of n,
// which is the coefficient of q^n in the series expansion
// of the discriminant modular form.
func RamanujanTau(z, n *fmpz.Int) *fmpz.Int {
	C.fmpz_ramanujan_tau((*C.fmpz)(z),(*C.fmpz)(n))
	return z
}
