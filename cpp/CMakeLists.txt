cmake_minimum_required(VERSION 3.16)

project(project)
include(FetchContent)
FetchContent_Declare(
        FunctionalPlus
        GIT_REPOSITORY https://github.com/Dobiasd/FunctionalPlus
        GIT_TAG        v0.2.18-p0
)
FetchContent_Declare(
        cpr
        GIT_REPOSITORY https://github.com/libcpr/cpr
        GIT_TAG        1.8.3
)
FetchContent_Declare(
        crow
        GIT_REPOSITORY https://github.com/ipkn/crow
        GIT_TAG        v0.1
)
FetchContent_Declare(
        oatpp
        GIT_REPOSITORY https://github.com/oatpp/oatpp
        GIT_TAG        1.3.0
)

FetchContent_MakeAvailable(FunctionalPlus cpr crow oatpp)

set(CMAKE_CXX_STANDARD 20)

add_executable(ff main.cpp)
target_compile_features(ff PUBLIC cxx_std_20)
target_link_libraries(ff PUBLIC fplus cpr CURL::libcurl oatpp)
#target_include_directories(ff PUBLIC)


#add_executable(Main smartPointers/Main.cpp)
#target_compile_features(Main PUBLIC cxx_std_20)
#
#add_executable(Main2 smartPointers/Main2.cpp)
#target_compile_features(Main2 PUBLIC cxx_std_20)
#
#add_executable(concept concept/concepts.cpp)
#set_property(TARGET concept PROPERTY CXX_STANDARD 20)
#target_compile_features(concept PUBLIC cxx_std_20)
#
#add_executable(polyBenchmark benchmark/benchmark1.cpp)
#target_link_libraries(polyBenchmark PRIVATE benchmark::benchmark)
#set_property(TARGET polyBenchmark PROPERTY CXX_STANDARD 20)