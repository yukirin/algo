package geo

import "math"

// P is point
type P struct {
	X, Y float64
}

// Intersection finds the intersection of two straight lines (a1x+b1y+c1, a2x+b2y+c2)
func Intersection(a1, b1, c1, a2, b2, c2 float64) (float64, float64) {
	return (b1*c2 - b2*c1) / (a1*b2 - a2*b1), (-a1*c2 + a2*c1) / (a1*b2 - a2*b1)
}

// Distance obtains the distance between the point and the line
func Distance(a, b, c, x, y float64) float64 {
	return math.Abs(a*x+b*y+c) / math.Sqrt(a*a+b*b)
}

// DistanceLS determine the distance between the point and the line segment
func DistanceLS(p P, a, b P) float64 {
	if Dot(b.X-a.X, b.Y-a.Y, 0, p.X-a.X, p.Y-a.Y, 0) < 0 {
		return math.Hypot(p.X-a.X, p.Y-a.Y)
	}

	if Dot(a.X-b.X, a.Y-b.Y, 0, p.X-b.X, p.Y-b.Y, 0) < 0 {
		return math.Hypot(p.X-b.X, p.Y-b.Y)
	}

	x, y, z := Cross(b.X-a.X, b.Y-a.Y, 0, p.X-a.X, p.Y-a.Y, 0)
	return math.Sqrt(x*x+y*y+z*z) / math.Hypot(b.X-a.X, b.Y-a.Y)
}

// Dot is inner product
func Dot(x1, y1, z1, x2, y2, z2 float64) float64 {
	return x1*x2 + y1*y2 + z1*z2
}

// Cross is outer product
func Cross(x1, y1, z1, x2, y2, z2 float64) (float64, float64, float64) {
	return y1*z2 - z1*y2, z1*x2 - x1*z2, x1*y2 - y1*x2
}

// IsOrthogonal determine whether two line segments are perpendicular to each other
func IsOrthogonal(a1, a2, b1, b2 P) bool {
	return Dot(a1.X-a2.X, a1.Y-a2.Y, 0, b1.X-b2.X, b1.Y-b2.Y, 0) == 0
}

// IsParallel tests whether parallel two line segments
func IsParallel(a1, a2, b1, b2 P) bool {
	x, y, z := Cross(a1.X-a2.X, a1.Y-a2.Y, 0, b1.X-b2.X, b1.Y-b2.Y, 0)
	return x == 0 && y == 0 && z == 0
}

// IsIntersection determine whether two line segments intersect
func IsIntersection(a1, a2, b1, b2 P) bool {
	ta := (b1.X-b2.X)*(a1.Y-b1.Y) + (b1.Y-b2.Y)*(b1.X-a1.X)
	tb := (b1.X-b2.X)*(a2.Y-b1.Y) + (b1.Y-b2.Y)*(b1.X-a2.X)
	tc := (a1.X-a2.X)*(b1.Y-a1.Y) + (a1.Y-a2.Y)*(a1.X-b1.X)
	td := (a1.X-a2.X)*(b2.Y-a1.Y) + (a1.Y-a2.Y)*(a1.X-b2.X)
	return tc*td < 0 && ta*tb < 0
}

// Rotate rotate the origin center
func Rotate(x, y, rad float64) (float64, float64) {
	return x*math.Cos(rad) - y*math.Sin(rad), x*math.Sin(rad) + y*math.Cos(rad)
}

// EquationL is Equation of a line (ax+by+c)
func EquationL(a, b P) (float64, float64, float64) {
	return b.Y - a.Y, a.X - b.X, a.Y*b.X - a.X*b.Y
}

// Area2D to calculate the area of a polygon
func Area2D(ps ...[]float64) float64 {
	sum := float64(0)
	ps = append(ps, []float64{ps[0][0], ps[0][1]})

	for i, v := range ps[:len(ps)-1] {
		x, y := v[0]-ps[i+1][0], v[1]+ps[i+1][1]
		sum += x * y
	}
	return math.Abs(sum) / 2
}

// PosL is tests whether a point is on the right or left side of the line segment
func PosL(p P, a, b P) int {
	x1, y1 := b.X-a.X, b.Y-a.Y
	x2, y2 := p.X-a.X, p.Y-a.Y
	c := x1*y2 - y1*x2
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
func InPolygon(p P, ps []P) bool {
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
func NearestP(p P, a, b P) P {
	if Dot(b.X-a.X, b.Y-a.Y, 0, p.X-a.X, p.Y-a.Y, 0) < 0 {
		return a
	}

	if Dot(a.X-b.X, a.Y-b.Y, 0, p.X-b.X, p.Y-b.Y, 0) < 0 {
		return b
	}
	ba := P{b.X - a.X, b.Y - a.Y}
	s := Dot(p.X-a.X, p.Y-a.Y, 0, ba.X, ba.Y, 0) / (ba.X*ba.X + ba.Y*ba.Y)
	return P{a.X + s*ba.X, a.Y + s*ba.Y}
}

// SMCircle is smallest enclosing circle
// http://eomole.hatenablog.com/entry/20100219/1292052417
func SMCircle(ps []P) (P, float64) {
	p, index := ps[0], 0
	ratio, r := 0.5, 0.0

	for c := 2 * math.Log2(1000000000); c >= 0; c-- {
		for j := 0; j < 100; j++ {
			max := 0.0
			for i, v := range ps {
				d := math.Hypot(v.X-p.X, v.Y-p.Y)
				if d > max {
					max, index = d, i
				}
			}
			p.X += (ps[index].X - p.X) * ratio
			p.Y += (ps[index].Y - p.Y) * ratio
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
func Angle(u, v P) float64 {
	l := math.Hypot(u.X, u.Y) * math.Hypot(v.X, v.Y)
	cos := Dot(u.X, u.Y, 0, v.X, v.Y, 0)
	return math.Acos(cos / l)
}
