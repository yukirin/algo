package sort2

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestQuickR(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 2000; i++ {
		ns := make([]int, 2000)
		for j := 0; j < 2000; j++ {
			ns[j] = r.Int()
		}

		QuickR(ns)

		if ok := sort.IntsAreSorted(ns); !ok {
			t.Errorf("sort.IntsAreSorted(ns) = %v; want %v", ok, true)
		}
	}
}
