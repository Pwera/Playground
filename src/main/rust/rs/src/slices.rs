pub fn slice() {
    // Slices
    let mut data = [1, 2, 3, 4, 5, 6];
    use_slice(&mut data[1..4]);// not includeing 4
    use_slice(&mut data);
    println!("{:?}", data);
}

fn use_slice(slice: &mut [i32]) {
    println!("first {}", slice[0]);
    slice[0] = 999;
}