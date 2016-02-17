package zobrist

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Table2D is zobrist 2d hash table
type Table2D [][][]int64

// Hash2D generate zobrist 2D hash
func (t Table2D) Hash2D(board [][]int) int64 {
	h := int64(0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			h ^= t[i][j][board[i][j]]
		}
	}
	return h
}

// Init2D init zobrist hash table
func Init2D(w, h, kind int, bitW int) (Table2D, error) {
	if bitW > 63 {
		return nil, fmt.Errorf("Invalid bit width")
	}

	var z big.Int
	x, y := big.NewInt(2), big.NewInt(int64(bitW))
	mod := z.Sub(z.Exp(x, y, nil), big.NewInt(1)).Int64()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	table := make(Table2D, h)
	for i := 0; i < h; i++ {
		table[i] = make([][]int64, w)
		for j := range table[i] {
			table[i][j] = make([]int64, kind)
			for k := 0; k < kind; k++ {
				table[i][j][k] = r.Int63() % mod
			}
		}
	}
	return table, nil
}
