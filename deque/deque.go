package deque

// Deque is double ended queue
type Deque struct {
	s       []int
	h, t, c int
}

// New to generate a deque, Create an empty deque if a == nil
func New(a []int, length int) *Deque {
	q := new(Deque)
	if a == nil {
		q.s = make([]int, length)
		return q
	}

	q.s, q.h, q.t, q.c = a, 0, len(a), len(a)
	return q
}

// PushBack add an element to the end
func (q *Deque) PushBack(x int) {
	if q.c == len(q.s) {
		q.update()
	}
	q.c++
	q.s[q.t] = x
	q.t++
	if q.t == len(q.s) {
		q.t = 0
	}
}

// PushFront add an element to the top
func (q *Deque) PushFront(x int) {
	if q.c == len(q.s) {
		q.update()
	}
	q.c++
	q.h--
	if q.h < 0 {
		q.h = len(q.s) - 1
	}
	q.s[q.h] = x
}

// PopBack retrieve the last element
func (q *Deque) PopBack() int {
	if q.c == 0 {
		return 0
	}

	q.c--
	q.t--
	if q.t < 0 {
		q.t = len(q.s) - 1
	}
	return q.s[q.t]
}

// PopFront take out the head of the element
func (q *Deque) PopFront() int {
	if q.c == 0 {
		return 0
	}

	q.c--
	v := q.s[q.h]
	q.h++
	if q.h == len(q.s) {
		q.h = 0
	}
	return v
}

// Back examine the end of the element
func (q *Deque) Back() int {
	index := q.t - 1
	if index < 0 {
		index = len(q.s) - 1
	}
	return q.s[index]
}

// Front examine the top element
func (q *Deque) Front() int {
	return q.s[q.h]
}

// Get get the element
func (q *Deque) Get(i int) int {
	index := q.h + i
	if index >= len(q.s) {
		index -= len(q.s)
	}
	return q.s[index]
}

// Len is len
func (q *Deque) Len() int {
	return q.c
}

func (q *Deque) update() {
	tmp := make([]int, len(q.s)*2)
	copy(tmp, q.s[q.h:])
	copy(tmp, q.s[:q.h])
	q.h, q.t, q.s = 0, q.c, tmp
}
