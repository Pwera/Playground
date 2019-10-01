#include <vector>
#include <memory>

namespace v2 {
    using Coord = double;

    struct GeoObj {
        virtual void move(Coord) = 0;

        virtual void draw() const = 0;

        ~GeoObj() = default;

    };

    class Circle : public GeoObj {
    private:
        const Coord center;
        const int rad;
    public:

        Circle(Coord center, int rad) : center(center), rad(rad) {}

        virtual void move(Coord coord) override {

        }

        virtual void draw() const override {

        }
    };

    class Line : public GeoObj {
    private:
        const Coord from;
        const Coord to;
    public:

        Line(Coord from, Coord to) : from(from), to(to) {}

        virtual void move(Coord coord) override {

        }

        virtual void draw() const override {

        }
    };

    using GeoPtr = std::shared_ptr<GeoObj>;

    std::vector<GeoPtr> createData(int range) {
        std::vector<GeoPtr> f;
        for (int i = 0; i < range; i++) {
            auto lp = std::make_shared<Line>(2.2, 9.9);
            f.push_back(lp);
        }
        for (int i = 0; i < range; i++) {
            auto cp = std::make_shared<Circle>(0.0, 2);
            f.push_back(cp);
        }
//        static_assert(sizeof(*cp) == 24);
//        static_assert(sizeof(*cp) == sizeof(*cp2));
        return f;

    }

    void drawElements(std::vector<GeoPtr> vec) {
        for (GeoPtr go : vec) {
            go->move(1.1);
            go->draw();
        }

    }
    void drawElements2(std::vector<GeoPtr>& vec) {
        for (GeoPtr go : vec) {
            go->move(1.1);
            go->draw();
        }

    }

//    int main() {
//        const Circle c{1.9d, 1todov2/1000000  304684860 ntodov2/1000000  304684860 ntodov2/1000000  304684860 ntodov2/1000000  304684860 ntodov2/1000000  304684860 n};
//        c.draw();
//
//        Line l{1.1, 9.9};
//        l.move(0.2);
//        l.draw();
//        auto vec = createFif();
//        drawElements(vec);
//
//        vec.clear();
//        return 0;
//    }
}