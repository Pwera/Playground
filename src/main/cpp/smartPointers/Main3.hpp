#include <vector>
#include <memory>
#include <variant>

namespace v3 {
    using Coord = double;

    template<class... Ts>
    struct overload : Ts ... {
        using Ts::operator()...;
    };
    template<class... Ts> overload(Ts...) -> overload<Ts...>;

    class Circle {
    private:
        const Coord center;
        const int rad;
    public:

        Circle(Coord center, int rad) : center(center), rad(rad) {}

        void move(Coord coord) {

        }

        void draw() const {

        }
    };

    class Line {
    private:
        const Coord from;
        const Coord to;
    public:

        Line(Coord from, Coord to) : from(from), to(to) {}

        void move(Coord coord) {

        }

        void draw() const {

        }
    };

    using GeoPtr = std::variant<Line, Circle>;

    std::vector<GeoPtr> createData(int range) {
        std::vector<GeoPtr> f;
        for (int i = 0; i < range; i++) {
            Line lp{2.2, 9.9};
            f.push_back(lp);
        }
        for (int i = 0; i < range; i++) {
            Circle cp{0.0, 2};
            f.push_back(cp);
        }
//        static_assert(sizeof(lp) == 16);
//        static_assert(sizeof(cp) == 16);
        return f;

    }

    void drawElements(const std::vector<GeoPtr> &vec) {

        for (const auto &go : vec) {
//            std::visit([](const auto &obj) {
//                obj.draw();
//            }, go);


//            std::visit(overload{
//                    [](Line &l) {
//                        l.draw();
//                        l.move(1.1);
//                    },
//                    [](Circle &c) {
//                        c.draw();
//                        c.move(2.2);
//                    }
//            }, go);

            std::visit(overload{
                    [](Line l) {
                        l.move(1.1d);
                        l.draw();
                    },
                    [](Circle c) {
                        c.move(1.1d);
                        c.draw();
                    }
            }, go);
        }

    }


//    int main() {
//        const Circle c{1.9d, 1};
//        c.draw();
//
//        Line l{1.1, 9.9};
//        l.move(0.2);
//        l.draw();
//        const auto vec = createFif();
//        drawElements(vec);
//
//        return 0;
//    }
}