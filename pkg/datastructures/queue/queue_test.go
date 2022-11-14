package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Create(3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())

	q.Enqueue(5)
	assert.Equal(t, 5, q.Dequeue())

	q.Enqueue(6)
	assert.Equal(t, 6, q.Dequeue())
}
