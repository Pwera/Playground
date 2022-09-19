#[cfg(test)]

#[test]
fn testhello(){
    assert_eq!("hello", "hello");
}

#[test]
#[should_panic]
fn testShouldPanic(){
    assert_eq!("hello", "hello__");
}

#[test]
#[ignore]
fn testIgnore(){

}

