package arraylist

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// ArrayList is an abstract data type (ADT) that holds a sequential
// array of `*int` elements.
type ArrayList struct {
	data []*int
}

// Create returns ArrayList objects that can hold the number of
// elements equal to the number given by `size`.
func Create(size int) *ArrayList {
	if size < 1 {
		panic(errors.New("size must be positive"))
	}

	if size == math.MaxInt32 {
		panic(fmt.Errorf("size must be less than %d", math.MaxInt32))
	}

	return &ArrayList{
		data: make([]*int, 0, size),
	}
}

// Size returns the capacity of this list.
func (l *ArrayList) Size() int {
	return cap(l.data)
}

// Count returns the number of elements in this list.
func (l *ArrayList) Count() int {
	return len(l.data)
}

// IsEmpty returns `true` when there are no elements in the given
// list; `false“ otherwise.
func (l *ArrayList) IsEmpty() bool {
	return len(l.data) <= 0
}

// IsFull returns `true“ when no more elements can be inserted into
// the given list; `false` otherwise.
func (l *ArrayList) IsFull() bool {
	return len(l.data) >= cap(l.data)
}

// Insert appends the given value to the end of the list if there
// is room.
func (l *ArrayList) Insert(value *int) {
	if len(l.data) >= cap(l.data) {
		panic(errors.New("index overflow"))
	}

	l.data = append(l.data, value)
}

// InsertAt sets the element located at the given index `i` to the
// given value if there were at least i+1 elements in the list.
func (l *ArrayList) InsertAt(i int, value *int) {
	if i < 1 {
		panic(errors.New("index must be positive"))
	}

	if i == math.MaxInt32 {
		panic(fmt.Errorf("index must be less than %d", math.MaxInt32))
	}

	if i >= len(l.data) {
		panic(errors.New("index overflow"))
	}

	l.data[i] = value
}

// Get returns the element at the given index `i`.
func (l *ArrayList) Get(i int) *int {
	if len(l.data) <= 0 {
		panic(errors.New("array is empty"))
	}

	if i >= len(l.data) {
		panic(errors.New("index overflow"))
	}

	return l.data[i]
}

// ToString returns a representation of the given list suitable for
// printing.
func (l *ArrayList) ToString() string {
	result := "["

	for i, v := range l.data {
		if i == 0 {
			result += strconv.FormatInt(int64(*v), 10)
		} else {
			result += ", " + strconv.FormatInt(int64(*v), 10)
		}
	}

	result += "]"

	return result
}

// Search performs a linear search on the list for the given value.
// An element with the given value is returned if the element is
// equal to the given value, or nil is returned.
func (l *ArrayList) Search(value *int) *int {
	for _, v := range l.data {
		if *v == *value {
			return v
		}
	}

	return nil
}

// SearchRecursively performs a recursive linear search on the list
// for the given value. An element with the given value is returned
// if the element is equal to the given value, or nil is returned.
func (l *ArrayList) SearchRecursively(value *int) *int {
	return searchRecursively(l.data, value)
}

func searchRecursively(values []*int, v *int) *int {
	if len(values) <= 0 {
		return nil
	}

	if *values[0] == *v {
		return values[0]
	}

	return searchRecursively(values[1:], v)
}
