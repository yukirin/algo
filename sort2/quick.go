package sort2

import (
	"math/rand"
	"time"
)

// QuickR is random quick sort
func QuickR(a []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var quick func([]int)
	quick = func(a []int) {
		if len(a) <= 1 {
			return
		}

		p := a[r.Intn(len(a)-1)+1]

		i, j := 0, len(a)-1
		for {
			for ; a[i] < p; i++ {
			}

			for ; a[j] > p; j-- {
			}

			if i >= j {
				break
			}

			a[i], a[j] = a[j], a[i]
			i++
			j--
		}

		quick(a[:i])
		quick(a[i:])
	}
	quick(a)
}
