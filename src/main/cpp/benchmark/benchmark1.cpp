#include <benchmark/benchmark.h>
#include "smartPointers/Main.hpp"
#include "smartPointers/Main2.hpp"
#include "smartPointers/Main3.hpp"
#include <iostream>

static void BM_todov1(benchmark::State &state) {
    std::vector<v1::GeoObj *> vector = v1::createData(state.range(0));
    for (auto _ : state) {
        v1::drawElements(vector);

    }
    v1::removeElements(vector);
}

static void BM_todov2(benchmark::State &state) {
    std::vector<v2::GeoPtr> vector = v2::createData(state.range(0));
    for (auto _ : state)
        v2::drawElements(vector);
}

static void BM_todov4(benchmark::State &state) {
    std::vector<v2::GeoPtr> vector = v2::createData(state.range(0));
    for (auto _ : state)
        v2::drawElements2(vector);
}

static void BM_todov3(benchmark::State &state) {
    std::vector<v3::GeoPtr> vector = v3::createData(state.range(0));
    for (auto _ : state)
        v3::drawElements(vector);
}

BENCHMARK(BM_todov1)->Arg(10000)->Arg(50000)->Arg(100000)->Arg(1000000);
BENCHMARK(BM_todov2)->Arg(10000)->Arg(50000)->Arg(100000)->Arg(1000000);
BENCHMARK(BM_todov4)->Arg(10000)->Arg(50000)->Arg(100000)->Arg(1000000);
BENCHMARK(BM_todov3)->Arg(10000)->Arg(50000)->Arg(100000)->Arg(1000000);

BENCHMARK_MAIN();