package queue

import (
	"fmt"
)

type Queue struct {
	n    int
	head int
	tail int
	data []int
}

func Create(n int) *Queue {
	return &Queue{
		n:    n,
		head: 0,
		tail: 0,
		data: make([]int, n),
	}
}

func (q *Queue) Enqueue(x int) {
	fmt.Printf("Enter Enqueue(x=%d, head=%d, tail=%d, n=%d, data=%v)\n", x, q.head, q.tail, q.n, q.data)

	q.data[q.tail] = x

	if q.tail == q.n-1 {
		q.tail = 0
	} else {
		q.tail = q.tail + 1
	}

	fmt.Printf("Leave Enqueue(head=%d, tail=%d, n=%d, data=%v)\n", q.head, q.tail, q.n, q.data)
}

func (q *Queue) Dequeue() int {
	fmt.Printf("Enter Dequeue(head=%d, tail=%d, n=%d, data=%v)\n", q.head, q.tail, q.n, q.data)

	x := q.data[q.head]

	if q.head == q.n-1 {
		q.head = 0
	} else {
		q.head = q.head + 1
	}

	fmt.Printf("Leave Dequeue(head=%d, tail=%d, n=%d, data=%v) -> %d\n", q.head, q.tail, q.n, q.data, x)
	return x
}
