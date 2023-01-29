cmake_minimum_required(VERSION 3.5)

project(IntroductionToAlgorithms LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)

include(FetchContent)
FetchContent_Declare(
    googletest
    URL https://github.com/google/googletest/archive/refs/tags/v1.13.0.zip
)

FetchContent_MakeAvailable(googletest)

# add_library(IntroductionToAlgorithms sorting/insertion_sort.cc)

################################
# Testing
################################

enable_testing()

add_executable(
    insertion_sort_test
    sorting/insertion_sort.cc
    sorting/insertion_sort_test.cc
)

target_link_libraries(
    insertion_sort_test
    GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(insertion_sort_test)
