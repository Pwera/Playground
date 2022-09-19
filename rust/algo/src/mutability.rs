use std::ops::AddAssign;

#[derive(Debug, Clone)]
struct Person {
    name: String,
    age: i32,
}

#[derive(Debug, Clone, Copy)]
struct Point {
    x: f64,
    y: i32,
}

impl Person {
    fn print(&self) -> String {
        format!("[__{:?}__]", self)
    }

    fn age_up(&mut self) {
        self.age += 1;
    }
}

fn get_age(s: &Person) -> &i32 {
    &s.age
}

pub fn mutabilityfunction() {
    let mut x = 34;
    let y = x;
    x += 5;
    println!("{} {}", x, y);

    let p1 = Person { name: "name".to_string(), age: 30 };

    let p2 = p1.clone(); // let p2 = p1 will fail
    println!("{:?} {:?}", p1, p2);

    let mut po = Point { x: 1f64, y: 2 };

    let po2 = po;
    po.x = 6666f64;

    println!("{:?} {:?}", po, po2);

    lifetime2();
    lifetime3();
    lifetime4();
    lifetime5();
}

fn lifetime2() {
    let mut p = Person { name: "name".to_string(), age: 0 };
    // p.age_up();

    let a = get_age(&p);


    // In Rust we can borrow as many time as borrow is immutable.
    // But if one is mutable, cannot borrow any time
    // &self get back what goes to function
    // self  never gives back // MOVING VALUE
    // &mut self, function has bto be invoked on mutable object, requires mutable object
    // cannot borrow p as mutable because it is also borrow as immutable
    // mutable borrow(&mut XX) can be placed if there is no immutable borrow
    println!("{}", a);
    p.age_up();
}

#[derive(Debug)]
struct LinkedList<T> {
    data: T,
    next: Option<Box<LinkedList<T>>>,   // if Option<LinkedList<T>> we don't know exact size of next
    // if there is nothing size of data pointer -> u64
    // In debug Box will say what my child is
}

impl<T: AddAssign> LinkedList<T> {
    pub fn add_up(&mut self, n: T) {
        self.data += n;
    }
}

fn lifetime3() {
    // use heap if don't know how big objects will be
    // Box
    let lk2 = LinkedList {
        data: 3,
        next: None,
    };

    let mut lk = LinkedList {
        data: 5,
        next: Some(Box::new(lk2)),
    };

    println!("{:?}", lk);

    lk.add_up(45);
    println!("{:?}", lk);

    if let Some(ref mut x) = lk.next {
        x.add_up(10);
    }
    println!("{:?}", lk);
}

struct AcceptString {
    field: String,
}

impl AcceptString {
    fn new(field: String) -> Self {
        Self { field }
    }
    fn doWork(&self) {
        println!("{:?}", self.field);
    }
}

fn lifetime4() {
    // let mut v: Vec<String> = Vec::new();
    //   v.iter_mut()
    //     .map(AcceptString::new);

    let mut y: Vec<String> = Vec::new();

    y.push("some text".to_string());
    y.push("".to_string());
    let vec = y.iter()
        .map(|x: &String| AcceptString::new(x.into()))
        .collect::<Vec<AcceptString>>();
    vec.iter()
        .for_each(|x| x.doWork());

    // println!("{:?}", vec);


    let z: Vec<String> = Vec::with_capacity(100);
}

fn lifetime5() {
    let mut s: &str = "   x  x xxx    "; //exist as a part of a file, set of bytes. Pointer to memmory to unchangable memory.
    // If we try to modify we have comp error.
    let p = s.trim();           // p is subset of s
    s = "??";
    println!("{} {}", p, s);

    let find_if = string_find_if("help me find home");
    println!("{}", find_if);
}

fn string_find_if(s: &str) -> &str {
    let n = 0;
    for (n, x) in s.char_indices() {
        if x == 'f' {
            return &s[n..];
        }
    }
    s
}