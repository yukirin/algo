package deque

import "container/list"
import "testing"

func BenchmarkDeque(b *testing.B) {
	q := New(nil, 2000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			q.PushBack(j)
			q.PushFront(j)
		}

		for q.Len() > 0 {
			q.PopFront()
		}
	}
}

func BenchmarkList(b *testing.B) {
	l := list.New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			l.PushBack(j)
			l.PushFront(j)
		}

		for l.Len() > 0 {
			l.Remove(l.Back())
		}
	}
}
