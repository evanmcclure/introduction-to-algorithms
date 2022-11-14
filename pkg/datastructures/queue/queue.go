package queue

import (
	"fmt"
)

type Queue struct {
	length int
	head   int
	tail   int
	data   []int
}

func Create(n int) *Queue {
	return &Queue{
		length: n,
		head:   1,
		tail:   1,
		data:   make([]int, n+1),
	}
}

func (q *Queue) Enqueue(x int) {
	fmt.Printf("Enqueue(x=%d, head=%d, tail=%d, length=%d)\n", x, q.head, q.tail, q.length)

	q.data[q.tail] = x

	if q.tail == q.length {
		q.tail = 1
	} else {
		q.tail = q.tail + 1
	}

	fmt.Printf("Enqueue: head=%d, tail=%d, length=%d, data=%v\n", q.head, q.tail, q.length, q.data)
}

func (q *Queue) Dequeue() int {
	x := q.data[q.head]

	if q.head == q.length {
		q.head = 1
	} else {
		q.head = q.head + 1
	}

	fmt.Printf("Dequeue: head=%d, tail=%d, length=%d, data=%v\n", q.head, q.tail, q.length, q.data)
	return x
}
