pub trait Shape{
    fn as_String(&self) -> String;
    fn area(&self) -> f32;
}

struct Rectangle{
    width: f32,
    height: f32
}
struct Ellipse{
    rad_x: f32,
    rad_y: f32
}
struct Circle{
    rad: f32,
}

impl Shape for Rectangle {
   fn as_String(&self) -> String {
        format!("Rectangle: width: {} height: {} area: {}", self.width, self.height, self.area())
    }
    fn area(&self) -> f32{
        self.width * self.height
    }
}
impl Shape for Circle {
   fn as_String(&self) -> String {
        format!("Circle: rad: {} area: {}", self.rad, self.area())
    }
    fn area(&self) -> f32{
        self.rad * self.rad * std::f32::consts::PI / 2f32
    }
}
fn main() {
    
 let mut v : Vec<&dyn Shape> = vec![];
 v.push(&Rectangle{width: 2., height: 2.});
 println!("{}",v[0].as_String());
 let  r =Circle{rad: 2.34};
 v.push(&r);
for &s in v.iter(){
    println!("{}", s.as_String());
}
println!("{}", (1..101)
    .map(|x| x * 2)
    .fold(0, |x, y| x + y));
}