#include <vector>

namespace v1 {
    using Coord = double;

    struct GeoObj {
        virtual void move(Coord) = 0;

        virtual void draw() const = 0;

        ~GeoObj() = default;

    };

    class Circle : public GeoObj {
    private:
        Coord center;
        int rad;
    public:

        Circle(Coord center, int rad) : center(center), rad(rad) {}

        virtual void move(Coord coord) override {

        }

        virtual void draw() const override {

        }
    };

    class Line : public GeoObj {
    private:
        Coord from;
        Coord to;
    public:

        Line(Coord from, Coord to) : from(from), to(to) {}

        virtual void move(Coord coord) override {

        }

        virtual void draw() const override {

        }
    };

    std::vector<GeoObj *> createData(int range) {
        std::vector<GeoObj *> f;
        for (int i = 0; i < range; i++) {
            Line *lp = new Line(1.1, .6);
            f.push_back(lp);
        }
        for (int i = 0; i < range; i++) {
            Circle *cp = new Circle(0.0, 2);
            f.push_back(cp);
        }
//        static_assert(sizeof(*lp) == 24);
//        static_assert(sizeof(*cp) == 24);
        return f;

    }

    void drawElements(std::vector<GeoObj *> &vec) {
        for (GeoObj *go : vec) {
            go->move(1.1);
            go->draw();
        }

    }

    void removeElements(std::vector<GeoObj *> &vec) {
        for (GeoObj *&go : vec) {
            delete go;
            go = nullptr;
        }

    }

//    int main() {
//        Circle c{1.9d, 1};
//        c.move(1.2d);
//        c.draw();
//
//        Line l{1.1, 9.9};
//        l.move(0.2);
//        l.draw();
//        auto vec = createFif();
//        drawElements(vec);
//        // if we dont don't clear vec we have memory leak
//        removeElements(vec);
//
//        return 0;
//    }
}