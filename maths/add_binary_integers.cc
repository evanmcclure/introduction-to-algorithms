#include <limits>

// add_binary_integers is an algorithm that solves the problem of
// adding two n-bit binary integers a and b, and stored in two
// n-element arrays A[0:n-1] and B[0:n-1], where each element is
// either 0 or 1, a being the sum from i = 0 to n-1 of A[i] x 2^i,
// and b being the sum from i = 0 to n-1 of B[i] x 2^i. The sum
// c = a + b of the two integers is stored in binary form in an
// (n + 1) element array C[0:n], where c is the sum from i = 0 to
// n of C[i] x 2^i.
//
// The input array a represents A and the input array b represents
// B. The output array c represents C.
//
// The size n is assumed to be 1 less than the number of elements
// in the array c.
void add_binary_integers(int a[], int b[], int n, int c[]) {
    int carry = 0;

    int i;
    for (i = 0; i < n; ++i) {
        int sum = a[i] + b[i] + carry;
        c[i] = sum % 2;
        carry = sum / 2;
    }

    c[i] = carry;
    return;
}
