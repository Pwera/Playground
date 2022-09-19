#![allow(unused_mut)]

use std::mem;

pub fn array() {
    //arrays
    let mut a: [i32; 5] = [1, 2, 3, 4, 5];
    let mut a2 = [1, 2, 3, 4, 5];
    println!("{} {} {:?}", a.len(), a[0], a);
    a[0] = 45;
    println!("{} {} {:?}", a.len(), a[0], a);
    if a == [45, 2, 3, 4, 5] {
        println!("same")
    }
//            if a == [45, 2, 3] { //compile tiem error
//                println!("same")
//            }else {
//                println!("not same")
//            }
    let b = [1; 10]; // 10 elements with 1
    for i in 0..b.len() {
        println!("{}", b[i])
    }

    println!("{} bytes", mem::size_of_val(&b));

    let c = [1u16; 10]; // 10 elements with 1
    println!("{} bytes", mem::size_of_val(&c));

    let d = [1u8; 10]; // 10 elements with 1
    println!("{} bytes", mem::size_of_val(&d));

    // 2 rows 3 columns
    let mtx: [[f32; 3]; 2] =
        [
            [1.0, 2.0, 3.0],
            [1.0, 2.0, 3.0]
        ];
    println!("{:?}", mtx)
}