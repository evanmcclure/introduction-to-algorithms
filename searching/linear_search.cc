#include <limits>

// linear_search is an efficient algorithm for searching a small
// number of elements for a value x, and it solves the searching
// problem.
//
// The input is an array a, which is a sequence of n numbers
// <a_1, a_2, ..., a_n>, and a value x. The output is an index
// i such that x equals a[i] or the special value -2147483648
// if x does not appear in a.
//
// linear_search uses 1-origin indexing, so the first element in
// the array a is ignored.
//
// The size n is assumed to be 1 less than the number of elements
// in the array a.
int linear_search(int a[], int n, int x) {
    for (int i = 1; i <= n; ++i) {
        if (x == a[i]) {
            return i;
        }
    }

    return -2147483648;
}
