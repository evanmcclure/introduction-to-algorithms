
// selection_sort solves the sorting problem.
//
// The input is an array a, which is a sequence of n numbers
// <a_1, a_2, ..., a_n>. The output is the array a, which is a
// permutation <a_1', a_2', ..., a_n'> of the input sequence such
// that a_1' <= a_2' <= ... <= a_n'.
//
// selection_sort uses 1-origin indexing, so the first element in
// the array a is ignored.
void selection_sort(int a[], int n) {

    for (int i = 1; i < n; i++) {
        // Find the smallest element of a[i:n].
        int k = i;
        for (int j = i; j <=n; j++) {
            if (a[j] < a[k]) {
                k = j;
            }
        }

        // Exchange the smallest element with a[i].
        int temp = a[i];
        a[i] = a[k];
        a[k] = temp;
    }
}

