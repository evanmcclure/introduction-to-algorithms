#include <limits>

#include <gtest/gtest.h>

#include "sorting.h"

TEST(SelectionSortTest, SelectionSort) {
  int a[] = { INT_MAX, 31, 41, 59, 26, 41, 58 };

  selection_sort(a, 6);

  int expected[] = { INT_MAX, 26, 31, 41, 41, 58, 59 };

  ASSERT_EQ(sizeof(a) / sizeof(int), sizeof(expected) / sizeof(int));
  
  for (int i = 0; i < sizeof(a) / sizeof(int); i++) {
    EXPECT_EQ(a[i], expected[i]) << "Arrays a and expected differ at index " << i;
  }
}
