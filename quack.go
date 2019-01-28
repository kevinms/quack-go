/*
A Quack is a FIFO Queue that can return the minimum value in the queue in
O(1) time. The name 'quack' is a smash up of 'queue' and 'stack', because
the queue is implemented using two stacks.

A Quack's worst case runtime of every public method is O(1) except Pop(),
which is amortized O(1).

	func lessInt(a, b interface{}) bool {
		return a.(int) < b.(int)
	}

	func main() {
		q := quack.NewQuack(lessInt)

		n := 1000000
		for i := 0; i < n; i++ {
			q.Push(rand.Int())
		}

		fmt.Printf("Len: %v, Min: %v\n", q.Len(), q.Min())

		for q.Len() > n/2 {
			q.Pop()
		}

		fmt.Printf("Len: %v, Min: %v\n", q.Len(), q.Min())
	}
*/
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
