#include <limits>

#include <gtest/gtest.h>

#include "searching.h"

TEST(LinearSearchTest, LinearSearch) {
  int a[] = { INT_MAX, 31, 41, 59, 26, 41, 58 };

  EXPECT_EQ(1, linear_search(a, 6, 31));

  EXPECT_EQ(6, linear_search(a, 6, 58));

  EXPECT_EQ(4, linear_search(a, 6, 26));

  EXPECT_EQ(-2147483648, linear_search(a, 6, 99));
}
