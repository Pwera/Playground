use crate::Line;
use crate::Point;

pub fn function() {
    //functions
    println!("{}", multiply(90, 34));
    let mut i = 0;
    increase(&mut i);
    println!("={}", i);
    let p1 = Point { x: 1.0, y: 0.0 };
    let p2 = Point { x: 0.0, y: 10.0 };
    let l1 = Line { startPoint: p1, endPoint: p2 };
    println!("{}", l1.len());

    let st = hello;
    st();

    let plus_one = |x: i32| -> i32 { x + 1 };
    println!("{}", plus_one(1));

    // let plus_two = |x| {
    //     let mut z =x;
    //     z+=2;
    //     z
    // };
}

fn multiply(a: i32, b: i32) -> i32 {
    a + b
}

fn increase(i: &mut i32) {
    *i += 1;
}

fn hello() {
    println!("helloo")
}