package quack

import (
	"container/list"
)

type node struct {
	min interface{}
	v   interface{}
}

// Stack is a LIFO stack that can return the minimum value in the stack in O(1)
// time.
//
// A Stack's worst case runtime of every public method is O(1).
type Stack struct {
	l    *list.List
	less LessFunc
}

// NewStack returns a new Stack.
func NewStack(less LessFunc) *Stack {
	return &Stack{l: list.New(), less: less}
}

// Push adds v onto the stack in O(1).
func (s *Stack) Push(v interface{}) {
	min := s.Min()
	if min == nil || s.less(v, min) {
		min = v
	}

	s.l.PushFront(node{min: min, v: v})
}

// Pop removes the oldest data from the stack in O(1).
func (s *Stack) Pop() interface{} {
	if s.Len() <= 0 {
		return nil
	}
	n := s.l.Remove(s.l.Front()).(node)
	return n.v
}

// Len returns the number of items in the stack in O(1).
func (s *Stack) Len() int {
	return s.l.Len()
}

// Min returns the smallest value in the stack in O(1).
func (s *Stack) Min() interface{} {
	if s.l.Len() <= 0 {
		return nil
	}

	d := s.l.Front().Value.(node)
	return d.min
}
