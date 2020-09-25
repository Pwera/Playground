cmake_minimum_required(VERSION 3.16)

project(project)
include(FetchContent)
FetchContent_Declare(
        googletest
        GIT_REPOSITORY https://github.com/google/googletest.git
        GIT_TAG        release-1.8.0
)
FetchContent_Declare(
        benchmark
        GIT_REPOSITORY https://github.com/google/benchmark.git
        GIT_TAG        v1.5.2
)
FetchContent_MakeAvailable(googletest benchmark)


add_executable(Main smartPointers/Main.cpp)
target_compile_features(Main PUBLIC cxx_std_20)

add_executable(Main2 smartPointers/Main2.cpp)
target_compile_features(Main2 PUBLIC cxx_std_20)

add_executable(concept concept/concepts.cpp)
set_property(TARGET concept PROPERTY CXX_STANDARD 20)

add_executable(polyBenchmark benchmark/benchmark1.cpp)
target_link_libraries(polyBenchmark PRIVATE benchmark::benchmark)
set_property(TARGET polyBenchmark PROPERTY CXX_STANDARD 20)