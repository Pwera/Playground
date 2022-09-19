
pub fn option(){
    let x = 3.0;
    let y = 2.0;
    let result: Option<f64> = if y != 0.0 { Some(x / y) } else { None };
    println!("{:?}", result);

    match result {
        Some(z) => println!(" ok {}", z),
        None => println!("not ok")
    }

    if let Some(z) = result { println!("z = {}", z); }
}

//pub fn option2(){
//    let x = 3.0;
//    let y = 2.0;
//
//    let result :Option<f64> =
//    if y != 0.0 {Some(x/y)} else {None};
//
//    println!("{:?}", result);
//
//    match result{
//
//    }
//
//}