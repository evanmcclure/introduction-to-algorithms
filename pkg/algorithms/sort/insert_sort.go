package sort

// InsertSort sorts array a in O(n^2) time using O(1) space.
func InsertSort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		for i > -1 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}

		a[i+1] = key
	}
}
