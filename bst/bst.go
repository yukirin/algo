package bst

import "fmt"

// Node is binary search tree Value
type Node struct {
	Val int64

	L, R *Node
	P    *Node
}

// Search search node
func Search(r *Node, n int64) (*Node, error) {
	for r != nil {
		switch {
		case r.Val == n:
			return r, nil
		case n < r.Val:
			r = r.L
		default:
			r = r.R
		}
	}
	return nil, fmt.Errorf("not found")
}

// Insert insert Node
func Insert(r, n *Node) {
	var p *Node
	for {
		switch {
		case r == nil:
			if n.Val < p.Val {
				p.L, n.P = n, p
				return
			}
			p.R, n.P = n, p
			return
		case n.Val < r.Val:
			r, p = r.L, r
		case n.Val >= r.Val:
			r, p = r.R, r
		}
	}
}

// Delete delete Node, Delete returns a new root
func Delete(r *Node, n int64) *Node {
	del, err := Search(r, n)
	if err != nil {
		return r
	}

	if del.L != nil && del.R != nil {
		sn := del.L
		for ; sn.R != nil; sn = sn.R {
		}

		del.Val = sn.Val
		if sn.L != nil {
			replace(sn, sn.L)
			return r
		}
		replace(sn, nil)
		return r
	}
	child := del.L
	if del.R != nil {
		child = del.R
	}
	replace(del, child)
	if del.P == nil {
		if child != nil {
			child.P = nil
		}
		return child
	}
	return r
}

func replace(n, m *Node) {
	if n.P == nil {
		return
	}
	if m != nil {
		m.P = n.P
	}
	if n.P.R == n {
		n.P.R = m
		return
	}
	n.P.L = m
}

// Enum lists the binary search tree
func Enum(n *Node, fn func(*Node)) {
	if n.L != nil {
		Enum(n.L, fn)
	}
	fn(n)
	if n.R != nil {
		Enum(n.R, fn)
	}
}
