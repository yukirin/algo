package convexhull

import "math"

// Point represents the plane of the coordinate
type Point struct {
	X, Y float64
}

// Solve is Gift wrapping algorithm
func Solve(ps []Point) []Point {
	var chList []Point
	a := nearest(ps)

	for {
		chList = append(chList, a)
		b := ps[0]

		for _, c := range ps[1:] {
			if a == b {
				b = c
				continue
			}

			p1 := &Point{b.X - a.X, b.Y - a.Y}
			p2 := &Point{c.X - a.X, c.Y - a.Y}
			v := product(p2, p1)
			if v > 0 || (v == 0 && math.Hypot(p2.X, p2.Y) > math.Hypot(p1.X, p1.Y)) {
				b = c
			}
		}

		a = b
		if a == chList[0] {
			break
		}
	}
	return chList
}

func nearest(ps []Point) Point {
	min := ps[0]
	for _, p := range ps {
		if p.Y < min.Y {
			min = p
		}
	}

	nearY := []Point{min}
	for _, p := range ps {
		if min.Y == p.Y {
			nearY = append(nearY, p)
		}
	}

	min = nearY[0]
	for _, p := range nearY {
		if p.X < min.X {
			min = p
		}
	}
	return min
}

func product(a, b *Point) float64 {
	return a.X*b.Y - b.X*a.Y
}
