package geo

import (
	"math"
	"math/cmplx"
)

// GravityP is the center of gravity of the triangle
func GravityP(ps []complex128) complex128 {
	return (ps[0] + ps[1] + ps[2]) / 3
}

// InnerC is inscried circle of the triangle
func InnerC(ps []complex128) (complex128, float64) {
	a, b, c := cmplx.Abs(ps[1]-ps[2]), cmplx.Abs(ps[2]-ps[0]), cmplx.Abs(ps[0]-ps[1])
	s := Area2D(ps)
	return (complex(a, 0)*ps[0] + complex(b, 0)*ps[1] + complex(c, 0)*ps[2]) / complex(a+b+c, 0), 2 * s / (a + b + c)
}

// CircumC is circumscribed circle of the triangle
func CircumC(ps []complex128) (complex128, float64) {
	ra, rb, rc := Angle(ps[1]-ps[0], ps[2]-ps[0]), Angle(ps[2]-ps[1], ps[0]-ps[1]), Angle(ps[0]-ps[2], ps[1]-ps[2])
	sa, sb, sc := math.Sin(2*ra), math.Sin(2*rb), math.Sin(2*rc)
	p := (ps[0]*complex(sa, 0) + ps[1]*complex(sb, 0) + ps[2]*complex(sc, 0)) / complex(sa+sb+sc, 0)
	return p, cmplx.Abs(ps[0] - p)
}

// OrthoP is orthocenter of the triangle
func OrthoP(ps []complex128) complex128 {
	ra, rb, rc := Angle(ps[1]-ps[0], ps[2]-ps[0]), Angle(ps[2]-ps[1], ps[0]-ps[1]), Angle(ps[0]-ps[2], ps[1]-ps[2])
	ta, tb, tc := math.Tan(ra), math.Tan(rb), math.Tan(rc)
	return (ps[0]*complex(ta, 0) + ps[1]*complex(tb, 0) + ps[2]*complex(tc, 0)) / complex(ta+tb+tc, 0)
}

// SignedArea is signed area of the triangle
func SignedArea(ps []complex128) float64 {
	return imag(cmplx.Conj(ps[1]-ps[0])*(ps[2]-ps[0])) * 0.5
}
