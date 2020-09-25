#include <numeric>
#include <array>
#include <ranges>
#include <iostream>
#include <list>
#include <cmath>
#include <concepts>
#include <vector>
#include <type_traits>
#include <string>
#include <chrono>
#include <ctime>
struct MyStruct{
    inline int addOne(int i){return i++;};
};
struct BetterMyStruct : MyStruct{};
MyStruct* myStruct;

template<typename T>
concept MySuperType = requires (T t){
    {t.addOne()} ->  std::integral;
};
template <class T, class U>
concept Derived = std::is_base_of<U, T>::value;

static_assert(std::is_void_v<void>);
static_assert(std::is_void_v<const void>);
static_assert(std::is_void_v< void const>);

static_assert(not std::is_void_v<int>);

static_assert(std::is_null_pointer_v<decltype(nullptr)>);

static_assert(std::is_pointer_v<void*>);
static_assert(std::is_pointer_v<decltype(myStruct)>);
static_assert(std::is_pointer_v<MyStruct*>);


static_assert(std::is_same_v<int,int>);
static_assert(std::is_same_v<int,std::conditional_t<std::is_pointer_v<MyStruct*>,int,MyStruct>>);
static_assert(std::is_same_v<MyStruct, std::conditional_t<not  std::is_pointer_v<MyStruct*>,int,MyStruct>>);

std::integral auto  func(std::derived_from<MyStruct> auto str){
    return str.addOne(1);
}

template <typename Vec>
using Scalar = std::decay_t<decltype(Vec()[0])>;

auto norm(const std::vector<double>& vec) -> double{
    double result = 0;
    for(size_t  i = 0; i<vec.size();i++){
        result += vec[i] * vec[i];
    }
    return std::sqrt(result);
}

template <typename Vec>
concept FloatVec = std::floating_point<Scalar<Vec>> &&
                   requires (Vec t){
                       {t.size() } -> std::integral;
                   };

template <FloatVec Vec>
auto norm2( const Vec & vec) -> Scalar<Vec>{
    Scalar<Vec> result = 0;
    using Size = decltype(vec.size());
    for(Size  i = 0; i<vec.size();i++){
        result += vec[i] * vec[i];
    }
    return std::sqrt(result);
}

struct Point2{
    float x;
    float y;

    float operator[](int i) const {
        return i == 0 ? x :y;
    }
    int size() const {
        return 2;
    }
};

template <typename Func, typename Arg, typename Ret>
concept FuncOneArg = requires(Arg a, Func func){
    {func(a)} -> std::same_as<Ret>;
};
template <typename Func, typename Arg, typename Ret>
concept FuncwithArg = std::regular_invocable<Func, Arg> &&
                      std::same_as<std::invoke_result_t<Func, Arg>, Ret>;

template <typename Callable> requires FuncwithArg<Callable, int, int>
int call_twice(Callable callable, int arg){
    return callable(arg) + callable(arg);
}

template<class T>
void Print(const T& arg) {
    auto time_point = std::chrono::system_clock::now();
    std::time_t ttp = std::chrono::system_clock::to_time_t(time_point);
    std::cout << " >>"<< std::ctime(&ttp) << arg << std::endl;
}

int main() {
    // constexpr const std::array myArray{1, 2, 3, 4, 5};
    // constexpr const auto sum = std::accumulate(myArray.begin(), myArray.end(), 0);
    BetterMyStruct betterMyStruct{};
    MyStruct myStruct{};
    std::integral auto ff =func(betterMyStruct);
    double d2 =func(betterMyStruct);
    // func(34);
    std::vector<double> a = {1,2,3};
    auto print = [](const int& n) { std::cout << " " << n; };

    std::cout << "before:";
    norm(a);
    norm2(a);
    std::for_each(a.cbegin(), a.cend(), print);
    Point2 p {3,4};
    norm2(p);

    FloatVec auto g  = Point2{4,5};
    norm2(g);

    // std::list d {1,2};
    // norm2(d);
    call_twice([](int i){
        std::cout<<"i = " << i <<"\n";
        return i+1;

    },4);
    std::string str = "Hello World!";
    int i = 1;
    Print(str);
    Print(i);

    // return sum;
}
