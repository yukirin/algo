package convexhull

import "github.com/yukirin/algo/geo"

// Solve2 is Quickhull
func Solve2(ps []geo.P) []geo.P {
	lp, rp := leftMost(ps), rightMost(ps)
	ret := make([]geo.P, 2, 500)
	ret[0], ret[1] = lp, rp
	return append(ret, solve(lp, rp, ps)...)
}

func solve(lp, rp geo.P, ps []geo.P) []geo.P {
	if len(ps) == 0 {
		return nil
	}

	a, b, c := geo.EquationL(lp, rp)
	updown, div := [][]geo.P{make([]geo.P, 0, 50), make([]geo.P, 0, 50)}, make([]geo.P, 2)
	maxU, maxD := float64(0), float64(0)
	ret := make([]geo.P, 0, 100)

	for _, p := range ps {
		d := geo.Distance(a, b, c, p.X, p.Y)
		switch side := geo.PosL(p, lp, rp); {
		case side > 0:
			if d > maxU {
				maxU, div[0] = d, p
			}
			updown[0] = append(updown[0], p)
		case side < 0:
			if d > maxD {
				maxD, div[1] = d, p
			}
			updown[1] = append(updown[1], p)
		}
	}

	if len(updown[0]) > 0 {
		ret = append(ret, div[0])
	}
	if len(updown[1]) > 0 {
		ret = append(ret, div[1])
	}

	left := [][]geo.P{updown[0][:0], updown[1][:0]}
	right := [][]geo.P{make([]geo.P, 0, 25), make([]geo.P, 0, 25)}
	tri := [][]geo.P{[]geo.P{lp, rp, div[0]}, []geo.P{lp, rp, div[1]}}
	inter := []geo.P{geo.NearestP(div[0], lp, rp), geo.NearestP(div[1], lp, rp)}
	for i := 0; i < 2; i++ {
		for _, p := range updown[i] {
			if geo.InPolygon(p, tri[i]) {
				continue
			}
			if geo.PosL(p, inter[i], div[i]) > 0 {
				left[i] = append(left[i], p)
				continue
			}
			right[i] = append(right[i], p)
		}
	}

	ret = append(ret, solve(lp, div[0], left[0])...)
	ret = append(ret, solve(div[0], rp, right[0])...)
	ret = append(ret, solve(rp, div[1], left[1])...)
	ret = append(ret, solve(div[1], lp, right[1])...)
	return ret
}

func leftMost(ps []geo.P) geo.P {
	p := ps[0]
	for _, v := range ps[1:] {
		if v.X < p.X {
			p = v
		}
	}
	return p
}

func rightMost(ps []geo.P) geo.P {
	p := ps[0]
	for _, v := range ps[1:] {
		if v.X > p.X {
			p = v
		}
	}
	return p
}
