package arraylist

import (
	"math"
	"strconv"
	"testing"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	// Assert that the list can't be created with a negative size.
	assert.PanicsWithError(t, "size must be positive", func() {
		Create(-1)
	})

	// Assert that the list can't be created without capacity.
	assert.PanicsWithError(t, "size must be positive", func() {
		Create(0)
	})

	// Assert that the list capacity has an upper bound.
	assert.PanicsWithError(t, "size must be less than "+strconv.FormatInt(math.MaxInt32, 10), func() {
		Create(math.MaxInt32)
	})

	// Assert that the smallest list can be created.
	Create(1)

	// Assert that the largest list can be created.
	Create(math.MaxInt32 - 1)

	// Assert that an empty list can be created with a certain size.
	l := Create(5)
	assert.Equal(t, 5, l.Size())
	assert.Equal(t, 0, l.Count())
	assert.True(t, l.IsEmpty())
	assert.False(t, l.IsFull())
	assert.PanicsWithError(t, "array is empty", func() {
		_ = l.Get(0)
	})
	assert.PanicsWithError(t, "array is empty", func() {
		_ = l.Get(1)
	})
	assert.PanicsWithError(t, "array is empty", func() {
		_ = l.Get(4)
	})

	// Assert that the list can be filled.
	l.Insert(pointy.Pointer(1))
	l.Insert(pointy.Pointer(2))
	l.Insert(pointy.Pointer(3))
	l.Insert(pointy.Pointer(5))
	l.Insert(pointy.Pointer(7))
	assert.Equal(t, 1, *l.Get(0))
	assert.Equal(t, 2, *l.Get(1))
	assert.Equal(t, 3, *l.Get(2))
	assert.Equal(t, 5, *l.Get(3))
	assert.Equal(t, 7, *l.Get(4))
	assert.PanicsWithError(t, "index overflow", func() {
		l.Insert(pointy.Pointer(11))
	})
	assert.PanicsWithError(t, "index overflow", func() {
		l.Get(5)
	})

	// Assert that the list can be printed.
	assert.Equal(t, "[1, 2, 3, 5, 7]", l.ToString())

	// Assert that the list can be searched.
	assert.Nil(t, l.Search(pointy.Pointer(11)))
	assert.Equal(t, 7, *l.Search(pointy.Pointer(7)))
	assert.Nil(t, l.SearchRecursively(pointy.Pointer(11)))
	assert.Equal(t, 7, *l.SearchRecursively(pointy.Pointer(7)))

	// Assert that a value can be inserted anywhere.
	assert.PanicsWithError(t, "index must be positive", func() {
		l.InsertAt(-1, pointy.Pointer(13))
	})
	assert.PanicsWithError(t, "index must be positive", func() {
		l.InsertAt(0, pointy.Pointer(13))
	})
	assert.PanicsWithError(t, "index must be less than "+strconv.FormatInt(math.MaxInt32, 10), func() {
		l.InsertAt(math.MaxInt32, pointy.Pointer(13))
	})
	assert.PanicsWithError(t, "index overflow", func() {
		l.InsertAt(5, pointy.Pointer(13))
	})
	l.InsertAt(2, pointy.Pointer(11))
	assert.Equal(t, 11, *l.Get(2))

	// resize
	// remove()
	// removeAt(i)
	// compare()
	// copy()
	// clone()

}
