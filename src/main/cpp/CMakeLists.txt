
#add_executable(Main smartPointers/Main.hpp)
#target_compile_features(Main PUBLIC cxx_std_17)
#
#add_executable(Main2 smartPointers/Main2.hpp)
#target_compile_features(Main2 PUBLIC cxx_std_17)
#
#add_executable(Main3 smartPointers/Main3.hpp)
#target_compile_features(Main3 PUBLIC cxx_std_17)

#set_property(TARGET Main3 PROPERTY CXX_STANDARD 17)



find_package(benchmark)
if(TARGET benchmark::benchmark)
    add_executable(polyBenchmark benchmark/benchmark1.cpp)
    target_link_libraries(polyBenchmark PRIVATE benchmark::benchmark)
    target_include_directories(polyBenchmark PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})
    set_property(TARGET polyBenchmark PROPERTY CXX_STANDARD 17)
endif()