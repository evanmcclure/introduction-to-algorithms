
// insertion_sort is an efficient algorithm for sorting a small
// number of elements.
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
