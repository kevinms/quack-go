package quack_test

import (
	"fmt"
	"github.com/kevinms/quack-go"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestStackPushPopLen(t *testing.T) {
	s := quack.NewStack(lessInt)
	if i := s.Pop(); i != nil {
		t.Fatalf("Pop() of empty queue should return nil, but got %v!", i)
	}

	// Must have LIFO semantics.
	in := []int{0, 1, 2, 3, 4, 5}
	for i, n := range in {
		s.Push(n)
		l := s.Len()
		if l != i+1 {
			t.Fatalf("Len() should be %d, but got %v?!", n, l)
		}
	}
	for i := len(in) - 1; i >= 0; i-- {
		if n := s.Pop(); n != in[i] {
			t.Fatalf("Pop() should return in reverse order, but got %d:%v?!", i, n)
		}
	}
}

func TestStackMin(t *testing.T) {
	s := quack.NewStack(lessInt)

	in := []int{5, 6, 4, 0, 1}
	mins := []int{5, 5, 4, 0, 0}

	for i, n := range in {
		s.Push(n)
		m := s.Min()
		if m != mins[i] {
			t.Fatalf("Min() should be %v, but got %v?!", mins[i], m)
		}
	}

	for i := len(in) - 1; i >= 0; i-- {
		m := s.Min()
		if m != mins[i] {
			t.Fatalf("Min() should be %v, but got %v?!", mins[i], m)
		}
		s.Pop()
	}

	m := s.Min()
	if m != nil {
		t.Fatalf("Min() should return 0 when empty, but got %v?!", m)
	}
}

func checkPop(t *testing.T, n interface{}, stack *[]int) {
	t.Helper()
	if n == nil && len(*stack) != 0 {
		t.Fatal("Stack is empty but it shouldn't be?!")
	}
	if n != nil && len(*stack) == 0 {
		t.Fatalf("Stack should be empty but got %d?!", n)
	}
	if len(*stack) > 0 {
		if n != (*stack)[len(*stack)-1] {
			t.Fatalf("Expected %d at top of stack, but got %d?!",
				(*stack)[len(*stack)-1], n)
		}
		*stack = (*stack)[:len(*stack)-1]
	}
}

func findMin(stack []int) interface{} {
	if len(stack) <= 0 {
		return nil
	}

	m := stack[0]
	for _, n := range stack[1:] {
		if n < m {
			m = n
		}
	}

	return m
}

func TestStackRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	s := quack.NewStack(lessInt)

	stack := make([]int, 0, 4096)

	// Push() and Pop() randomly.
	for i := 0; i < 10000; i++ {
		if rand.Uint64() > math.MaxUint64/2 {
			n := rand.Int()
			s.Push(n)
			stack = append(stack, n)
		} else {
			n := s.Pop()
			checkPop(t, n, &stack)
		}
		l := s.Len()
		if l != len(stack) {
			t.Fatalf("Expected stack length %d, got %d?!", len(stack), l)
		}

		m := s.Min()
		if m != findMin(stack) {
			fmt.Println(stack)
			t.Fatalf("Expected min of %d, but got %d?!\n", findMin(stack), m)
		}
	}

	// Pop() any remaining.
	for len(stack) > 0 {
		n := s.Pop()
		checkPop(t, n, &stack)
	}
	if len(stack) > 0 {
		t.Fatal("Exepected empty stack?!")
	}
}
