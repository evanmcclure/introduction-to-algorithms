package arraylist

import (
	"errors"
)

// SortedArrayList is an abstract data type (ADT) that holds a
// sequential array of `*int` elements that are sorted.
type SortedArrayList struct {
	*ArrayList
}

// CreateSortedArrayList returns SortedArrayList objects that can
// hold the number of elements equal to the number given by `size`.
// Elements are inserted in order.
func CreateSortedArrayList(size int) *SortedArrayList {
	arrayList := Create(size)

	return &SortedArrayList{
		ArrayList: arrayList,
	}
}

// Insert puts the given value in its proper location in the given
// using index-insert (insert sort).
func (l *SortedArrayList) Insert(value *int) {
	if len(l.data) >= cap(l.data) {
		panic(errors.New("index overflow"))
	}

	// Insert the value at the end in order to give space to access
	// data using an index. The value will be shifted off. So, it
	// will have to be added before this function returns.
	l.data = append(l.data, value)

	pos := 0

	// Find the position where the value should be inserted.
	for i, v := range l.data {
		if *value < *v {
			pos = i
			break
		}
	}

	// Shift everything beyond that positon to the right.
	for i := len(l.data) - 1; i > pos; i-- {
		l.data[i] = l.data[i-1]
	}

	l.data[pos] = value
}
