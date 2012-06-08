// A go.flint example: Lambert W function power series.

// This is basically code from Fredrik Johansson's blog March 11, 2011
// http://fredrikj.net/blog/2011/03/a-flint-example-lambert-w-function-power-series/
// translated to Go and go.flint.

package main

import (
	"fmt"
	mp "github.com/frithjof-schulze/go.flint/fmpq"
)

// As today’s example, let us implement the Lambert W
// function for the power series ring Q[[x]]. The Lambert W
// function is defined implicitly by the equation
//   x = W(x)exp(W(z)),
// which can be solved using Newton iteration with the
// update step w = w-(w exp(w)-x)/((w+1) exp(w)).
//
// Power series Newton iteration is just like numerical
// Newton iteration, except that the convergence behavior is
// much simpler: starting with a correct first-order
// expansion, each iteration at least doubles the number of
// correct coefficients.
//
// This is a simple recursive implementation with
// asymptotically optimal performance (up to constant
// factors). Beyond the base case W(x) = 0 + O(x), the
// function just computes w to accuracy ceil(n/2), and then
// extends it to accuracy n using a single Newton step.
func lambertw(w, x *mp.Poly, n int64) {
	if n == 1 {
		w.SetInt64(0)
		return
	}

	lambertw(w, x, (n+1)/2)

	t := new(mp.Poly)
	t.ExpSeries(w, n)

	u := new(mp.Poly)
	v := new(mp.Poly)
	u.MulLow(t, w, n)
	v.Sub(u, x)
	t.Add(u, t)
	u.DivSeries(v, t, n)
	w.Sub(w, u)
}

func main() {
	x := mp.NewPoly(0)
	w := mp.NewPoly(0)
	x.SetCoeff64(1, 1)
	lambertw(w, x, 10)

	s := w.String()
	fmt.Printf("%v\n", s)
}

// The output of this program should be:
//   531441/4480*x^9 - 16384/315*x^8 + 16807/720*x^7 - 54/5*x^6 +
//   125/24*x^5 - 8/3*x^4 + 3/2*x^3 - 1*x^2 + 1*x
