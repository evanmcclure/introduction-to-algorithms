package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	a := []int{5, 2, 4, 6, 1, 3}

	InsertSort(a)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, a)
}
