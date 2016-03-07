package geo

import "math"

// P is point
type P struct {
	x, y float64
}

// Intersection finds the intersection of two straight lines (a1x+b1y+c1, a2x+b2y+c2)
func Intersection(a1, b1, c1, a2, b2, c2 float64) (float64, float64) {
	return (b1*c2 - b2*c1) / (a1*b2 - a2*b1), (-a1*c2 + a2*c1) / (a1*b2 - a2*b1)
}

// Distance obtains the distance between the point and the line
func Distance(a, b, c, x, y float64) float64 {
	return math.Abs(a*x+b*y+c) / math.Sqrt(a*a+b*b)
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
	return Dot(a1.x-a2.x, a1.y-a2.y, 0, b1.x-b2.x, b1.y-b2.y, 0) == 0
}

// IsParallel tests whether parallel two line segments
func IsParallel(a1, a2, b1, b2 P) bool {
	x, y, z := Cross(a1.x-a2.x, a1.y-a2.y, 0, b1.x-b2.x, b1.y-b2.y, 0)
	return x == 0 && y == 0 && z == 0
}

// IsIntersection determine whether two line segments intersect
func IsIntersection(a1, a2, b1, b2 P) bool {
	ta := (b1.x-b2.x)*(a1.y-b1.y) + (b1.y-b2.y)*(b1.x-a1.x)
	tb := (b1.x-b2.x)*(a2.y-b1.y) + (b1.y-b2.y)*(b1.x-a2.x)
	tc := (a1.x-a2.x)*(b1.y-a1.y) + (a1.y-a2.y)*(a1.x-b1.x)
	td := (a1.x-a2.x)*(b2.y-a1.y) + (a1.y-a2.y)*(a1.x-b2.x)
	return tc*td < 0 && ta*tb < 0
}

// Rotate rotate the origin center
func Rotate(x, y, rad float64) (float64, float64) {
	return x*math.Cos(rad) - y*math.Sin(rad), x*math.Sin(rad) + y*math.Cos(rad)
}

// EquationL is Equation of a line (ax+by+c)
func EquationL(a, b P) (float64, float64, float64) {
	return b.y - a.y, a.x - b.x, a.y*b.x - a.x*b.y
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
