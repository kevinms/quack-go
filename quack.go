// Implements a FIFO Queue that can return the minimum value in the queue in
// O(1) time. The queue is composed of two stacks, hence the name 'quack'.
//
// The worst case runtime of every public method is O(1) except Pop(), which is
// amortized O(1).
//
// Operation  Runtime
// Push()     O(1) amortized
// Pop()      O(1)
// Len()      O(1)
// Min()      O(1)
package quack

type LessFunc func(a, b interface{}) bool

type Quack struct {
	in   *Stack
	out  *Stack
	less LessFunc
}

// NewQuack returns a new Quack.
func NewQuack(less LessFunc) *Quack {
	return &Quack{
		in:   NewStack(less),
		out:  NewStack(less),
		less: less,
	}
}

// Pushes v onto the quack in O(1).
func (q *Quack) Push(v interface{}) {
	q.in.Push(v)
}

// Pops the oldest data from the quack in amortized O(1).
func (q *Quack) Pop() interface{} {
	if i := q.out.Pop(); i != nil {
		return i
	}
	for q.in.Len() > 0 {
		q.out.Push(q.in.Pop())
	}

	return q.out.Pop()
}

// Returns the number of items in the quack in O(1).
func (q *Quack) Len() int {
	return q.in.Len() + q.out.Len()
}

// Returns the smallest value in the quack in O(1).
func (q *Quack) Min() interface{} {
	v1 := q.in.Min()
	v2 := q.out.Min()

	if v1 == nil {
		return v2
	}
	if v2 == nil {
		return v1
	}

	if q.less(v1, v2) {
		return v1
	}
	return v2
}
