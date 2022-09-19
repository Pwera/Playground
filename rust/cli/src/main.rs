extern crate clap;
use clap::{clap_app, crate_version};
use pulldown_cmark::{html::push_html, Event, Parser};
fn main() {
    let clap = clap_app!(
        mdrend =>
            (version: crate_version!())
            (author: "Piotr Wera")
            (@arg input: +required "Sets the input file"))
    .get_matches();
    let file = clap.value_of("input");
    println!("Input {:?}", file);
    let infile = std::fs::read_to_string(file.unwrap()).expect("Could not read file");

    let mut res = String::new();
    let ps = Parser::new(&infile);
    let ps: Vec<Event> = ps.into_iter().collect();
    for p in &ps {
        println!("{:?}", p);
    }
    push_html(&mut res, ps.into_iter());
}
