// SliceTricks golang
// https://github.com/golang/go/wiki/SliceTricks
package main

import (
	"fmt"
)

func main() {
	a := newSlice()
	b := newSlice()

	// AppendVecotr
	a = append(a, b...)
	fmt.Println(a)

	// Copy
	a = newSlice()
	b = make([]int, len(a))
	copy(b, a)
	fmt.Println(b)
	// or
	b = append([]int(nil), a...)
	fmt.Println(b)

	// Cut
	a = newSlice()
	a = append(a[:3], a[7:]...)
	fmt.Println(a)

	// Delete
	a = newSlice()
	a = append(a[:3], a[3+1:]...)
	fmt.Println(a)
	// or
	a = newSlice()
	a = a[:3+copy(a[3:], a[3+1:])]
	fmt.Println(a)

	// Delete without preserving order
	a = newSlice()
	a[3] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Println(a)

	// Fix Cut
	a = newSlice()
	i, j := 3, 7
	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = 0 // or the zero value of T
	}
	a = a[:len(a)-j+i]
	fmt.Println(a)

	// Fix Delete
	a = newSlice()
	copy(a[3:], a[3+1:])
	a[len(a)-1] = 0 // or the zero value of T
	a = a[:len(a)-1]
	fmt.Println(a)
	// or, more simply:
	a = newSlice()
	a, a[len(a)-1] = append(a[:3], a[3+1:]...), 0
	fmt.Println(a)

	// Fix Delete without preserving order
	a = newSlice()
	a[3] = a[len(a)-1]
	a[len(a)-1] = 0
	a = a[:len(a)-1]
	fmt.Println(a)

	// Expand
	a = newSlice()
	a = append(a[:i], append(make([]int, 5), a[i:]...)...)
	fmt.Println(a)

	// Extend
	a = newSlice()
	a = append(a, make([]int, 5)...)
	fmt.Println(a)

	// Insert
	a = newSlice()
	a = append(a[:3], append([]int{22}, a[i:]...)...)
	fmt.Println(a)

	// Insert (recommend)
	a = newSlice()
	a = append(a, 0)
	copy(a[3+1:], a[3:])
	a[3] = 55
	fmt.Println(a)

	// InsertVector
	a = newSlice()
	b = newSlice()
	a = append(a[:i], append(b, a[i:]...)...)
	fmt.Println(a)

	// Pop
	a = newSlice()
	x, a := a[len(a)-1], a[:len(a)-1]
	fmt.Println(x, a)

	// Push
	a = newSlice()
	a = append(a, 3)
	fmt.Println(a)

	// Shift
	a = newSlice()
	x, a = a[0], a[1:]
	fmt.Println(x, a)

	// Unshift
	a = newSlice()
	a = append([]int{8}, a...)
	fmt.Println(a)

	// Filtering without allocating
	a = newSlice()
	b, f := a[:0], func(n int) bool { return n%2 == 1 }
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}

	// Reversing
	a = newSlice()
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)
	// The same thing, except with two indices:
	a = newSlice()
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	fmt.Println(a)
}

func newSlice() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}
