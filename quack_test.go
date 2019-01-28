package quack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := NewQuack(lessInt)
	if q.in == nil || q.out == nil {
		t.Fatal("Failed to initialize Stack(s)?!")
	}
	if q.less == nil {
		t.Fatal("Failed to initialize less func pointer?!")
	}
}

type action struct {
	what string
	n    int
	min  interface{}
}

func TestQuackMin(t *testing.T) {
	q := NewQuack(lessInt)

	actions := []action{
		{"push", 5, 5},
		{"push", 6, 5},
		{"push", 4, 4},
		{"push", 0, 0},
		{"push", 1, 0},
		{"pop", 5, 0},
		{"pop", 6, 0},
		{"pop", 4, 0},
		{"pop", 0, 1},
		{"pop", 1, nil},
		{"push", 5, 5},
		{"push", 6, 5},
		{"push", 4, 4},
		{"pop", 5, 4},
		{"push", 5, 4},
		{"push", 3, 3},
	}

	length := 0
	for _, a := range actions {
		switch a.what {
		case "push":
			length++
			q.Push(a.n)
		case "pop":
			if length > 0 {
				length--
			}
			n := q.Pop()
			if n != a.n {
				t.Fatalf("Pop() should be %v, but got %v?!", a.n, n)
			}
		default:
			t.Fatalf("Unexpected test action %s?!", a.what)
		}

		m := q.Min()
		if m != a.min {
			fmt.Println(a)
			t.Fatalf("Min() should be %v, but got %v?!", a.min, m)
		}
		l := q.Len()
		if l != length {
			t.Fatalf("Len() should be %v, but got %v?!", length, l)
		}
	}
}
