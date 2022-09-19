#![allow(unused_mut)]
#![allow(unused_assignments)]

pub fn ownership() {
    let v = vec![1, 2, 3];

    // copying a pointer
    let v2 = v;
//    println!("{:?}", v); // compile time error
    println!("{:?}", v2);


    let n = vec![1, 2, 3];
    let foo = |v: Vec<i32>| ();
    foo(n); // we no longer can use n variable
//    println!("{:?}", n); compile time error


    let o = vec![1, 2, 3];
    let func = |x: Vec<i32>| -> Vec<i32> {
        x
    };
    let oo = func(o);
}

pub fn borrowing() {
    let v = vec![1, 2, 3];
    {
        let func = |x: &Vec<i32>| {
//        x.push(23);
            println!("{:?}", x);
        };
        func(&v);
    }

    println!("{:?}", v);

    // mutable vec
    let mut z = vec![1, 2, 3];

    for i in &z {
        println!("{}", i)
    }
    z.push(4);

    for i in &z {
        println!("{}", i)
    }


    {
        let func = |mut x: Vec<i32>| {
            x.push(23);
            println!("{:?}", x);
        };

        func(z);
    }

//    println!("{:?}", z);
}