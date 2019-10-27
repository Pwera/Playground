#![allow(unused_variables)]
#![allow(dead_code)]
#![allow(unused_mut)]
#![allow(unused_assignments)]

use std::rc::Rc;
use std::thread;
use std::sync::{Arc,Mutex};

struct Worker {
    name: String
}

struct Worker2<'a> {
    name: &'a str
}

struct Worker3 {
    name: Rc<String>
}
struct Worker4 {
    name: Arc<String>,
    state: Arc<Mutex<String>>
}

impl Worker {
    fn get_rf(&self) -> &String {
        &self.name
    }
}

struct Company<'z> {
    name: String,
    ceo: &'z Worker,
}

pub fn lifetime() {
    let boss = Worker { name: "ceo".to_string() };
    let tesla = Company { name: "test".to_string(), ceo: &boss };

    let mut z: &String;
    {
        let p = Worker { name: String::from("John") };
        z = p.get_rf();
    }
}

pub fn lifetime2() {
    let s = "??";
    let z = Worker2 { name: s };
}

impl Worker {
    fn new(name: String) -> Worker {
        Worker { name: name }
    }
    fn greet(&self) {
        println!("{}", self.name);
    }
}

impl Worker3 {
    fn new(name: Rc<String>) -> Worker3 {
        Worker3 { name: name }
    }
    fn greet(&self) {
        println!("{}", self.name);
    }
}
impl Worker4 {
    fn new(name: Arc<String>,state: Arc<Mutex<String>>) -> Worker4 {
        Worker4 { name: name, state: state  }
    }
    fn greet(&self) {
        let mut state = self.state.lock().unwrap();
        println!("{}", self.name);
        state.clear();
        state.push_str("modified");
    }
}

pub fn reference_counting() {
    let n = "work".to_string();
    let w = Worker::new(n);
    w.greet();
//    println!("{:?}", n); // wont compile


    let n2 = Rc::new("work".to_string());
    {
        println!("{}", Rc::strong_count(&n2) );
    }
    let w3 = Worker3::new(n2.clone());
    {
        println!("{}", Rc::strong_count(&n2) );
    }
    println!("{:?}", n2);
    {
        println!("{}", Rc::strong_count(&n2) );
    }
}


pub fn arc(){
    let n3 = Arc::new("work".to_string());
    let state = Arc::new(Mutex::new("state".to_string()));
    let w = Worker4::new(n3.clone(), state.clone());
    let t = thread::spawn( move || {
        w.greet();
    });
    t.join().unwrap();
    println!("{}",n3);

}