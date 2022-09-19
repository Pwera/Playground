#![allow(dead_code)]


mod pm;
mod traits;
mod tuples;
mod functions;
mod generics;
mod strings;
mod vectors;
mod slices;
mod arrays;
mod options;
mod commons;
mod sd;
mod ownerships;
mod lifetime;
mod mpsc;

const MEM: u8 = 42;
static mut XX: u8 = 45;

struct Point {
    x: f64,
    y: f64,
}

struct Line {
    startPoint: Point,
    endPoint: Point,
}

impl Line {
    fn len(&self) -> f64 {
        let dx = self.startPoint.x - self.endPoint.x;
        let dy = self.startPoint.y - self.endPoint.y;
        (dx * dx + dy * dy).sqrt()
    }
}

enum Color {
    Red,
    Green,
    Blue,
    RgbColor(u8, u8, u8),
    Cmyk { cyan: u8 },
}

union IntOrFloat {
    i: i32,
    f: f32,
}

fn main() {
    commons::common();
    commons::structure();
    commons::enumm();
    commons::union();
    options::option();
    arrays::array();
    vectors::vector();
    slices::slice();
    strings::string();
    tuples::tuple();
    pm::pattern();
    generics::generic();
    functions::function();
    {
        // high order functions
        // functions that take functions
        // functions that return functions -> generators
    }
    traits::traits();
    traits::trait3();
    sd::static_();
    sd::dynamic_();
    sd::dynamic_2();
    sd::dynamic_3();
    ownerships::ownership();
    ownerships::borrowing();
    lifetime::lifetime();
    lifetime::lifetime2();
    lifetime::reference_counting();
    lifetime::arc();
    mpsc::f();

}
