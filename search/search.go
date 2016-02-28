// Package search implements the binary, ternary, golden section Search
// http://d.hatena.ne.jp/nodchip/20090303/1236058357
package search

import "math"

// Loop count of search
const (
	BSMaxLoop = 56
	TSMaxLoop = 88
	GSMaxLoop = 78
)

// GRatio is golden ratio
var GRatio = (1 + math.Sqrt(5)) * 0.5

// Binary is binary search
func Binary(l, r, k float64, f func(float64) float64) float64 {
	for i := 0; i < BSMaxLoop; i++ {
		m := (l + r) / 2
		if f(m) > k {
			r = m
			continue
		}
		l = m
	}
	return (l + r) / 2
}

// TernaryU is ternary serch(convex upward)
func TernaryU(l, r float64, f func(float64) float64) float64 {
	for i := 0; i < TSMaxLoop; i++ {
		low, high := (l*2+r)/3, (l+r*2)/3
		if f(low) > f(high) {
			r = high
			continue
		}
		l = low
	}
	return (l + r) * 0.5
}

// TernaryD is ternary search(convex downward)
func TernaryD(l, r float64, f func(float64) float64) float64 {
	for i := 0; i < TSMaxLoop; i++ {
		low, high := (l*2+r)/3, (l+r*2)/3
		if f(low) < f(high) {
			r = high
			continue
		}
		l = low
	}
	return (l + r) * 0.5
}

// GoldenU is golden section search(convex upward)
func GoldenU(l, r float64, f func(float64) float64) float64 {
	low, high := (l*GRatio+r)/(1+GRatio), (l+GRatio*r)/(1+GRatio)
	lV, rV := f(low), f(high)
	for i := 0; i < GSMaxLoop; i++ {
		if lV > rV {
			r, rV = (l+GRatio*r)/(1+GRatio), lV
			lV = f((l*GRatio + r) / (1 + GRatio))
			continue
		}
		l, lV = (l*GRatio+r)/(1+GRatio), rV
		rV = f((l + GRatio*r) / (1 + GRatio))
	}

	if lV > rV {
		r = (l + GRatio*r) / (1 + GRatio)
	} else {
		l = (l*GRatio + r) / (1 + GRatio)
	}
	return (l + r) * 0.5
}

// GoldenD is golden section search(convex downward)
func GoldenD(l, r float64, f func(float64) float64) float64 {
	low, high := (l*GRatio+r)/(1+GRatio), (l+GRatio*r)/(1+GRatio)
	lV, rV := f(low), f(high)
	for i := 0; i < GSMaxLoop; i++ {
		if lV < rV {
			r, rV = (l+GRatio*r)/(1+GRatio), lV
			lV = f((l*GRatio + r) / (1 + GRatio))
			continue
		}
		l, lV = (l*GRatio+r)/(1+GRatio), rV
		rV = f((l + GRatio*r) / (1 + GRatio))
	}

	if lV < rV {
		r = (l + GRatio*r) / (1 + GRatio)
	} else {
		l = (l*GRatio + r) / (1 + GRatio)
	}
	return (l + r) * 0.5
}
