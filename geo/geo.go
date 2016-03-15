package geo

import (
	"math"
	"math/cmplx"
)

// Intersection finds the intersection of two straight lines (a1x+b1y+c1, a2x+b2y+c2)
func Intersection(a1, b1, c1, a2, b2, c2 float64) (float64, float64) {
	return (b1*c2 - b2*c1) / (a1*b2 - a2*b1), (-a1*c2 + a2*c1) / (a1*b2 - a2*b1)
}

// Distance obtains the distance between the point and the line
func Distance(a, b, c float64, z complex128) float64 {
	return math.Abs(a*real(z)+b*imag(z)+c) / math.Sqrt(a*a+b*b)
}

// DistanceLS determine the distance between the point and the line segment
func DistanceLS(p, a, b complex128) float64 {
	ab, ba, pa, pb := a-b, b-a, p-a, p-b
	if Dot(ba, pa) < 0 {
		return cmplx.Abs(pa)
	}

	if Dot(ab, pb) < 0 {
		return cmplx.Abs(pb)
	}

	v, z := Cross(ba, 0, pa, 0)
	return math.Sqrt(real(v)*real(v)+imag(v)*imag(v)+z*z) / cmplx.Abs(ba)
}

// Dot is inner product
func Dot(a, b complex128) float64 {
	return real(a * cmplx.Conj(b))
}

// Cross is outer product
func Cross(a complex128, z1 float64, b complex128, z2 float64) (complex128, float64) {
	return complex(imag(a)*z2-z1*imag(b), z1*real(b)-real(a)*z2), real(a)*imag(b) - imag(a)*real(b)
}

// IsOrthogonal determine whether two line segments are perpendicular to each other
func IsOrthogonal(a, b, c, d complex128) bool {
	if real((a-b)/(c-d)) == 0 {
		return true
	}
	return false
}

// IsParallel tests whether parallel two line segments
func IsParallel(a, b, c, d complex128) bool {
	if imag((a-b)/(c-d)) == 0 {
		return true
	}
	return false
}

// IsIntersection determine whether two line segments intersect
func IsIntersection(a1, a2, b1, b2 complex128) bool {
	x1, x2, x3, x4 := real(a1), real(a2), real(b1), real(b2)
	y1, y2, y3, y4 := imag(a1), imag(a2), imag(b1), imag(b2)
	ta := (x3-x4)*(y1-y3) + (y3-y4)*(x3-x1)
	tb := (x3-x4)*(y2-y3) + (y3-y4)*(x3-x2)
	tc := (x1-x2)*(y3-y1) + (y1-y2)*(x1-x3)
	td := (x1-x2)*(y4-y1) + (y1-y2)*(x1-x4)
	return tc*td <= 0 && ta*tb <= 0
}

// Rotate rotate the origin center
func Rotate(p complex128, rad float64) complex128 {
	return p * cmplx.Rect(1, rad)
}

// EquationL is Equation of a line (ax+by+c)
func EquationL(a, b complex128) (float64, float64, float64) {
	return imag(b) - imag(a), real(a) - real(b), imag(a)*real(b) - real(a)*imag(b)
}

// Area2D to calculate the area of a polygon
func Area2D(ps ...complex128) float64 {
	sum, ps := float64(0), append(ps, ps[0])
	for i, v := range ps[:len(ps)-1] {
		x, y := real(v)-real(ps[i+1]), imag(v)+imag(ps[i+1])
		sum += x * y
	}
	return math.Abs(sum) / 2
}

// PosL is tests whether a point is on the right or left side of the line segment
func PosL(p, a, b complex128) int {
	ba, pa := b-a, p-a
	c := real(ba)*imag(pa) - imag(ba)*real(pa)
	switch {
	case c > 0:
		return 1 // left
	case c < 0:
		return -1 // right
	default:
		return 0 // on the line
	}
}

// InPolygon to the inside and outside judgment of a convex polygon
func InPolygon(p complex128, ps []complex128) bool {
	ps = append(ps, ps[0])
	sign := PosL(p, ps[0], ps[1])
	for i, v := range ps[1 : len(ps)-1] {
		if sign*PosL(p, v, ps[i+2]) < 0 {
			return false
		}
	}
	return true
}

// NearestP determmine the nearest point of the line segment from point
func NearestP(p, a, b complex128) complex128 {
	ab, ba, pa, pb := a-b, b-a, p-a, p-b
	if Dot(ba, pa) < 0 {
		return a
	}

	if Dot(ab, pb) < 0 {
		return b
	}
	return a + complex(Dot(pa, ba)/(real(ba)*real(ba)+imag(ba)*imag(ba)), 0)*ba
}

// SMCircle is smallest enclosing circle
// http://eomole.hatenablog.com/entry/20100219/1292052417
func SMCircle(ps []complex128) (complex128, float64) {
	p, index := ps[0], 0
	ratio, r := complex(0.5, 0), 0.0

	for c := 2 * math.Log2(1000000000); c >= 0; c-- {
		for j := 0; j < 100; j++ {
			max := 0.0
			for i, v := range ps {
				d := cmplx.Abs(v - p)
				if d > max {
					max, index = d, i
				}
			}
			p += (ps[index] - p) * ratio
			r = max
		}
		ratio /= 2
	}
	return p, r
}

// Rad converts from the degree to the radian
func Rad(d float64) float64 {
	return d * math.Pi / 180
}

// Deg converts from the radian to the degree
func Deg(r float64) float64 {
	return r * 180 / math.Pi
}

// Angle determine the angle(radian) of two vectors
func Angle(u, v complex128) float64 {
	return cmplx.Phase(u / v)
}
