
// insertion_sort is an efficient algorithm for sorting a small
// number of elements, and it solves the sorting problem.
//
// The input is an array a, which is a sequence of n numbers
// <a_1, a_2, ..., a_n>. The output is the array a, which is a
// permutation <a_1', a_2', ..., a_n'> of the input sequence such
// that a_1' <= a_2' <= ... <= a_n'.
//
// insertion_sort uses 1-origin indexing, so the first element in
// the array a is ignored.
//
// The size n is assumed to be 1 less than the number of elements
// in the array a.
void insertion_sort(int a[], int n) {
    for (int i = 2; i <= n; ++i) {
        int key = a[i];
        int j = i - 1;
        // Insert a[i] into the sorted subarray a[1:i-1].
        while (j > 0 && a[j] > key) {
            a[j+1] = a[j];
            --j;
        }
        a[j+1] = key;
    }
    return;
}
