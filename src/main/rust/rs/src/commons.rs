#![allow(unused_mut)]

use std::mem;
use crate::{Point, XX, Line, Color, IntOrFloat};

pub fn common() {
    let a: u8 = 127;
    {
        println!("u8 = {} \t\t{}", a, mem::size_of_val(&a));
        let b = 1273456789;
        println!("u32 = {} {}", b, mem::size_of_val(&b));
        let d = 'x';
        println!("char {} \t\t{}", d, mem::size_of_val(&d));
        let e = 2.5;
        println!("f64 {} \t\t{}", e, mem::size_of_val(&e));
        {
            unsafe {
                XX = 2;
            }
        }
        let x = Box::new(4);
        println!("{} {} ", &x, *x);
        let p1 = origin();
        let p2 = Box::new(origin());
        println!("p1: {} bytes", mem::size_of_val(&p1));
        println!("p2: {} bytes", mem::size_of_val(&p2));
        let temp = 45;
        let res = if temp < 23 { "ok" } else { "bad" };
        for x in 1..3 {
            println!("X")
        }

        for (pos, y) in (30..41).enumerate() { // 41 excluded
            println!("{} : {}", pos, y)
        }

        let code = 56;
        let country = match code {
            1 => "PL",
            2 => "GER",
            1...999 => "SOME", //999 included
            _ => "UNK"
        };
    }
}

pub fn structure() {
//Structures
    let p = Point { x: 1.0, y: 9.9 };
    println!("{} : {}", p.x, p.y);

    let line = Line { startPoint: p, endPoint: Point { x: 0.0, y: 0.0 } };
}

pub fn enumm() {
    // enums
    let c = Color::Red;
    let c2: Color = Color::Blue;

    let res = match c {
        Color::Red => "Red",
        Color::Blue => "Blue",
        Color::RgbColor(0, 0, 0) => "Black",
        _ => "Green",
    };
}

pub fn union() {
    //union
    let mut u = IntOrFloat { f: 4.4 };
    u.i = 5;
    let value = unsafe { u.i };
}


fn origin() -> Point {
    Point { x: 0.0, y: 0.0 }
}

//fn mathcUnion(u: IntOrFloat) {
//    unsafe {
////        match u {
////            IntOrFloat { i: 42 } => {
////                println!("X")
////            }
////            IntOrFloat { f } => {
////                println!("Y")
////            }
////        }
//    }
//}