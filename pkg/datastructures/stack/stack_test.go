package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Create(3)

	assert.True(t, s.Empty())

	s.Push(1)

	assert.False(t, s.Empty())

	s.Push(2)
	s.Push(3)

	assert.PanicsWithError(t, "overflow", func() {
		s.Push(4)
	})

	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())

	assert.True(t, s.Empty())

	assert.PanicsWithError(t, "underflow", func() {
		s.Pop()
	})
}
