package arraylist

import (
	"testing"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestSortedArrayList(t *testing.T) {
	// Assert that the list is sorted regardless of how it gets filled.
	l := CreateSortedArrayList(5)
	l.Insert(pointy.Pointer(7))
	l.Insert(pointy.Pointer(1))
	l.Insert(pointy.Pointer(5))
	l.Insert(pointy.Pointer(2))
	l.Insert(pointy.Pointer(3))
	assert.Equal(t, "[1, 2, 3, 5, 7]", l.ToString())
}
