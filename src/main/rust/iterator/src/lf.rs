pub struct Hider{
    pub public: String,
    hidden: String,
}

impl Hider {
    pub fn new(public: String, hidden: String) -> Self{
        Hider{
            public, hidden
        }
    }

    pub fn edit<F>(&mut self, f: F) where F: FnOnce(&mut String){

    }
}


fn main(){
    let mut h = Hider::new("p".to_string(), "h".to_string());
    h.edit(|x| {
        println!(x.to_string());
    })("????");
}