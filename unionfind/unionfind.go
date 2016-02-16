package unionfind

// Tree is union find tree
type Tree []int

// Init is n要素で初期化
func (t *Tree) Init(n int) {
	for i := 0; i < n; i++ {
		*t = append(*t, i)
	}
}

// Find is 木の根を求める
func (t Tree) Find(x int) int {
	if t[x] == x { // 根
		return x
	}

	t[x] = t.Find(t[x]) // 経路圧縮
	return t[x]
}

// Same is xとyが同じ集合に属する否か
func (t Tree) Same(x, y int) bool {
	return t.Find(x) == t.Find(y)
}

// Union is xとyの属する集合を併合
func (t Tree) Union(x, y int) {
	x, y = t.Find(x), t.Find(y)

	if x == y {
		return
	}

	t[x] = y
}
