cmake_minimum_required(VERSION 3.5)

project(IntroductionToAlgorithms LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)

include(FetchContent)
FetchContent_Declare(
    googletest
    URL https://github.com/google/googletest/archive/refs/tags/v1.13.0.zip
)

FetchContent_MakeAvailable(googletest)

################################
# Testing
################################

enable_testing()

add_executable(
    all_code
    searching/linear_search.cc
    searching/linear_search_test.cc
    sorting/insertion_sort.cc
    sorting/insertion_sort_test.cc
)

target_link_libraries(
    all_code
    GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(all_code)
