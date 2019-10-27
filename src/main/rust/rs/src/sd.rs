#![allow(unused_variables)]
#![allow(dead_code)]
#[warn(unused_mut)]

trait Printable {
    fn format(&self) -> String;
}

impl Printable for i32 {
    fn format(&self) -> String {
        format!("i32: {}", *self)
    }
}

impl Printable for String {
    fn format(&self) -> String {
        format!("String: {}", *self)
    }
}

// monomorphisation
fn print_it<T: Printable>(z: T) {
    println!("{}", z.format());
}

fn print_it2(z: &Printable) {
    println!("{}", z.format());
}

pub fn static_() {
    // compile time
    let a: i32 = 123;
    let b = "xyz".to_string();
    println!("{}", a.format());
    println!("{}", b.format());
    print_it(a);
    print_it(b);
}

pub fn dynamic_() {
    // runtime
    let a: i32 = 123;
    let b = "xyz".to_string();
    print_it2(&a);
    print_it2(&b);
}

struct Circle { radius: f64 }

struct Square { side: f64 }

trait Shape {
    fn area(&self) -> f64;
}

impl Shape for Square {
    fn area(&self) -> f64 {
        self.side * self.side
    }
}

impl Shape for Circle {
    fn area(&self) -> f64 {
        self.radius * self.radius * std::f64::consts::PI
    }
}

pub fn dynamic_2() {
    let shapes:[&Shape; 4] = [
        &Circle{radius:0.0},
        &Circle{radius:14.0},
        &Square{side:5.6},
        &Square{side:15.3}
    ];
    for (i, shape) in shapes.iter().enumerate() {
        println!("{:?}", shape.area());
    }
}

pub fn dynamic_3() {
    let mut h:Vec<&Shape> = Vec::new();
    h.push(&Circle{radius:0.0});
    h.push(&Square{side:10.0});
    for (i, shape) in h.iter().enumerate() {
        println!("{:?}", shape.area());
    }
    for u in h{
        println!("{:?}", u.area());
    }

    let mut g:Vec<Box<Shape>> = Vec::new();
    g.push(Box::new(Circle{radius:0.0}));
    g.push(Box::new(Square{side:10.0}));

    for u in g.iter(){
        println!("{:?}", u.area());
    }

}




