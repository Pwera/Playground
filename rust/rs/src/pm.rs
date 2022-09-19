#![allow(unused_variables)]
#![allow(dead_code)]
#[warn(unused_mut)]

pub fn how_many(x: i32) -> &'static str {
    match x {
        0 => "no",
        1 | 2 => "one or two",
        z @ 9...11 => "a lot of",
        _ if (x % 2 == 0) => "some",
        _ => "a few"
    }
}

pub fn pattern() {
    for x in 0..13 {
        println!("{} I have {} orages", x, how_many(x))
    }
    let point = (0, 4);
    match point {
        (0, 0) => println!("origin"),
        (0, y) => println!("x axis {}", y),
        (x, 0) => println!("y axis {} ", x),
        (x, y) => println!("{} {}", x, y)
    }
}