pub fn tuple() {
// Tuples
    let x = 3;

    let y = 4;
    let tup = use_tuple(x, y);
    println!("{:?}", tup);
    println!("{} {}", tup.0, tup.1);

    let (a, b) = tup; // destructuring
}

fn use_tuple(x: i32, y: i32) -> (i32, i32) { // tuple implicit
    (x + y, x * x)
}