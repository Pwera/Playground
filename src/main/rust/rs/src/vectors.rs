pub fn vector() {
    // Vector
    let mut a = Vec::new();
    a.push(1);
    a.push(2);
    a.push(3);
    a.push(4);
    println!("{:?}", a);
    println!("{}", a[0]);
    a[0] = 213123;
    println!("{:?}", a);
    //println!("{}", a[23]);
    //return option type
    match a.get(6) {
        Some(z) => println!("{}", z),
        None => println!("{:?}", a)
    }

    let mut c = Vec::new();
    c.push(7);

    println!("{:?}", c.pop());
    println!("{:?}", c.pop());

    while let Some(x) = a.pop() {
        println!("{} ", x);
    }
}