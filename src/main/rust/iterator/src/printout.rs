fn main() {
    print_a(&vec!["hello".to_string(), "world".to_string()]);

    print_c(vec!["hello".to_string(), "world".to_string()].into_iter());
    // print_c(["hello".to_string(), "world".to_string()].into_iter());

    print_any((&["hello".to_string(), "world".to_string()]).into_iter());
    print_any((&["hello", "world"]).into_iter());
    print_any(vec!["hello".to_string(), "world".to_string()].into_iter());
    print_any(vec!["hello", "world"].into_iter());
}

fn print_a(v: &Vec<String>) {
    for (i, val) in v.into_iter().enumerate() {
        println!("{} == {}", i, val);
    }
}

fn print_c<I: Iterator<Item=String>>(v: I) {
    for (i, val) in v.enumerate() {
        println!("{} == {}", i, val);
    }
}

fn print_any<S: AsRef<str>, I: IntoIterator<Item=S>>(v: I) {
    for (i, val) in v.into_iter().enumerate() {
        println!("{} == {}", i, val.as_ref());
    }
}