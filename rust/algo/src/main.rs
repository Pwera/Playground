use crate::mutability::mutabilityfunction;

mod error;
mod myservice;
mod mutability;
mod sort;
mod rand;

// use crate::error::Res;

#[derive(Debug)]
pub struct Person {
    name: String,
    age: i32,
}

impl Person {
    pub fn new(self) -> String {
        format!("{:?}", self)
    }
}

#[derive(Debug)]
pub enum Color {
    Red(String),
    Green,
    Blue,
}

pub struct Stepper {
    curr: i32,
    step: i32,
    max: i32,
}

impl Iterator for Stepper {
    type Item = i32;
    fn next(&mut self) -> Option<Self::Item> {
        if self.curr >= self.max {
            return None;
        }
        let res = self.curr;
        self.curr += self.step;
        Some(res)
    }
}

fn main() {
    let person = Person {
        name: "Peter".to_string(),
        age: 3,
    };
    let s: String = person.new();
    println!("{}", s);

    // let c= Color::Green;
    let c = Color::Red("some".to_string());

    match c {
        Color::Blue => println!("blue"),
        Color::Red(s) => println!("red {}", s),
        Color::Red(_) => println!("red"),
        _ => println!("??"),
    }

    let res1 = error::devide(0);
    match res1 {
        Ok(v) => println!("Thing: {:?}", v), // same as Result::Ok(v) => println!("Thing: {:?}", v),
        _ => println!("Something else")
    }

    // if we care about only one case
    if let Ok(v) = res1 { // same as if let Result::Ok(v) = res1 {
        println!("Thing: {:?}", v);
    }

    // loop 1
    let mut n = 0;
    loop {
        n += 1;
        if n > 4 {
            break;
        }
        println!("{}", n);
    }

    // loop 2
    n = 0;
    while n < 4 {
        n += 1;
        println!("{}", n);
    }

    // loop 3
    for i in 1..5 {
        println!("{}", i);
    }

    let mut st = Stepper {
        curr: 2,
        step: 3,
        max: 15,
    };

    loop {
        match st.next() {
            Some(v) => println!("match inside loop -> {:?}", v),
            None => break,
        }
    }

    let mut st2 = Stepper {
        curr: 3,
        step: 4,
        max: 15,
    };

    while let Some(n) = st2.next() {
        println!("while let Some -> {:?}", n)
    }

    let st3 = Stepper { // doesn't have to be mutable
        curr: 3,
        step: 4,
        max: 15,
    };
    for i in st3 {
        println!("for i in Stepper -> {:?}", i)
    }

    mutabilityfunction();
    sort::sort();
}



