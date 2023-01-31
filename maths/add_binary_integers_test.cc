#include <limits>

#include <gtest/gtest.h>

#include "maths.h"

TEST(AddBinaryIntegersTest, AddBinaryIntegers) {
  int a[] = { 1, 0, 0, 0 }; // the value 1 as a 4-bit binary integer
  int b[] = { 1, 1, 0, 0 }; // the value 3 as a 4-bit binary integer
  int c[5] = {0}; // the value of adding array a to array b as a 5-bit binary integer

  add_binary_integers(a, b, 4, c);

  int expected[] = { 0, 0, 1, 0, 0 }; // the value of adding array a to array b as a binary 5-bit integer, which is 4

  ASSERT_EQ(sizeof(c) / sizeof(int), sizeof(expected) / sizeof(int));
  
  for (int i = 0; i < sizeof(c) / sizeof(int); i++) {
    EXPECT_EQ(c[i], expected[i]) << "Arrays c and expected differ at index " << i;
  }
}
