package quack_test

import (
	"fmt"
	"quack"
)

func lessInt(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func ExampleQuack() {
	q := quack.NewQuack(lessInt)

	a := []int{2, 0, 8, 3, 4}
	for _, n := range a {
		q.Push(n)
	}

	for q.Len() > 0 {
		fmt.Println(q.Min())
		q.Pop()
	}

	// Output: 0
	// 0
	// 3
	// 3
	// 4
}

func ExampleStack() {
	s := quack.NewStack(lessInt)

	a := []int{4, 3, 8, 0, 2}
	for _, n := range a {
		s.Push(n)
	}

	for s.Len() > 0 {
		fmt.Println(s.Min())
		s.Pop()
	}

	// Output: 0
	// 0
	// 3
	// 3
	// 4
}
