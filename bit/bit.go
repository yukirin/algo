package bit

// Add add value to BIT[pos]
func Add(bit []int64, pos int, v int64) {
	for x := pos; x < len(bit); x += x & -x {
		bit[x] += v
	}
}

// Sum return BIT[1~end]
func Sum(bit []int64, end int) int64 {
	ret := int64(0)
	for x := end; x > 0; x -= x & -x {
		ret += bit[x]
	}
	return ret
}

// Init init BIT
func Init(bit []int64) {
	for x := 1; x < len(bit); x++ {
		next := x + (x & -x)
		if next < len(bit) {
			bit[next] += bit[x]
		}
	}
}

// Add2D add value to BIT[b][a]
func Add2D(bit [][]int64, a, b int, v int64) {
	for y := b; y < len(bit); y += y & -y {
		for x := a; x < len(bit[0]); x += x & -x {
			bit[y][x] += v
		}
	}
}

// Sum2D return BIT[1~b][1~a]
func Sum2D(bit [][]int64, a, b int) int64 {
	ret := int64(0)
	for y := b; y > 0; y -= y & -y {
		for x := a; x > 0; x -= x & -x {
			ret += bit[y][x]
		}
	}
	return ret
}

// Init2D init BIT[][]
func Init2D(bit [][]int64) {
	for y := 1; y < len(bit); y++ {
		for x := 1; x < len(bit[0]); x++ {
			w, h := x+(x&-x), y+(y&-y)
			if h < len(bit) && w < len(bit[0]) {
				bit[h][w] += bit[y][x]
			}
		}
	}
}
