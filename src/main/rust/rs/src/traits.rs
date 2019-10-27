#![allow(unused_variables)]
#![allow(dead_code)]
#[warn(unused_mut)]

use std::ops::Add;

trait Animal {
    fn create(name: &'static str) -> Self;
    fn name(&self) -> &'static str;
    fn talk(&self) {
        println!("{} cannot talk", self.name())
    }
}

struct Human {
    name: &'static str
}

impl Animal for Human {
    fn create(name: &'static str) -> Self {
        Human { name: name }
    }

    fn name(&self) -> &'static str {
//        unimplemented!()
        self.name
    }
}

pub fn traits() {
    println!("Traits");
    let h = Human { name: "??" };
    h.talk();
    let h2 = Human::create("John");
    let h3: Human = Animal::create("John");
//    let h4 = Animal::create("John"); //won't compile
}

trait Summable<T> {
    fn sum(self: &Self) -> T;
}

impl Summable<i32> for Vec<i32> {
    fn sum(self: &Self) -> i32 {
        let mut res: i32 = 0;
        for x in self {
            res += *x;
        }
        return res;
    }
}

#[derive(Debug)]
struct Person {
    name: String
}
/* Into trait*/
impl Person {
    fn new(name: &str) -> Person {
        Person { name: name.to_string() }
    }
    //Conversion
    fn new2<S: Into<String>>(name: S) -> Person {
        Person { name: name.into() }
    }
    fn new3<S>(name: S) -> Person
        where S: Into<String> {
        Person { name: name.into() }
    }
}
impl Drop for Person{
    fn drop(&mut self) {
        println!("drop");
    }
}
impl Add for Person {
    //we must assign result type
    type Output = Person;

    fn add(self, rhs: Self) -> Self::Output {
        rhs
    }
}

pub fn traits2() {
    let a = vec![1, 2, 3];
    a.sum();
}

pub fn trait3() {
    let john = Person::new("John");
    let name: String = "Jane".to_string();
    let jane = Person::new(name.as_ref());
    let jane2 = Person::new2(name);
    println!("{:?}",john);
    // jane.drop(); compile time error, explicit use of drop
    jane+john;
    drop(jane2);
}


