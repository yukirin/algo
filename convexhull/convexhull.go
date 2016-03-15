package convexhull

import "math/cmplx"

// Solve is Gift wrapping algorithm
func Solve(ps []complex128) []complex128 {
	var chList []complex128
	a := nearest(ps)

	for {
		chList = append(chList, a)
		b := ps[0]

		for _, c := range ps[1:] {
			if a == b {
				b = c
				continue
			}

			p1, p2 := b-a, c-a
			v := product(p2, p1)
			if v > 0 || (v == 0 && cmplx.Abs(p2) > cmplx.Abs(p1)) {
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

func nearest(ps []complex128) complex128 {
	min := ps[0]
	for _, p := range ps {
		if imag(p) < imag(min) {
			min = p
		}
	}

	nearY := []complex128{min}
	for _, p := range ps {
		if imag(min) == imag(p) {
			nearY = append(nearY, p)
		}
	}

	min = nearY[0]
	for _, p := range nearY {
		if real(p) < real(min) {
			min = p
		}
	}
	return min
}

func product(a, b complex128) float64 {
	return real(a)*imag(b) - real(b)*imag(a)
}
