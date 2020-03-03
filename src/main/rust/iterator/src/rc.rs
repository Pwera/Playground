use std::rc::Rc;
use std::cell::RefCell;

#[derive(Debug)]
pub struct WithLife<'a> {
    s: &'a String,
}

#[derive(Debug)]
pub struct NoLife {
    s: Rc<RefCell<String>>,
}


fn main() -> Result<(), std::io::Error> {
    let (l,r) = make_with_nolife("src/testdata.txt")?;
    let mut s = l.s.borrow_mut();
    println!("l = {:?}", l);
    println!("r = {:?}", r);
    println!("s = {:?}", s);
    Ok(())
}

// fn make_with_life<'a>(fname: &str) -> Result<(WithLife<'a>, WithLife<'a>), std::io::Error> {
//     let s = std::fs::read_to_string(fname)?;
//     Ok((WithLife { s: &s }, WithLife { s: &s }))
// }

fn make_with_nolife(fname: &str) -> Result<(NoLife, NoLife), std::io::Error> {
    let s = std::fs::read_to_string(fname)?;
    let rc = Rc::new(RefCell::new(s));
    Ok((NoLife { s: rc.clone() }, NoLife { s: rc }))
}
