# quack-go
Queue that can return the minimum element in O(1) time where no operation is worse than O(1) amortized.

The name 'quack' is a smash up of 'queue' and 'stack', because the queue is implemented using two stacks.

### Example usage:

```
package main

import (
  "fmt"
  quack "github.com/kevinms/quack-go"
  "math/rand"
)

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
```
