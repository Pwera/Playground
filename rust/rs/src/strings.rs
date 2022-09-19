#![allow(unused_mut)]

pub fn string() {
    // String
    let s = "hello";  //string slice
    let s2: &'static str = "hello";
//            s2="";      // compile time error
//            let h = s[0]; // compile time error

    for c in s2.chars().rev() {
        println!("{}", c);
    }

    // String on heap
    let mut letters = String::new();

    let abc = String::from("hello");
    let abv = "hello".to_string();
}